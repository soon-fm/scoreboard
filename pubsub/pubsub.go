package pubsub

type Closer interface {
	Close() error
}

type Subscriber interface {
	Subscribe(string) (<-chan string, error)
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
