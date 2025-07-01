package cmd

import (
	"fmt"
	"os"

	"github.com/flames31/spotify-client/internal/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

func RootCmd(app *config.App) *cobra.Command {
	command := &cobra.Command{
		Use:   "spotify-client",
		Short: "Your own Spotify client to interact with Spotify!",
		Long:  `Spoify client is a CLI app that fetches data from your spotify account and provides various commands to view and edit data`,
	}

	command.AddCommand(connectCmd(app))
	return command
}

func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		cwd, err := os.Getwd()
		cobra.CheckErr(err)

		viper.AddConfigPath(cwd)
		viper.SetConfigType("json")
		viper.SetConfigName(".spotify-client")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stdout, "Using config file:", viper.ConfigFileUsed())
	}
}
