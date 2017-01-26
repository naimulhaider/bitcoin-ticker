package main

import "time"

func FeedCurrencyData(data *Data, quit chan string) {
	sources := GetCurrencySources()
	data.totalCUR = len(sources)

	for {
		var eurusd float64
		active := 0

		for _, source := range sources {
			err := source.Update()
			if err != nil {
				continue
			}
			active++
			eurusd += source.GetEURUSD()
		}

		if active == 0 {
			quit <- "No currency feeds are active!"
			return
		}

		data.eurUSD <- eurusd / float64(active)
		data.activeCUR <- active

		time.Sleep(time.Duration(IntervalConfig) * time.Second)
	}
}
