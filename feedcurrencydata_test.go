package main

import (
	"testing"
	"time"
)

func TestFeedCurrencyData(t *testing.T) {

	data := NewData()

	go FeedCurrencyData(data)

	timeout := time.After(time.Duration(IntervalConfig-1) * time.Second)

	cnt := 0

	for {
		select {
		case <-data.eurUSD:
			cnt++
		case <-data.activeCUR:
			cnt++
		case <-timeout:
			t.FailNow()
		}
		if cnt == 2 {
			break
		}
	}
}
