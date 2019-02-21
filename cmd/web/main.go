package main

import (
	"github.com/pdbrito/rebalancer-web"
	"log"
)

func main() {
	log.Fatal(rebalancerweb.StartServer())
}
