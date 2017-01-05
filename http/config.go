// Package configuration
//
// Example TOML:
// [http]
// listen = "0.0.0.0:5000"
//
// Environment Variables:
// SCOREBOARD_HTTP_LISTEN="0.0.0.0:5000"

package http

import "github.com/spf13/viper"

// Set sane defaults and bind environment vars
func init() {
	viper.BindEnv("http.listen")
	viper.SetDefault("http.listen", ":5000")
}

// Configurer inteface config types need to implement
type Configurer interface {
	Listen() string
}

// Empty type that implements Configurer interface
type Config struct{}

// Returns configured listen address
func (c Config) Listen() string {
	return viper.GetString("http.listen")
}

// Config Constructor
func NewConfig() Config {
	return Config{}
}
