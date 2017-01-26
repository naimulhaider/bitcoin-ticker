package main

import (
	"sync"
	"time"
)

func FeedCurrencyData(data *Data, quit chan string) {
	sources := GetCurrencySources()
	data.totalCUR = len(sources)

	for {
		var eurusd float64
		active := 0

		var wg sync.WaitGroup

		for _, source := range sources {
			wg.Add(1)
			go func(src CurrencyDataSource) {
				err := src.Update()
				if err != nil {
					return
				}
				active++
				eurusd += src.GetEURUSD()
				wg.Done()
			}(source)
		}

		wg.Wait()

		if active == 0 {
			quit <- "No currency feeds are active!"
			return
		}

		data.eurUSD <- eurusd / float64(active)
		data.activeCUR <- active

		time.Sleep(time.Duration(IntervalConfig) * time.Second)
	}
}
