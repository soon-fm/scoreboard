package pubsub

type Message interface {
	Topic() string
	Payload() string
}

type Closer interface {
	Close()
}

type Subscriber interface {
	Subscribe(...string) (<-chan Message, error)
}

type SubscribeCloser interface {
	Subscriber
	Closer
}

type Publisher interface {
	Publish(string, []byte) error
}

type PublishCloser interface {
	Publisher
	Closer
}
