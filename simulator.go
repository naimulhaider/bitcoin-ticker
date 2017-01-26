package main

import (
	"math/rand"
	"time"
)

func DataSimulator(data *Data, quit chan int) {

	for i := 0; i < 100; i++ {
		rand.Seed(time.Now().UnixNano())
		marker := rand.Intn(5)
		rand.Seed(time.Now().UnixNano())

		switch marker {
		case 1:
			magicNum := rand.Float64()
			data.btcUSD <- magicNum
		case 2:
			magicNum := rand.Float64()
			data.btcEUR <- magicNum
		case 3:
			magicNum := rand.Float64()
			data.eurUSD <- magicNum
		case 4:
			magicNum := rand.Intn(5)
			data.activeBTC <- magicNum
		case 5:
			magicNum := rand.Intn(5)
			data.activeCUR <- magicNum
		}

		time.Sleep(1 * time.Second)
	}

	// quit <- 0
}
