package main

import (
	"fmt"
	"time"
)

func Printer(data *Data, quit chan string) {
	var btcUSD, btcEUR, eurUSD float64 = 0, 0, 0
	var activeBTC, activeCUR int = 0, 0

	printTicker := func() {
		tm := time.Now().Format(time.RFC3339)
		fmt.Printf(
			"\rLast Update: %v BTC/USD: %f   EUR/USD: %f   BTC/EUR: %f Active sources: BTC/USD (%d of %d)  EUR/USD (%d of %d)",
			tm, btcUSD, eurUSD, btcEUR, activeBTC, data.totalBTC, activeCUR, data.totalCUR,
		)
	}

	printTicker()

	for {
		select {
		case btcUSD = <-data.btcUSD:
			printTicker()
		case btcEUR = <-data.btcEUR:
			printTicker()
		case eurUSD = <-data.eurUSD:
			printTicker()
		case activeBTC = <-data.activeBTC:
			printTicker()
		case activeCUR = <-data.activeCUR:
			printTicker()
		case msg := <-quit:
			fmt.Printf("\n%v", msg)
			return
		}
	}

}
