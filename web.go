package rebalancerweb

import (
	"io"
	"net/http"
)

// HealthCheckHandler returns a HTTP 200 response along with a json object
// indicating liveness.
func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, `{"alive":true}`)
}

// StartServer will pass HTTP requests made on localhost:8080 through our router
func StartServer() error {
	router := http.NewServeMux()
	router.HandleFunc("/healthcheck", HealthCheckHandler)
	s := &http.Server{Addr: "localhost:8080", Handler: router}
	return s.ListenAndServe()
}
