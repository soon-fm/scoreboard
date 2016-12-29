// Provides configuration types for Redis Pub/Sub
//
// Example TOML:
// [redis]
// address = "localhost:6379"  # Address of redis server in host:port format
// password = "foo" # Optional, remove or leave blank
// db = 0 # Optional DB number, remove or leave blank
//
// Environment Variables:
// SCOREBOARD_REDIS_ADDRESS = "localhost:6379"
// SCOREBOARD_REDIS_PASSWORD = "foo"
// SCOREBOARD_REDIS_DB = "0"

package config

import "github.com/spf13/viper"

// Set redis default configuration
func init() {
	viper.SetDefault("redis.address", "localhost:6379")
	viper.SetDefault("redis.password", "")
	viper.SetDefault("redis.db", 0)
}

// A simple type for accessing redis configuration
type redis struct{}

// Returns the reids connection address in host:ip format
func (r redis) Address() string {
	viper.BindEnv("REDIS.ADDRESS")
	return viper.GetString("redis.address")
}

// Returns redis server password
func (r redis) Password() string {
	viper.BindEnv("REDIS.PASSWORD")
	return viper.GetString("redis.password")
}

// Returns logging format to use
func (r redis) DB() int {
	viper.BindEnv("REDIS.DB")
	return viper.GetInt("redis.db")
}

// Returns a new redis config
func NewRedisConfig() redis {
	return redis{}
}
