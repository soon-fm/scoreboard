package run

import (
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"

	"scoreboard/db"
	"scoreboard/http"
	"scoreboard/logger"
	"scoreboard/pubsub"
	"scoreboard/pubsub/redis"
)

// Run package logger
var log = logger.WithField("pkg", "run")

// Application exit OS signals
var quitSignals = []os.Signal{
	syscall.SIGINT,
	syscall.SIGTERM,
	syscall.SIGQUIT,
}

// Configuration defaults
func init() {
	viper.SetConfigType("toml")
	viper.SetConfigName("config")
	viper.AddConfigPath("/etc/scoreboard")
	viper.AddConfigPath("$HOME/.config/scoreboard")
	viper.SetEnvPrefix("scoreboard")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
}

// Central run function
func Run() error {
	// Read configuration
	if err := ReadConfig(); err != nil {
		log.WithError(err).Warn("error loading configuration")
	}
	// Setup Logger
	logger.Setup(logger.NewConfig())
	// Ensure our InfluxDB Exists
	if err := db.Create(db.NewConfig()); err != nil {
		return err
	}
	// Redis Pub/Sub Events
	ps := redis.New(redis.NewConfig())
	fmevents, err := ps.Subscribe(pubsub.PlayerEvents{})
	if err != nil {
		return err
	}
	fmevents.Read()
	defer ps.Close()
	// HTTP API Server
	go http.ListenAndServe(http.NewConfig())
	// Run until a quit os signal is received
	UntilQuit()
	return nil
}

// Binds config path loading flag to viper
func BindConfigPathFlag(flag *pflag.Flag) {
	viper.BindPFlag("config.path", flag)
}

// Load configuration from file
func ReadConfig() error {
	path := viper.GetString("config.path")
	if _, err := os.Stat(path); err != nil {
		viper.SetConfigFile(path)
	}
	return viper.ReadInConfig()
}

// Run this method until the passed in os.Signals are triggered
// Returns the recieved signal
func UntilSignal(signals ...os.Signal) os.Signal {
	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, signals...)
	sig := <-sigc // Blocking
	logger.WithField("signal", sig).Debug("recieved signal")
	return sig
}

// Run until a quit signal is recieved
func UntilQuit() os.Signal {
	return UntilSignal(quitSignals...)
}
