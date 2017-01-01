// Package that allows connections to a Redis Pub/Sub service
// Please see "scoreboard/config" package for configuration options

package redis

import (
	"scoreboard/logger"
	"scoreboard/pubsub"

	redis "gopkg.in/redis.v5"
)

// Package logger
var log = logger.WithField("pkg", "pubsub/redis")

// Config interface, please see config package for more details
type Configurer interface {
	Address() string
	Password() string
	DB() int
}

type Message struct {
	msg *redis.Message
}

func (m *Message) Topic() string {
	return m.msg.Channel
}

func (m *Message) Payload() string {
	return m.msg.Payload
}

type PubSub struct {
	config Configurer
	client *redis.Client
	pubsub *redis.PubSub
}

func (p *PubSub) subscribe(channels ...string) error {
	pubsub, err := p.client.PSubscribe(channels...)
	if err != nil {
		return err
	}
	p.pubsub = pubsub
	return nil
}

func (p *PubSub) receiveMessages(ch chan pubsub.Message) {
	log.Debug("start pub/sub receive")
	defer log.Debug("exit pub/sub recieve")
	defer close(ch) // Close the channel on exit
	for {
		msg, err := p.pubsub.ReceiveMessage()
		if err != nil {
			log.WithError(err).Warn("recieve message error")
			return
		}
		ch <- &Message{msg}
	}
}

func (p *PubSub) Subscribe(channels ...string) (<-chan pubsub.Message, error) {
	log.WithField("channels", channels).Debug("pub/sub subscribe")
	defer log.Debug("exit pub/sub subscribe")
	ch := make(chan pubsub.Message)
	if err := p.subscribe(channels...); err != nil {
		return nil, err
	}
	go p.receiveMessages(ch)
	return (<-chan pubsub.Message)(ch), nil
}

func (p *PubSub) Publish(ch string, msg []byte) error {
	return nil
}

func (p *PubSub) Close() {
	defer log.Debug("close subscriptions")
	if p.pubsub != nil {
		if err := p.pubsub.Close(); err != nil {
			log.WithError(err).Error("close error")
		}
	}
}

func New(config Configurer) *PubSub {
	return &PubSub{
		config: config,
		client: redis.NewClient(&redis.Options{
			Addr:     config.Address(),
			Password: config.Password(),
			DB:       config.DB(),
		}),
	}
}
