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
		btcUSD:    make(chan float64),
		btcEUR:    make(chan float64),
		eurUSD:    make(chan float64),
		activeBTC: make(chan int),
		activeCUR: make(chan int),
	}
	return data
}
