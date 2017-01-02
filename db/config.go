// db package configuration
//
// Example TOML Configuration:
// [influxdb]
// address = "http://localhost:8086" # Required Influx DB HTTP API address
// db = "myDb" # Required DB Name
// username = "username" # Optional Username - omit of not required
// password = "password" # Optional Password - omit of not required
//
// Environment Variables
//
// SCOREBOARD_INFLUXDB_ADDRESS
// SCOREBOARD_INFLUXDB_DB
// SCOREBOARD_INFLUXDB_USERNAME
// SCOREBOARD_INFLUXDB_PASSWORD

package db

import "github.com/spf13/viper"

func init() {
	viper.SetDefault("influxdb.address", "http://localhost:8086")
	viper.SetDefault("influxdb.username", "")
	viper.SetDefault("influxdb.password", "")
	viper.SetDefault("influxdb.db", "")
}

// Accepted by the DB constructor, implemented by Config type below
type Configurer interface {
	Address() string
	Username() string
	Password() string
	DB() string
}

// Empty struct to implement Configurer interface
type Config struct{}

// Returns the influx db http api address
func (c Config) Address() string {
	viper.BindEnv("INFLUXDB.ADDRESS")
	return viper.GetString("influxdb.address")
}

// Returns the influx db user name
func (c Config) Username() string {
	viper.BindEnv("INFLUXDB.USERNAME")
	return viper.GetString("influxdb.username")
}

// Returns the influx db password
func (c Config) Password() string {
	viper.BindEnv("INFLUXDB.PASSWORD")
	return viper.GetString("influxdb.password")
}

// Returns the influx db to use
func (c Config) DB() string {
	viper.BindEnv("INFLUXDB.DB")
	return viper.GetString("influxdb.db")
}

// Config constructor
func NewConfig() Config {
	return Config{}
}
