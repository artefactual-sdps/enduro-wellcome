package auth

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type AccessToken struct {
	Value     string `json:"access_token"`
	ExpiresIn int    `json:"expires_in"`
	Type      string `json:"token_type"`
}

// Get OAuth2 access token.
func GetAccessToken(providerURL, grantType, clientID, secretID string) (*AccessToken, error) {
	data := url.Values{}
	data.Set("grant_type", grantType)
	data.Set("client_id", clientID)
	data.Set("client_secret", secretID)

	req, err := http.NewRequest(http.MethodPost, providerURL, strings.NewReader(data.Encode()))
	if err != nil {
		return nil, fmt.Errorf("could not create HTTP request:  %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("couldn't send access token request: %s\n", err)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("couldn't read access token response: %s\n", err)
	}
	res.Body.Close()

	if res.StatusCode >= 400 {
		return nil, fmt.Errorf("access token error response: status: %q, body: %q", res.Status, body)
	}

	var token AccessToken
	if err = json.Unmarshal(body, &token); err != nil {
		return nil, fmt.Errorf("malformed access token response: %v", err)
	}
	if token.Value == "" {
		return nil, fmt.Errorf("empty access token response: %s", body)
	}

	return &token, nil
}
