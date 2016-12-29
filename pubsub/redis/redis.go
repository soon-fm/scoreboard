// Package that allows connections to a Redis Pub/Sub service
// Please see "scoreboard/config" package for configuration options

package redis

import (
	"scoreboard/logger"

	redis "gopkg.in/redis.v5"
)

// Config interface, please see config package for more details
type Config interface {
	Address() string
	Password() string
	DB() int
}

type PubSub struct {
	config Config
	client *redis.Client
	pubsub *redis.PubSub
	msgs   chan string
}

func (p *PubSub) subscribe(channels ...string) error {
	pubsub, err := p.client.PSubscribe(channels...)
	if err != nil {
		return err
	}
	p.pubsub = pubsub
	p.msgs = make(chan string)
	return nil
}

func (p *PubSub) Subscribe(channels ...string) (<-chan string, error) {
	if err := p.subscribe(channels...); err != nil {
		return nil, err
	}
	go func() {
		defer logger.Debug("pubsub: exit subscribe goroutine")
		defer close(p.msgs)
		for {
			msg, err := p.pubsub.ReceiveMessage()
			if err != nil {
				logger.WithError(err).Warn("pubsub: recieve message error")
				return
			}
			p.msgs <- msg.Payload
		}
	}()
	return (<-chan string)(p.msgs), nil
}

func (p *PubSub) Publish(ch string, msg []byte) error {
	return nil
}

func (p *PubSub) Close() {
	defer logger.Debug("pubsub: close subscriptions")
	if p.pubsub != nil {
		if err := p.pubsub.Close(); err != nil {
			logger.WithError(err).Error("pubsub: close error")
		}
	}
}

func New(config Config) *PubSub {
	return &PubSub{
		config: config,
		client: redis.NewClient(&redis.Options{
			Addr:     config.Address(),
			Password: config.Password(),
			DB:       config.DB(),
		}),
	}
}
