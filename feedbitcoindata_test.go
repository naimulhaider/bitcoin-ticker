package main

import (
	"testing"
	"time"
)

func TestFeedBitcoinData(t *testing.T) {

	data := NewData()

	go FeedBitcoinData(data)

	timeout := time.After(time.Duration(IntervalConfig-1) * time.Second)

	cnt := 0

	for {
		select {
		case <-data.btcUSD:
			cnt++
		case <-data.btcEUR:
			cnt++
		case <-data.activeBTC:
			cnt++
		case <-timeout:
			t.FailNow()
		}
		if cnt == 3 {
			break
		}
	}
}
