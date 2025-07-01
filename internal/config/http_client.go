package config

import (
	"github.com/go-resty/resty/v2"
)

func NewRestyClient() *restyClient {
	return &restyClient{
		Client: resty.New().SetBaseURL("https://api.spotify.com/v1"),
	}
}

type HTTPClient interface {
	R() *resty.Request
}

type restyClient struct {
	Client *resty.Client
}

func (rc *restyClient) R() *resty.Request {
	return rc.Client.R()
}
