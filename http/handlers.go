package http

import (
	"encoding/json"
	"net/http"
	"reflect"
	"sort"
	"time"

	"github.com/nvellon/hal"

	"scoreboard/api"
	"scoreboard/db"
	"scoreboard/scores"
)

type Score struct {
	ID          string `json:"id"`
	DisplayName string `json:"displayName"`
	GivenName   string `json:"givenName"`
	FamilyName  string `json:"familyName"`
	Position    int    `json:"position"`
	Score       int64  `json:"score"`
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
	dateStr := r.URL.Query().Get("date")
	date, err := time.Parse(time.RFC3339, dateStr)
	if err != nil {
		date = time.Now().UTC()
	}
	results, err := scores.ByWeek(queryer, date)
	if err != nil {
		log.WithError(err).Error("failed to get scores")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	sort.Sort(sort.Reverse(results)) // Sort the results so they are in order
	sb := ScoreBoard{Scores: make([]Score, len(results))}
	api := ctx.Value("api").(*api.Client)
	for i, result := range results {
		user, err := api.User(result.User)
		if err != nil {
			log.WithError(err).WithField("user", result.User).Error("failed to query api for user")
			continue
		}
		sb.Scores[i] = Score{
			ID:          result.User,
			DisplayName: user.DisplayName,
			GivenName:   user.GivenName,
			FamilyName:  user.FamilyName,
			Position:    i + 1,
			Score:       result.Score,
		}
	}
	j, err := json.Marshal(hal.NewResource(sb, r.URL.String()))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

func ScoresYear(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	queryer, ok := ctx.Value("influxdb").(db.Queryer)
	if !ok {
		log.WithField("type", reflect.TypeOf(ctx.Value("influxdb"))).
			Error("influxdb type on context incorrect")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	dateStr := r.URL.Query().Get("date")
	date, err := time.Parse(time.RFC3339, dateStr)
	if err != nil {
		date = time.Now().UTC()
	}
	results, err := scores.ByYear(queryer, date)
	if err != nil {
		log.WithError(err).Error("failed to get scores")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	sort.Sort(sort.Reverse(results)) // Sort the results so they are in order
	sb := ScoreBoard{Scores: make([]Score, len(results))}
	api := ctx.Value("api").(*api.Client)
	for i, result := range results {
		user, err := api.User(result.User)
		if err != nil {
			log.WithError(err).WithField("user", result.User).Error("failed to query api for user")
			continue
		}
		sb.Scores[i] = Score{
			DisplayName: user.DisplayName,
			ID:          result.User,
			Position:    i + 1,
			Score:       result.Score,
		}
	}
	j, err := json.Marshal(hal.NewResource(sb, r.URL.String()))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}
