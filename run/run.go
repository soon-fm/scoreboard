package run

import (
	"os"
	"os/signal"
	"syscall"

	"scoreboard/config"
	"scoreboard/logger"
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

// Central run function
func Run() {
	ps := redis.New(config.NewRedisConfig())
	msgs, err := ps.Subscribe("foo")
	if err != nil {
		log.WithError(err).Fatal("failed to subscribe")
	}
	go func() {
		for msg := range msgs {
			log.Debug(msg)
		}
	}()
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
