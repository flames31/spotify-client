package cmd

import (
	"fmt"

	"github.com/flames31/spotify-client/internal/config"
	"github.com/flames31/spotify-client/internal/spotify"
	"github.com/spf13/cobra"
)

func connectCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "connect",
		Short: "This command connects to spotify to retrieve an auth token.",
		Long: `
This command sends a request to spotify's 3rd party API.
This will redirect you to a link where you can authorize the use of spotifys API.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth := &config.Auth{
				Client: config.NewSpotifyAuthClient(),
				Config: config.NewViperConfig(),
			}
			spotify.RequestToken(auth)

			fmt.Println("Succesfully fetched token!")
		},
	}
}
