package spotify

import (
	"errors"
	"testing"

	"github.com/flames31/spotify-client/internal/config"
)

func TestRequestToken(t *testing.T) {
	mockAuth := &mockAuth{
		cred:  "dummy",
		token: "token",
		err:   nil,
	}
	mockConfig := &mockConfig{
		mockMap: map[string]any{
			"SPOTIFY_CLIENT_ID": "valid",
			"SPOTIFY_SECRET":    "valid",
		}}
	err := RequestToken(&config.Auth{
		Client: mockAuth,
		Config: mockConfig,
	})

	if err != nil {
		t.Errorf("unexpected err :%v", err)
	}
}

type mockAuth struct {
	cred  string
	token string
	err   error
}

type mockConfig struct {
	mockMap map[string]any
	wrote   bool
}

func (m *mockAuth) GetToken(url string, body interface{}) (string, error) {
	if m.token == "valid" {
		return m.token, m.err
	}
	return "", errors.New("invalid token")
}

func (m *mockAuth) GetCredentials() (string, string) {
	return m.cred, m.cred
}

func (m *mockConfig) Get(key string) any {
	return m.mockMap[key]
}

func (m *mockConfig) Set(key string, val any) {
	m.mockMap[key] = val
}

func (m *mockConfig) Write() error {
	m.wrote = true
	return nil
}
