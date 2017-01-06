// Package that allows connections to a Redis Pub/Sub service

package redis

import (
	"scoreboard/logger"
	"scoreboard/pubsub"
	"sync"

	redis "gopkg.in/redis.v5"
)

// Package logger
var log = logger.WithField("pkg", "scoreboard/pubsub/redis")

// Config interface, please see config package for more details
type Configurer interface {
	Address() string
	Password() string
	DB() int
}

// Implements the pubsub.Message interface
type Message struct {
	msg *redis.Message
}

// Returns the channel name the message originated from
func (m *Message) Name() string {
	return m.msg.Channel
}

// Returns the raw message payloaad
func (m *Message) Payload() string {
	return m.msg.Payload
}

// Pubsub subscription implementing the pubsub.ReadCloser interface
type Subscription struct {
	pubsub *redis.PubSub
	wg     *sync.WaitGroup
	ch     chan pubsub.Message
}

// Returns a sync.Waitgroup pointer
func (s *Subscription) waitGroup() *sync.WaitGroup {
	if s.wg == nil {
		s.wg = &sync.WaitGroup{}
	}
	return s.wg
}

// Returns Message channel
func (s *Subscription) channel() chan pubsub.Message {
	if s.ch == nil {
		s.ch = make(chan pubsub.Message)
	}
	return s.ch
}

// Consumes messages from pubsub placing them on a message channel
func (s *Subscription) receive() {
	log.Debug("start subscription receive")
	defer log.Debug("exit subscription receive")
	wg := s.waitGroup()
	wg.Add(1)
	defer wg.Done()
	ch := s.channel()
	for {
		msg, err := s.pubsub.ReceiveMessage()
		if err != nil {
			log.WithError(err).Warn("recieve message error")
			return // Exit routine on error
		}
		ch <- &Message{msg}
	}
}

// Read messages from the pubsub connection
func (s *Subscription) Read() <-chan pubsub.Message {
	ch := s.channel()
	go s.receive()
	return (<-chan pubsub.Message)(ch)
}

// Close pubsub connection
func (s *Subscription) Close() error {
	// Close pubsub subscription
	if s.pubsub != nil {
		s.pubsub.Close()
	}
	// Wait for pubsub routines to exit
	wg := s.waitGroup()
	wg.Wait()
	// Close the message channel
	if s.ch != nil {
		close(s.ch)
		s.ch = nil
	}
	return nil
}

// Redis client implementing the pubsub.SubscribeCloser interface
type Client struct {
	config        Configurer
	client        *redis.Client
	subscriptions map[string]*Subscription
}

// Subscribes to a topic, opening a pubsub connection, returning a pubsub.Reader
// for consuming messages on the topic
func (c *Client) Subscribe(topic string) (pubsub.Reader, error) {
	log.WithField("topic", topic).Info("subscribe to topic")
	pubsub, err := c.client.Subscribe(topic)
	if err != nil {
		return nil, err
	}
	subscription := &Subscription{pubsub: pubsub}
	c.subscriptions[topic] = subscription
	return subscription, nil
}

// Close redis connection
func (c *Client) Close() error {
	defer logger.Info("closed redis pubsub client")
	// Close running subscriptions
	for _, subscription := range c.subscriptions {
		if err := subscription.Close(); err != nil {
			log.WithError(err).Error("failed to close subscription")
		}
	}
	// Close redis client
	if c.client != nil {
		return c.client.Close()
	}
	return nil
}

// Construct a new redis client
func New(config Configurer) *Client {
	return &Client{
		config: config,
		client: redis.NewClient(&redis.Options{
			Addr:     config.Address(),
			Password: config.Password(),
			DB:       config.DB(),
		}),
		subscriptions: make(map[string]*Subscription),
	}
}
