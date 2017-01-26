package main

type BitcoinDataSource interface {
	GetUSD() float64
	GetEUR() float64
	Update() error
}

type CurrencyDataSource interface {
	GetEURUSD() float64
	Update() error
}

func GetBitcoinSources() []BitcoinDataSource {
	ret := []BitcoinDataSource{}
	ret = append(ret, &CoindeskSource{})
	ret = append(ret, &BlockchainSource{})
	return ret
}

func GetCurrencySources() []CurrencyDataSource {
	ret := []CurrencyDataSource{}
	ret = append(ret, &FixerSource{})
	return ret
}
