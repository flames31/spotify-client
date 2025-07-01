package main

import (
	"log"

	"github.com/flames31/spotify-client/cmd"
	"github.com/flames31/spotify-client/internal/config"
)

func main() {
	newApp := &config.App{
		Client: config.NewRestyClient(),
		Config: config.NewConfig(),
	}
	root := cmd.RootCmd(newApp)
	if err := root.Execute(); err != nil {
		log.Fatalf("error : %v", err)
	}
}
