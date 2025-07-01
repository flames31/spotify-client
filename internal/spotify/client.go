package spotify

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/flames31/spotify-client/internal/config"
)

const SPOTIFY_API_URL = "https://accounts.spotify.com/api/token"

func Connect(app *config.App) error {
	clientID := app.Config.Get("SPOTIFY_CLIENT_ID")
	clientSecret := app.Config.Get("SPOTIFY_SECRET")

	body := fmt.Sprintf("grant_type=client_credentials&client_id=%v&client_secret=%v", clientID, clientSecret)
	resp, err := app.Client.GetToken(SPOTIFY_API_URL, body)
	if err != nil {
		return fmt.Errorf("error getting token from spotify :%w", err)
	}

	var resBody map[string]interface{}
	if err := json.Unmarshal(resp, &resBody); err != nil {
		return fmt.Errorf("error unmarshalling response from spotify :%w", err)
	}

	spotifyToken := resBody["access_token"].(string)
	if spotifyToken == "" {
		return errors.New("access_token not received from spotify")
	}

	app.Config.Set("SPOTIFY_TOKEN", spotifyToken)
	err = app.Config.Write()
	if err != nil {
		return fmt.Errorf("write config: %w", err)
	}

	fmt.Println("Succesfully fetched spotify token!")

	return nil
}
