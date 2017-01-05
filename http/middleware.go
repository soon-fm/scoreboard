package http

import (
	"context"
	"net/http"
	"scoreboard/db"
)

func JSONContentType(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json+hal")
		next.ServeHTTP(w, r)
	})
}

func InfluxDBMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		client, err := db.New(db.NewConfig())
		if err != nil {
			log.WithError(err).Error("influxdb middleware error")
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		defer client.Close()
		ctx := context.WithValue(r.Context(), "influxdb", client)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
