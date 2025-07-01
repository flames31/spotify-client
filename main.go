package main

import (
	"log"

	"github.com/flames31/spotify-client/cmd"
	"github.com/flames31/spotify-client/internal/config"
	"github.com/go-resty/resty/v2"
	"github.com/spf13/viper"
)

func main() {
	newApp := &config.App{
		Client: resty.New(),
		Viper:  viper.GetViper(),
	}
	root := cmd.RootCmd(newApp)
	if err := root.Execute(); err != nil {
		log.Fatalf("error : %v", err)
	}
}
