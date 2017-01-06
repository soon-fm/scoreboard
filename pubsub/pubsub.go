package pubsub

type Topic interface {
	Name() string
	Handler() HandlerFunc
}

type Message interface {
	Name() string
	Payload() string
}

type Closer interface {
	Close() error
}

type Subscriber interface {
	Subscribe(Topic) (Reader, error)
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
	Read()
}

type ReadCloser interface {
	Reader
	Closer
}

type HandlerFunc func(Message) error
