package config

import (
	"fmt"

	"github.com/go-resty/resty/v2"
	"github.com/spf13/viper"
)

type restyClient struct{}

func (rc *restyClient) GetToken(url string, body interface{}) ([]byte, error) {
	client := resty.New()
	resp, err := client.R().
		SetHeader("Content-Type", "application/x-www-form-urlencoded").
		SetBody(body).
		Post(url)

	if err != nil {
		return []byte{}, fmt.Errorf("error sending request to spotify : %w", err)
	}

	return resp.Body(), nil
}

type config struct{}

func (c *config) Get(key string) any {
	return viper.Get(key)
}

func (c *config) Set(key string, val any) {
	viper.Set(key, val)
}

func (c *config) Write() error {
	if err := viper.WriteConfigAs(viper.ConfigFileUsed()); err != nil {
		return fmt.Errorf("write config: %w", err)
	}

	return nil
}
