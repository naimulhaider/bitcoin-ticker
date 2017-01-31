package main

type Data struct {
	btcUSD    chan float64
	btcEUR    chan float64
	eurUSD    chan float64
	activeBTC chan int
	activeCUR chan int
	totalBTC  int
	totalCUR  int
}

func NewData() *Data {
	data := &Data{
		btcUSD:    make(chan float64, 100),
		btcEUR:    make(chan float64, 100),
		eurUSD:    make(chan float64, 100),
		activeBTC: make(chan int, 100),
		activeCUR: make(chan int, 100),
	}
	return data
}
