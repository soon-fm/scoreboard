package http

import (
	"encoding/json"
	"net/http"
	"reflect"
	"sort"

	"github.com/nvellon/hal"

	"scoreboard/db"
	"scoreboard/scores"
)

type Score struct {
	User  string `json:"user"`
	Score int64  `json:"score"`
}

type ScoreBoard struct {
	Scores []Score `json:"scores"`
}

func ScoresWeek(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	queryer, ok := ctx.Value("influxdb").(db.Queryer)
	if !ok {
		log.WithField("type", reflect.TypeOf(ctx.Value("influxdb"))).
			Error("influxdb type on context incorrect")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	results, err := scores.ThisWeek(queryer)
	if err != nil {
		log.WithError(err).Error("failed to get scores")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	sort.Sort(results) // Sort the results so they are in order
	sb := ScoreBoard{Scores: make([]Score, len(results))}
	for i, result := range results {
		sb.Scores[i] = Score{result.User, result.Score}
	}
	j, err := json.Marshal(hal.NewResource(sb, r.URL.String()))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}
