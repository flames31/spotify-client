package config

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/go-resty/resty/v2"
	"github.com/spf13/viper"
)

type Auth struct {
	Client AuthClient
	Config Config
}

type AuthClient interface {
	GetToken(url string, body interface{}) (string, error)
	GetCredentials() (string, string)
}

type SpotifyAuthClient struct {
	Client                 *resty.Client
	ClientID, ClientSecret string
}

func (s *SpotifyAuthClient) GetToken(url string, body interface{}) (string, error) {
	resp, err := s.Client.R().
		SetHeader("Content-Type", "application/x-www-form-urlencoded").
		SetBody(body).
		Post(url)

	if err != nil {
		return "", fmt.Errorf("error sending request to spotify : %w", err)
	}

	var resBody map[string]interface{}
	if err := json.Unmarshal(resp.Body(), &resBody); err != nil {
		return "", fmt.Errorf("error unmarshalling response from spotify :%w", err)
	}

	spotifyToken := resBody["access_token"].(string)
	if spotifyToken == "" {
		return "", errors.New("access_token not received from spotify")
	}

	return spotifyToken, nil
}

func (s *SpotifyAuthClient) GetCredentials() (string, string) {
	return s.ClientID, s.ClientSecret
}

func NewSpotifyAuthClient() *SpotifyAuthClient {
	clientID := viper.GetString("SPOTIFY_CLIENT_ID")
	clientSecret := viper.GetString("SPOTIFY_SECRET")

	return &SpotifyAuthClient{
		Client:       resty.New().SetBaseURL("https://accounts.spotify.com"),
		ClientID:     clientID,
		ClientSecret: clientSecret,
	}
}
