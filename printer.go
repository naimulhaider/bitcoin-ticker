package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"time"
)

func ClearTerminal() {
	if runtime.GOOS == "linux" {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else if runtime.GOOS == "windows" {
		cmd := exec.Command("cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func Printer(data *Data) {
	var btcUSD, btcEUR, eurUSD float64 = 0, 0, 0
	var activeBTC, activeCUR int = 0, 0

	printTicker := func() {
		ClearTerminal()
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
		}
	}

}
