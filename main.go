package main

import (
	"github.com/pdbrito/rebalancer"
	"github.com/pdbrito/rebalancer-web/http"
	"github.com/shopspring/decimal"
	"log"
)

func main() {
	pricelister := func() rebalancer.Pricelist {
		return rebalancer.Pricelist{
			"ETH": decimal.NewFromFloat(200),
			"BTC": decimal.NewFromFloat(5000),
		}
	}
	server := rebalancerweb.NewServer(pricelister)
	log.Fatal(server.ListenAndServe())
}
