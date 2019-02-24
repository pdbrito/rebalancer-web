package main

import (
	"github.com/pdbrito/rebalancer-web/http"
	"log"
)

func main() {
	server := rebalancerweb.NewServer()
	log.Fatal(server.ListenAndServe())
}
