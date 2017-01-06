package pubsub

type PlayerEvents struct{}

func (p PlayerEvents) Name() string { return "fm:events" }

func (p PlayerEvents) Handler() HandlerFunc {
	return func(msg Message) error {
		log.WithField("payload", msg.Payload()).Debug("recieved player event")
		return nil
	}
}
