package pubsub

import (
	"encoding/json"
	"errors"
	"scoreboard/db"
	"scoreboard/logger"
	"time"
)

type PlayerEvents struct{}

func (p PlayerEvents) Name() string { return "fm:events" }

func (p PlayerEvents) Handler() HandlerFunc {
	return func(msg Message) error {
		log.WithField("payload", msg.Payload()).Debug("recieved player event")
		var payload map[string]interface{}
		if err := json.Unmarshal([]byte(msg.Payload()), &payload); err != nil {
			return err
		}
		raw, ok := payload["event"]
		if !ok {
			// No type in the event payload
			return nil
		}
		event, ok := raw.(string)
		if !ok {
			// "type" is not a string
			return nil
		}
		switch event {
		case "play":
			return p.play(payload)
		case "stop":
			return p.stop(payload)
		}
		return nil
	}
}

func (p PlayerEvents) userFromPayload(payload map[string]interface{}) (string, error) {
	raw, ok := payload["user"]
	if !ok {
		return "", errors.New("user not in payload")
	}
	val, ok := raw.(string)
	if !ok {
		return "", errors.New("user not a string")
	}
	return val, nil
}

// Handle player play event
func (p PlayerEvents) play(payload map[string]interface{}) error {
	log.Debug("handle play event")
	user, err := p.userFromPayload(payload)
	if err != nil {
		return err
	}
	score := db.NewScore(user, 1, time.Now().UTC())
	return db.With(func(db *db.DB) error {
		log.WithFields(logger.F{
			"user":  score.User(),
			"value": score.Value(),
		}).Info("add score for user")
		return db.Write(score)
	})
}

// Handle player stop event
func (p PlayerEvents) stop(payload map[string]interface{}) error {
	log.Debug("handle stop event")
	user, err := p.userFromPayload(payload)
	if err != nil {
		return err
	}
	score := db.NewScore(user, -1, time.Now().UTC())
	return db.With(func(db *db.DB) error {
		log.WithFields(logger.F{
			"user":  score.User(),
			"value": score.Value(),
		}).Info("add score for user")
		return db.Write(score)
	})
}
