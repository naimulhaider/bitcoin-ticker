package main

import "fmt"

func Printer(data *Data, quit chan int) {
	var btcUSD, btcEUR, eurUSD float64 = 0, 0, 0
	var activeBTC, activeCUR int = 0, 0

	printTicker := func() {
		fmt.Printf("\rBTC/USD: %f   EUR/USD: %f   BTC/EUR: %f Active sources: BTC/USD (%d of %d)  EUR/USD (%d of %d)", btcUSD, eurUSD, btcEUR, activeBTC, len(BitcoinSources), activeCUR, len(CurrencySources))
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
		case <-quit:
			return
		}
	}

}
