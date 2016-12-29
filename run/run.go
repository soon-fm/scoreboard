package run

import (
	"os"
	"os/signal"
	"syscall"

	"scoreboard/config"
	"scoreboard/logger"
	"scoreboard/pubsub/redis"
)

// Application exit OS signals
var quitSignals = []os.Signal{
	syscall.SIGINT,
	syscall.SIGTERM,
	syscall.SIGQUIT,
}

// Central run function
func Run() {
	ps := redis.New(config.NewRedisConfig())
	defer ps.Close()
	UntilQuit()
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
