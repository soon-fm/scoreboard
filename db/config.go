// db package configuration
//
// Example TOML Configuration:
// [db]
// host = "localhost" # Datbabase host
// port = 5432 # Datbabase port
// db = "myDb" # Required DB Name
// username = "username" # Optional Username - omit of not required
// password = "password" # Optional Password - omit of not required
// migration_path = "/path/to/migrations" # Optional Datbase migration path
//
// Environment Variables
//
// SCOREBOARD_DB_HOST
// SCOREBOARD_DB_PORT
// SCOREBOARD_DB_DB
// SCOREBOARD_DB_USERNAME
// SCOREBOARD_DB_PASSWORD
// SCOREBOARD_DB_MIGRATION_PATH

package db

import (
	"fmt"

	"github.com/spf13/viper"
)

func init() {
	viper.SetDefault("db.host", "localhost")
	viper.SetDefault("db.port", 5432)
	viper.SetDefault("db.username", "postgres")
	viper.SetDefault("db.password", "postgres")
	viper.SetDefault("db.db", "scoreboard")
	viper.SetDefault("db.migration_path", "./migrations")
}

// Accepted by the DB constructor, implemented by Config type below
type Configurer interface {
	Host() string
	Port() int
	Username() string
	Password() string
	DB() string
	ConnectionURL() string
	MigrationPath() string
}

// Empty struct to implement Configurer interface
type Config struct{}

// Returns the db host
func (c Config) Host() string {
	viper.BindEnv("DB.HOST")
	return viper.GetString("db.host")
}

// Returns the db port
func (c Config) Port() int {
	viper.BindEnv("DB.PORT")
	return viper.GetInt("db.port")
}

// Returns the db user name
func (c Config) Username() string {
	viper.BindEnv("DB.USERNAME")
	return viper.GetString("db.username")
}

// Returns the db password
func (c Config) Password() string {
	viper.BindEnv("DB.PASSWORD")
	return viper.GetString("db.password")
}

// Returns the db to use
func (c Config) DB() string {
	viper.BindEnv("DB.DB")
	return viper.GetString("db.db")
}

// Returns the db migration path
func (c Config) MigrationPath() string {
	viper.BindEnv("DB.MIGRATION_PATH")
	return viper.GetString("db.migration_path")
}

// Returns a connection url
func (c Config) ConnectionURL() string {
	return fmt.Sprintf(
		"host=%s port=%d dbname=%s user=%s password=%s sslmode=disable",
		c.Host(),
		c.Port(),
		c.DB(),
		c.Username(),
		c.Password())
}

// Config constructor
func NewConfig() Config {
	return Config{}
}
