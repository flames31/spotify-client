package spotify

import (
	"errors"
	"fmt"
	"os"

	"github.com/flames31/spotify-client/internal/config"
)

func RequestToken(auth *config.Auth) error {
	clientID, clientSecret := auth.Client.GetCredentials()
	if clientID == "" || clientSecret == "" {
		return errors.New("clientId and/or clientSecret not set")
	}
	body := fmt.Sprintf("grant_type=client_credentials&client_id=%v&client_secret=%v", clientID, clientSecret)
	token, err := auth.Client.GetToken("/api/token", body)
	if err != nil {
		fmt.Printf("error fetching token: %v", err)
		os.Exit(1)
	}

	auth.Config.Set("SPOTIFY_TOKEN", token)
	if err := auth.Config.Write(); err != nil {
		fmt.Println("write config: %w", err)
		os.Exit(1)
	}

	return nil
}
