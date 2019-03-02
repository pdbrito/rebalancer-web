package rebalancerweb

import (
	"encoding/json"
	"github.com/pdbrito/rebalancer"
	"io"
	"net/http"
)

// HealthCheckHandler returns a HTTP 200 response along with a json object
// indicating liveness.
func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, `{"alive":true}`)
}

// PricelistHandler returns a handler for the pricelist endpoint
func PricelistHandler(pricelist rebalancer.Pricelist) http.HandlerFunc {
	b, _ := json.Marshal(pricelist)
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		io.WriteString(w, string(b))
	}
}

// NewServer returns a *http.Server ready to serve our routes over http
func NewServer(pricelist rebalancer.Pricelist) *http.Server {
	router := http.NewServeMux()
	router.HandleFunc("/healthcheck", HealthCheckHandler)
	router.HandleFunc("/pricelist", PricelistHandler(pricelist))
	return &http.Server{Addr: "localhost:8080", Handler: router}
}
