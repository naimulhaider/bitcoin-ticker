package main

func main() {

	data := NewData()
	quit := make(chan int)

	go DataSimulator(data, quit)

	Printer(data, quit)

}
