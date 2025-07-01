package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func NewViperConfig() *viperConfig {
	return &viperConfig{}
}

type Config interface {
	Get(key string) any
	Set(key string, val any)
	Write() error
}

type viperConfig struct{}

func (c *viperConfig) Get(key string) any {
	return viper.Get(key)
}

func (c *viperConfig) Set(key string, val any) {
	viper.Set(key, val)
}

func (c *viperConfig) Write() error {
	if err := viper.WriteConfigAs(viper.ConfigFileUsed()); err != nil {
		return fmt.Errorf("write config: %w", err)
	}

	return nil
}
