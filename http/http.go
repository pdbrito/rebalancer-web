package rebalancerweb

import (
	"encoding/json"
	"github.com/pdbrito/rebalancer"
	"io"
	"net/http"
)

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, `{"alive":true}`)
}

func pricelistHandler(pricelist rebalancer.Pricelist) http.HandlerFunc {
	b, _ := json.Marshal(pricelist)
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		io.WriteString(w, string(b))
	}
}

// NewServer returns a *http.Server ready to serve our routes over http
func NewServer(pricelist rebalancer.Pricelist) *http.Server {
	router := http.NewServeMux()
	router.HandleFunc("/healthcheck", healthCheckHandler)
	router.HandleFunc("/pricelist", pricelistHandler(pricelist))
	return &http.Server{Addr: "localhost:8080", Handler: router}
}
