package spotify

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/flames31/spotify-client/internal/config"
	"github.com/spf13/viper"
)

const SPOTIFY_API_URL = "https://accounts.spotify.com/api/token"

func Connect(app *config.App) error {
	clientID := app.Viper.Get("SPOTIFY_CLIENT_ID")
	clientSecret := app.Viper.Get("SPOTIFY_SECRET")

	resp, err := app.Client.R().
		SetHeader("Content-Type", "application/x-www-form-urlencoded").
		SetBody(fmt.Sprintf("grant_type=client_credentials&client_id=%v&client_secret=%v", clientID, clientSecret)).
		Post(SPOTIFY_API_URL)

	if err != nil {
		return fmt.Errorf("error sending request to spotify :%w", err)
	}
	var resBody map[string]interface{}
	if err := json.Unmarshal(resp.Body(), &resBody); err != nil {
		return fmt.Errorf("error unmarshalling response from spotify :%w", err)
	}

	spotifyToken := resBody["access_token"].(string)
	if spotifyToken == "" {
		return errors.New("access_token not received from spotify")
	}

	viper.Set("SPOTIFY_TOKEN", spotifyToken)
	if err := viper.WriteConfigAs(viper.ConfigFileUsed()); err != nil {
		return fmt.Errorf("write config: %w", err)
	}

	fmt.Println("Succesfully fetched spotify token!")

	return nil
}
