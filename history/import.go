package history

import (
	"encoding/json"
	"io/ioutil"
	"time"

	"scoreboard/db"
	"scoreboard/logger"
)

var log = logger.WithField("pkg", "scoreboard/history")

type Records struct {
	Records []Record `json:"RECORDS"`
}

type Record struct {
	Created string `json:"created"`
	Updates string `json:"updated"`
	ID      string `json:"id"`
	User    string `json:"user_id"`
	Track   string `json:"track_id"`
}

// Open a json dump of a playlist history db and import
// the data into the scores databases
func Import(path string) error {
	raw, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	records := &Records{}
	if err := json.Unmarshal(raw, records); err != nil {
		return err
	}
	return db.With(func(sess *db.DB) error {
		for _, record := range records.Records {
			t, err := time.Parse("01/02/2006 15:04:05", record.Created)
			if err != nil {
				return err
			}
			if err := sess.Write(score); err != nil {
				return err
			}
		}
		return nil
	})
}
