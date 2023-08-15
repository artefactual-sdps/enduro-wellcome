package ingests_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"gotest.tools/v3/assert"

	"github.com/artefactual-sdps/enduro/internal/wellcome"
	"github.com/artefactual-sdps/enduro/internal/wellcome/auth"
	"github.com/artefactual-sdps/enduro/internal/wellcome/ingests"
)

type params struct {
	token         *auth.AccessToken
	aipLocation   *ingests.AIPlocation
	bagExternalID string
}

type expected struct {
	ingestID ingests.IngestID
	errMsg   string
}

type mockHttpServer struct {
	responseHeader http.Header
	responseBody   []byte
	responseStatus int
}

func newMockServer() *mockHttpServer {
	return &mockHttpServer{
		responseHeader: http.Header{
			"Location": {"/ingests/myid-1234"},
		},
		responseBody:   []byte(`{"id":"myid-1234"}`),
		responseStatus: http.StatusCreated,
	}
}

func (m *mockHttpServer) Start(t *testing.T) *httptest.Server {
	t.Helper()

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var data ingests.IngestData
		d := json.NewDecoder(r.Body)
		err := d.Decode(&data)
		assert.NilError(t, err)

		if data.SourceLocation.Path == "" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{"error":"missing required value: sourceLocation.path"}`))
			return
		}

		assert.Equal(t, r.Header.Get("Content-Type"), "application/json")
		assert.Assert(t, r.Header.Get("Authorization") != "")

		assert.Assert(t, data.Bag.Info.ExternalID != "")
		assert.Equal(t, data.Bag.Info.Type, "BagInfo")

		for k, v := range m.responseHeader {
			w.Header().Set(k, v[0])
		}
		w.WriteHeader(m.responseStatus)
		w.Write(m.responseBody)
	}))

	return server
}

func (m *mockHttpServer) SetResponseStatus(v int) *mockHttpServer {
	m.responseStatus = v

	return m
}

func (m *mockHttpServer) SetResponseBody(v string) *mockHttpServer {
	m.responseBody = []byte(v)

	return m
}

func (m *mockHttpServer) SetResponseHeader(k, v string) *mockHttpServer {
	m.responseHeader.Set(k, v)

	return m
}

func TestSubmit(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		server *mockHttpServer
		args   params
		want   expected
	}{
		{
			name: "Returns an ingest ID",
			args: params{
				token: &auth.AccessToken{
					Value:     "test-1234",
					ExpiresIn: 3600,
					Type:      "granted",
				},
				aipLocation: &ingests.AIPlocation{
					Bucket: wellcome.BucketConfig{
						Endpoint:  "http://minio.enduro-sdps:9000",
						PathStyle: true,
						Key:       "minio",
						Secret:    "minio123",
						Region:    "us-west-1",
						Bucket:    "wellcome-uploads",
					},
					Key: "Bag.zip",
				},
				bagExternalID: "3791fe45-581a-469a-bfa1-f27fcf7abff0",
			},
			server: newMockServer(),
			want: expected{
				ingestID: "myid-1234",
			},
		},
		{
			name: "Bad request",
			args: params{
				token: &auth.AccessToken{
					Value:     "test-1234",
					ExpiresIn: 3600,
					Type:      "granted",
				},
				aipLocation: &ingests.AIPlocation{},
			},
			server: newMockServer(),
			want: expected{
				errMsg: `ingest bag: unexpected response: status: "400 Bad Request", body: "{\"error\":\"missing required value: sourceLocation.path\"}"`,
			},
		},
		{
			name: "Internal Server Error",
			args: params{
				token: &auth.AccessToken{
					Value:     "test-1234",
					ExpiresIn: 3600,
					Type:      "granted",
				},
				aipLocation: &ingests.AIPlocation{
					Bucket: wellcome.BucketConfig{
						Endpoint:  "http://minio.enduro-sdps:9000",
						PathStyle: true,
						Key:       "minio",
						Secret:    "minio123",
						Region:    "us-west-1",
						Bucket:    "wellcome-uploads",
					},
					Key: "Bag.zip",
				},
				bagExternalID: "3791fe45-581a-469a-bfa1-f27fcf7abff0",
			},
			server: newMockServer().SetResponseStatus(500).SetResponseBody(""),
			want: expected{
				errMsg: `ingest bag: unexpected response: status: "500 Internal Server Error", body: ""`,
			},
		},
		{
			name: "Missing location header",
			args: params{
				token: &auth.AccessToken{
					Value:     "test-1234",
					ExpiresIn: 3600,
					Type:      "granted",
				},
				aipLocation: &ingests.AIPlocation{
					Bucket: wellcome.BucketConfig{
						Endpoint:  "http://minio.enduro-sdps:9000",
						PathStyle: true,
						Key:       "minio",
						Secret:    "minio123",
						Region:    "us-west-1",
						Bucket:    "wellcome-uploads",
					},
					Key: "Bag.zip",
				},
				bagExternalID: "3791fe45-581a-469a-bfa1-f27fcf7abff0",
			},
			server: newMockServer().
				SetResponseHeader("Location", "").
				SetResponseBody(`{}`),
			want: expected{
				errMsg: `ingest bag: missing "Location" response header`,
			},
		},
		{
			name: "Missing ingest ID",
			args: params{
				token: &auth.AccessToken{
					Value:     "test-1234",
					ExpiresIn: 3600,
					Type:      "granted",
				},
				aipLocation: &ingests.AIPlocation{
					Bucket: wellcome.BucketConfig{
						Endpoint:  "http://minio.enduro-sdps:9000",
						PathStyle: true,
						Key:       "minio",
						Secret:    "minio123",
						Region:    "us-west-1",
						Bucket:    "wellcome-uploads",
					},
					Key: "Bag.zip",
				},
				bagExternalID: "3791fe45-581a-469a-bfa1-f27fcf7abff0",
			},
			server: newMockServer().
				SetResponseHeader("Location", "/ingests").
				SetResponseBody(`{"id":""}`),
			want: expected{
				errMsg: `ingest bag: missing ingest ID in response`,
			},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			server := tt.server.Start(t)
			defer server.Close()

			svc := ingests.NewService(&wellcome.IngestsConfig{
				ApiURL:  server.URL,
				Type:    "create",
				SpaceID: "born-digital",
			})

			ingestID, err := svc.Submit(tt.args.token, tt.args.aipLocation, tt.args.bagExternalID)
			if tt.want.errMsg != "" {
				assert.Error(t, err, tt.want.errMsg)
				return
			}
			assert.NilError(t, err)
			assert.Equal(t, ingestID, ingests.IngestID("myid-1234"))
		})
	}
}
