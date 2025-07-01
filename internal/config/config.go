package config

import (
	"github.com/go-resty/resty/v2"
	"github.com/spf13/viper"
)

type App struct {
	Client *resty.Client
	Viper  *viper.Viper
}
