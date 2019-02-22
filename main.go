package main

import (
	"github.com/pdbrito/rebalancer-web/http"
	"log"
)

func main() {
	log.Fatal(rebalancerweb.StartServer())
}
