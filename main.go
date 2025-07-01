package main

import (
	"log"

	"github.com/flames31/spotify-client/cmd"
	"github.com/flames31/spotify-client/internal/config"
)

func main() {
	app := &config.App{
		Client: config.NewRestyClient(),
		Config: config.NewViperConfig(),
	}
	root := cmd.RootCmd(app)
	if err := root.Execute(); err != nil {
		log.Fatalf("error : %v", err)
	}
}
