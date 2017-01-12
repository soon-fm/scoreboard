package api

import "github.com/spf13/viper"

const (
	viper_host_key   = "api.host"
	viper_scheme_key = "api.scheme"
)

func init() {
	// API Host
	viper.BindEnv(viper_host_key)
	viper.SetDefault(viper_host_key, "localhost:80")
	// API Scheme
	viper.BindEnv(viper_scheme_key)
	viper.SetDefault(viper_scheme_key, "http")
}

type Configurer interface {
	Host() string
	Scheme() string
}

type Config struct{}

func (c Config) Host() string {
	return viper.GetString(viper_host_key)
}

func (c Config) Scheme() string {
	return viper.GetString(viper_scheme_key)
}

func NewConfig() Config {
	return Config{}
}
