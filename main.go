package main

func main() {

	data := NewData()
	quit := make(chan string, 1)

	// go DataSimulator(data, quit)

	go FeedBitcoinData(data, quit)
	go FeedCurrencyData(data, quit)

	Printer(data, quit)

}
