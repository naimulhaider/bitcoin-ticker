package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type FixerSource struct {
	Rates struct {
		USD float64 `json:"USD"`
	} `json:"rates"`
}

func (s *FixerSource) GetEURUSD() float64 {
	return s.Rates.USD
}

func (s *FixerSource) Update() error {

	timeout := time.Duration(IntervalConfig-1) * time.Second
	client := http.Client{
		Timeout: timeout,
	}

	resp, err := client.Get("http://api.fixer.io/latest")
	if err != nil {
		return fmt.Errorf("Failed to fetch from coindesk: %v", err)
	}
	defer resp.Body.Close()

	src := FixerSource{}
	err = json.NewDecoder(resp.Body).Decode(&src)
	if err != nil {
		return fmt.Errorf("Unexpected Response: %v", err)
	}

	s.Rates = src.Rates

	return nil
}
