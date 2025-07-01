package spotify

import (
	"errors"
	"testing"

	"github.com/flames31/spotify-client/internal/config"
)

func TestConnectPass(t *testing.T) {
	t.Run("succesfully connects and fetches token from spotify", func(t *testing.T) {
		client := &mockClient{
			res: []byte(`{"access_token":"<mock_token>"}`),
		}

		cfg := &mockConfig{}
		cfg.mockMap = map[string]any{
			"SPOTIFY_CLIENT_ID": "test",
			"SPOTIFY_SECRET":    "test",
		}
		mockApp := &config.App{
			Client: client,
			Config: cfg,
		}

		err := Connect(mockApp)
		if err != nil {
			t.Errorf("unexpected error on connect :%v", err)
		}

		if cfg.mockMap["SPOTIFY_TOKEN"] != "<mock_token>" {
			t.Errorf("incorrect access token : wanted %v got %v", "<mock_token>", cfg.mockMap["SPOTIFY_TOKEN"])
		}
	})

	t.Run("connect correctly sends an error", func(t *testing.T) {
		client := &mockClient{
			res: []byte(`{"access_token":"<mock_token>"}`),
			err: errors.New("req failed"),
		}

		cfg := &mockConfig{}
		cfg.mockMap = map[string]any{
			"SPOTIFY_CLIENT_ID": "test",
			"SPOTIFY_SECRET":    "test",
		}
		mockApp := &config.App{
			Client: client,
			Config: cfg,
		}

		err := Connect(mockApp)
		if err == nil {
			t.Error("expected error on connect")
		}
	})
}

type mockClient struct {
	res []byte
	err error
}

type mockConfig struct {
	mockMap map[string]any
	wrote   bool
}

func (mc *mockClient) GetToken(url string, body interface{}) ([]byte, error) {
	return mc.res, mc.err
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
