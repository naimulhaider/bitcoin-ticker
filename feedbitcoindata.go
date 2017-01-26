package main

import (
	"sync"
	"time"
)

func FeedBitcoinData(data *Data, quit chan string) {
	sources := GetBitcoinSources()
	data.totalBTC = len(sources)

	for {
		var eur, usd float64 = 0, 0
		active := 0

		var wg sync.WaitGroup

		for _, source := range sources {
			wg.Add(1)
			go func(src BitcoinDataSource) {
				err := src.Update()
				if err != nil {
					return
				}
				active++
				eur += src.GetEUR()
				usd += src.GetUSD()
				wg.Done()
			}(source)
		}

		wg.Wait()

		if active == 0 {
			quit <- "No bitcoin feeds are active!"
			return
		}

		data.btcEUR <- eur / float64(active)
		data.btcUSD <- usd / float64(active)
		data.activeBTC <- active

		time.Sleep(time.Duration(IntervalConfig) * time.Second)
	}

}
