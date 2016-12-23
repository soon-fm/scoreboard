// Logger Configuration
//
// Example TOML:
// [log]
// level = "debug"
// logfile = "/path/to/scoreboard.log
// format = "logstash"
//
// [logstash]
// type = "mylogstashtype"
//
// Environment Variables:
// SCOREBOARD_LOG_LEVEL = "info"
// SCOREBOARD_LOG_LOGFILE = "/path/to/scoreboard.log"
// SCOREBOARD_LOG_FORMAT = "logstash"
//
// CLI Flags:
// -l/--level info

package config

import (
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

// Set logging configuration defaults
func init() {
	viper.SetDefault("log.level", "info")
	viper.SetDefault("log.format", "text")
	viper.SetDefault("log.logfile", "")
	viper.SetDefault("log.console_output", true)
	viper.SetDefault("logstash.type", "")
}

// Allows us to bind a cli flag to a viper config option for log.level
func BindLogLevelFlag(flag *pflag.Flag) {
	viper.BindPFlag("log.level", flag)
}

// A simple type for accessing logging configuration
type log struct{}

// Returns the logging verbosity level, binds to environment variable
func (l log) Level() string {
	viper.BindEnv("LOG.LEVEL")
	return viper.GetString("log.level")
}

// Returns absolute path to logfile
func (l log) LogFile() string {
	viper.BindEnv("LOG.LOGFILE")
	return viper.GetString("log.logfile")
}

// Returns logging format to use
func (l log) Format() string {
	viper.BindEnv("LOG.FORMAT")
	return viper.GetString("log.format")
}

// Returns console log output bool
func (l log) ConsoleOutput() bool {
	return viper.GetBool("log.console_output")
}

// Returns logstash format type
func (l log) LogstashType() string {
	return viper.GetString("logstash.type")
}

// Consturcts a new Log config
func NewLogConifg() log {
	return log{}
}
