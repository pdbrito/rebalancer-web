package main

import (
	"github.com/pdbrito/rebalancer"
	"github.com/pdbrito/rebalancer-web/http"
	"github.com/shopspring/decimal"
	"log"
)

func main() {
	pricelist := rebalancer.Pricelist{
		"ETH": decimal.NewFromFloat(200),
		"BTC": decimal.NewFromFloat(5000),
	}
	server := rebalancerweb.NewServer(pricelist)
	log.Fatal(server.ListenAndServe())
}
