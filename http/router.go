package http

import (
	"encoding/json"
	"net/http"
	"scoreboard/version"
	"time"

	"github.com/nvellon/hal"
	"github.com/pressly/chi"
)

type HealthCheck struct {
	Version   string    `json:"version"`
	BuildTime time.Time `json:"buildTime"`
}

// Constructs a new HTTP router
func NewRouter() *chi.Mux {
	// Top Level Router
	router := chi.NewRouter()
	// Global Middleware
	router.Use(JSONContentType)
	// Healthcheck Handlers
	router.Get("/__healthcheck__", func(w http.ResponseWriter, r *http.Request) {
		v, _ := version.Version()
		bt, _ := version.BuildTime()
		j, err := json.Marshal(hal.NewResource(HealthCheck{v, bt}, r.URL.String()))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(j)
	})
	// Mount Other Routers
	router.Mount("/scores", ScoresRouter())
	return router
}

// Routes for /scores/...
func ScoresRouter() *chi.Mux {
	router := chi.NewRouter()
	// Middleare for this route
	router.Use(InfluxDBMiddleware)
	router.Use(APIMiddleware)
	// Handlers
	router.Get("/week", ScoresWeek)
	return router
}
