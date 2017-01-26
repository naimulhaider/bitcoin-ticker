
BITCOIN TICKER

There are two types of sources:

type BitcoinDataSource interface {
	GetUSD() float64 // BTC to USD
	GetEUR() float64 // BTC to EUR
	Update() error // update values
}

type CurrencyDataSource interface {
	GetEURUSD() float64 // EUR to USD
	Update() error // update values
}

To add a source, simply implement one of the interface


Currently these sources are added:

Bitcoin:
https://api.coindesk.com/v1/bpi/currentprice.json

Currency:
http://api.fixer.io/latest


The ticker updates data from all its sources every 10 seconds, and calculates the average from all the available sources
