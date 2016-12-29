package run

import (
	"os"
	"os/signal"
	"sync"
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
	wg := &sync.WaitGroup{}
	f := func(wg *sync.WaitGroup, closeC chan bool) {
		defer wg.Done()
		defer logger.Debug("run: exit run goroutine")
		pubsub := redis.New(config.NewRedisConfig())
		msgs, err := pubsub.Subscribe("foo")
		if err != nil {
			logger.WithError(err).Fatal("pubsub: subscribe error")
			return
		}
		defer pubsub.Close()
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			defer logger.Debug("run: exit msg read goroutine")
			for msg := range msgs {
				logger.Debug(msg)
			}
		}(wg)
		<-closeC
	}
	closeC := make(chan bool, 1)
	wg.Add(1)
	go f(wg, closeC)
	UntilQuit()
	close(closeC)
	wg.Wait()
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
