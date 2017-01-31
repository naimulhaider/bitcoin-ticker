package main

func main() {

	data := NewData()

	// go DataSimulator(data, quit)

	go FeedBitcoinData(data)
	go FeedCurrencyData(data)

	Printer(data)

}
