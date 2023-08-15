package auth_test

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"gotest.tools/v3/assert"

	"github.com/artefactual-sdps/enduro/internal/wellcome/auth"
)

type params struct {
	grantType string
	clientID  string
	secretID  string
}

type expected struct {
	token  *auth.AccessToken
	errMsg string
}

type mockHttpServer struct {
	requestBody []byte

	responseBody   []byte
	responseStatus int
}

func newMockServer() *mockHttpServer {
	return &mockHttpServer{
		requestBody:    []byte("client_id=enduro&client_secret=be6c5a55&grant_type=client_credentials"),
		responseBody:   []byte(`{"access_token":"eyJraWQi...","expires_in":3600,"token_type":"Bearer"}`),
		responseStatus: http.StatusOK,
	}
}

func (m *mockHttpServer) Start(t *testing.T) *httptest.Server {
	t.Helper()

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		b, err := io.ReadAll(r.Body)
		assert.NilError(t, err)
		values, err := url.ParseQuery(string(b))
		assert.NilError(t, err)

		for _, k := range []string{
			"client_id", "client_secret", "grant_type",
		} {
			if values.Get(k) == "" {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte(fmt.Sprintf(`{"error":"missing required value: %s"}`, k)))
				return
			}
		}

		assert.DeepEqual(t, b, m.requestBody)
		assert.Equal(t, r.Header.Get("Content-Type"), "application/json")

		w.WriteHeader(m.responseStatus)
		w.Write(m.responseBody)
	}))

	return server
}

func (m *mockHttpServer) SetResponseStatus(d int) *mockHttpServer {
	m.responseStatus = d

	return m
}

func (m *mockHttpServer) SetResponseBody(s string) *mockHttpServer {
	m.responseBody = []byte(s)

	return m
}

func TestGetAccessToken(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		server *mockHttpServer
		args   params
		want   expected
	}{
		{
			name: "Returns a valid access token",
			args: params{
				grantType: "client_credentials",
				clientID:  "enduro",
				secretID:  "be6c5a55",
			},
			server: newMockServer(),
			want: expected{
				token: &auth.AccessToken{
					Value:     "eyJraWQi...",
					ExpiresIn: 3600,
					Type:      "Bearer",
				},
			},
		},
		{
			name:   "Bad request",
			server: newMockServer(),
			want: expected{
				errMsg: `access token error response: status: "400 Bad Request", body: "{\"error\":\"missing required value: client_id\"}"`,
			},
		},
		{
			name: "Internal Server Error",
			args: params{
				grantType: "client_credentials",
				clientID:  "enduro",
				secretID:  "be6c5a55",
			},
			server: newMockServer().SetResponseStatus(500).SetResponseBody(""),
			want: expected{
				errMsg: `access token error response: status: "500 Internal Server Error", body: ""`,
			},
		},
		{
			name: "Invalid access token error",
			args: params{
				grantType: "client_credentials",
				clientID:  "enduro",
				secretID:  "be6c5a55",
			},
			server: newMockServer().SetResponseBody(`{"access_token":""}`),
			want: expected{
				errMsg: `empty access token response: {"access_token":""}`,
			},
		},
		{
			name: "Invalid JSON response",
			args: params{
				grantType: "client_credentials",
				clientID:  "enduro",
				secretID:  "be6c5a55",
			},
			server: newMockServer().SetResponseBody("hello"),
			want: expected{
				errMsg: "malformed access token response: invalid character 'h' looking for beginning of value",
			},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			server := tt.server.Start(t)
			defer server.Close()

			token, err := auth.GetAccessToken(server.URL, tt.args.grantType, tt.args.clientID, tt.args.secretID)
			if tt.want.errMsg != "" {
				assert.Error(t, err, tt.want.errMsg)
				return
			}
			assert.NilError(t, err)
			assert.DeepEqual(t, token, &auth.AccessToken{
				Value:     "eyJraWQi...",
				ExpiresIn: 3600,
				Type:      "Bearer",
			})
		})
	}
}
