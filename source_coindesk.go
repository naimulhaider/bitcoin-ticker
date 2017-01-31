package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type CoindeskSource struct {
	Bpi struct {
		USD struct {
			RateFloat float64 `json:"rate_float"`
		} `json:"USD"`
		EUR struct {
			RateFloat float64 `json:"rate_float"`
		} `json:"EUR"`
	} `json:"bpi"`
}

func (s *CoindeskSource) GetUSD() float64 {
	return s.Bpi.USD.RateFloat
}

func (s *CoindeskSource) GetEUR() float64 {
	return s.Bpi.EUR.RateFloat
}

func (s *CoindeskSource) Update() error {

	timeout := time.Duration(IntervalConfig-1) * time.Second
	client := http.Client{
		Timeout: timeout,
	}

	resp, err := client.Get("https://api.coindesk.com/v1/bpi/currentprice.json")
	if err != nil {
		return fmt.Errorf("Failed to fetch from coindesk: %v", err)
	}
	defer resp.Body.Close()

	src := CoindeskSource{}
	err = json.NewDecoder(resp.Body).Decode(&src)
	if err != nil {
		return fmt.Errorf("Unexpected Response: %v", err)
	}

	s.Bpi = src.Bpi

	return nil
}
