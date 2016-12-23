// Reads in configuration from files

package config

import (
	"os"
	"strings"

	"github.com/spf13/viper"
)

// Setup Defaults
func init() {
	viper.SetConfigType("toml")
	viper.SetConfigName("config")
	viper.AddConfigPath("/etc/scoreboard")
	viper.AddConfigPath("$HOME/.config/scoreboard")
	viper.SetEnvPrefix("scoreboard")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
}

// Reads configuration in from a file. If the passed in path exists
// then the config file is set to be that path, else default paths
// are used to read from
func Read(path string) error {
	if _, err := os.Stat(path); err != nil {
		viper.SetConfigFile(path)
	}
	return viper.ReadInConfig()
}
