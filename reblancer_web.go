package rebalancerweb

import (
	"io"
	"net/http"
)

// HealthCheckHandler returns a HTTP 200 response along with a json object
// indicating liveness.
func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, `{"alive": true}`)
}
