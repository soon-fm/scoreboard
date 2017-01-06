package pubsub

type Message interface {
	Name() string
	Payload() string
}

type Closer interface {
	Close() error
}

type Subscriber interface {
	Subscribe(string) (Reader, error)
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

type Reader interface {
	Read() <-chan Message
}

type ReadCloser interface {
	Reader
	Closer
}
