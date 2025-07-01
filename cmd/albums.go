package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/flames31/spotify-client/internal/config"
	"github.com/flames31/spotify-client/internal/spotify"
	"github.com/spf13/cobra"
)

func albumsCmd(app *config.App) *cobra.Command {
	return &cobra.Command{
		Use:   "list-playlists",
		Short: "This command fetches playlists of current user.",
		Long: `
This command fetches playlists from spotify's 3rd party API.
You can set a limit and offset as optional parameters. (default is 10 and 5)`,
		Run: func(cmd *cobra.Command, args []string) {
			albums, err := spotify.GetAlbums(app)
			fmt.Println(albums)
			if err != nil {
				log.Printf("error fetching albums from spotify :%v", err)
				os.Exit(1)
			}
		},
	}
}
