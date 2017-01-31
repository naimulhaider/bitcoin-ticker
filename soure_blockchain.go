package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type BlockchainSource struct {
	USD struct {
		Last float64 `json:"last"`
	} `json:"USD"`
	EUR struct {
		Last float64 `json:"last"`
	} `json:"EUR"`
}

func (s *BlockchainSource) GetUSD() float64 {
	return s.USD.Last
}

func (s *BlockchainSource) GetEUR() float64 {
	return s.EUR.Last
}

func (s *BlockchainSource) Update() error {
	resp, err := http.Get("https://blockchain.info/ticker")
	if err != nil {
		return fmt.Errorf("Failed to fetch from blockchain: %v", err)
	}
	defer resp.Body.Close()

	src := BlockchainSource{}
	err = json.NewDecoder(resp.Body).Decode(&src)
	if err != nil {
		return fmt.Errorf("Unexpected Response: %v", err)
	}

	s.USD = src.USD
	s.EUR = src.EUR

	return nil
}
