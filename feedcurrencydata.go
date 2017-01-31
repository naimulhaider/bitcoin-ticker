package main

import (
	"fmt"
	"log"
	"time"
)

func FeedCurrencyData(data *Data, quit chan string) {

	sources := GetCurrencySources()
	data.totalCUR = len(sources)

	fetchData := func(source CurrencyDataSource, pipe chan CurrencyDataSource) {
		err := source.Update()
		if err != nil {
			log.Println(fmt.Errorf("Failed to fetch data, err: %v", err))
			return
		}
		pipe <- source
	}

	updateData := func() {

		pipe := make(chan CurrencyDataSource, len(sources))

		foundOne := false

		for _, source := range sources {
			go fetchData(source, pipe)
		}

		responsiveSources := 0

		timeout := time.After(time.Duration(IntervalConfig-1) * time.Second)

		for {
			doneListening := false
			select {
			case src := <-pipe:
				if foundOne == false { // this is the quickest response
					data.eurUSD <- src.GetEURUSD()
					foundOne = true
				}
				responsiveSources++
				if responsiveSources == len(sources) {
					doneListening = true // all of the sources responded
					break
				}
			case <-timeout:
				// timed out
				doneListening = true
				break
			}

			if doneListening {
				close(pipe)
				break
			}
		}

		data.activeCUR <- responsiveSources
		return
	}

	for {
		go updateData()
		time.Sleep(time.Duration(IntervalConfig) * time.Second)
	}

}
