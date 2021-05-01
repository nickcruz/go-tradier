package main

import (
	"fmt"
	"os"
	"time"

	"github.com/nickcruz/go-tradier"
)

func examplePrintQuotes(client *tradier.Client) {
	quotes, err := client.GetQuotes([]string{"GOOG", "SPX"})
	if err != nil {
		panic(err)
	}

	for _, quote := range quotes {
		fmt.Printf("%v: bid $%.02f (%v shares), ask $%.02f (%v shares)\n",
			quote.Symbol, quote.Bid, quote.BidSize, quote.Ask, quote.AskSize)
	}
}

func examplePrintOptionChain(client *tradier.Client) {
	chain, err := client.GetOptionChain("GOOG", time.Date(2021, 05, 07, 0, 0, 0, 0, time.UTC), true)
	if err != nil {
		panic(err)
	}
	fmt.Print(chain[0])
	// for _, quote := range chain {
	// 	fmt.Printf("%v: bid $%.02f (%v shares), ask $%.02f (%v shares)\n",
	// 		quote.Symbol, quote.Bid, quote.BidSize, quote.Ask, quote.AskSize)
	// }
}

func main() {
	tradierApiKey := os.Getenv("TRADIER_API_KEY")
	params := tradier.DefaultParams(tradierApiKey)
	client := tradier.NewClient(params)

	// examplePrintQuotes(client)
	examplePrintOptionChain(client)
}
