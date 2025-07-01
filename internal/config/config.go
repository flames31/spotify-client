package config

type App struct {
	Client HTTPClient
	Config Viper
}

type HTTPClient interface {
	GetToken(url string, body interface{}) ([]byte, error)
}

type Viper interface {
	Get(key string) any
	Set(key string, val any)
	Write() error
}
