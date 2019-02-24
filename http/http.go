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

// NewServer returns a *http.Server ready to serve our routes over http
func NewServer() *http.Server {
	router := http.NewServeMux()
	router.HandleFunc("/healthcheck", HealthCheckHandler)
	return &http.Server{Addr: "localhost:8080", Handler: router}
}
