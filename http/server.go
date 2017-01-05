package http

import "net/http"

func ListenAndServe(config Configurer) {
	log.WithField("listen", config.Listen()).Info("starting http server")
	http.ListenAndServe(config.Listen(), NewRouter())
}
