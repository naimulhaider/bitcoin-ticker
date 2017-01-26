package main

import "time"

func FeedBitcoinData(data *Data, quit chan string) {
	sources := GetBitcoinSources()
	data.totalBTC = len(sources)

	for {
		var eur, usd float64 = 0, 0
		active := 0

		for _, source := range sources {
			err := source.Update()
			if err != nil {
				continue
			}
			active++
			eur += source.GetEUR()
			usd += source.GetUSD()
		}

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
