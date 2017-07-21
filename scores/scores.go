package scores

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/influxdata/influxdb/models"

	"scoreboard/db"
)

// Stores a user id and a score value
type Score struct {
	User  string
	Score int64
}

// Holds a slice of scores
type Scores []Score

// Implement Sorter interface
func (s Scores) Len() int           { return len(s) }
func (s Scores) Swap(a, b int)      { s[a], s[b] = s[b], s[a] }
func (s Scores) Less(a, b int) bool { return s[a].Score < s[b].Score }

func processScores(rows []models.Row) Scores {
	scores := make([]Score, len(rows))
	for x, row := range rows {
		user, ok := row.Tags["user"]
		if !ok {
			continue
		}
		for i, col := range row.Columns {
			switch col {
			case "total":
				value, err := row.Values[0][i].(json.Number).Int64() // TODO: make better
				if err == nil {
					scores[x] = Score{user, value}
				}
			}
		}
	}
	return scores
}

// Returns the summed user scores for the week
func ByWeek(queryer db.Queryer, date time.Time) (Scores, error) {
	result, err := db.ScoresByWeek(queryer, date)
	if err != nil {
		return nil, err
	}
	if len(result) == 0 {
		return nil, errors.New("no results returned")
	}
	return processScores(result[0].Series), nil
}

func ByYear(queryer db.Queryer, date time.Time) (Scores, error) {
	result, err := db.ScoresByYear(queryer, date)
	if err != nil {
		return nil, err
	}
	if len(result) == 0 {
		return nil, errors.New("no results returned")
	}
	return processScores(result[0].Series), nil
}
