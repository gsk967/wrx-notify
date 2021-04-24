package utils

import (
	"log"
	"strconv"
)

type ExchangeResponse struct {
	At     int64 `json:"at"`
	Ticker struct {
		Buy  string `json:"buy"`
		Sell string `json:"sell"`
		Low  string `json:"low"`
		High string `json:"high"`
		Last string `json:"last"`
		Vol  string `json:"vol"`
	} `json:"ticker"`
}

func (tickerExchangeResponse ExchangeResponse) PrintLastValue(exchange string) {
	log.Printf("Exchange %s At %s , Last exchange value : %s", exchange,
		GetTimeStamp(tickerExchangeResponse.At), tickerExchangeResponse.Ticker.Last)
}

func (tickerExchangeResponse ExchangeResponse) ShowPriceNotification(exchange string, high float64, low float64) {
	currentPrice, _ := strconv.ParseFloat(tickerExchangeResponse.Ticker.Last, 10)
	if currentPrice <= low {
		ShowNotification(exchange, currentPrice, false)
	}
	if currentPrice >= high {
		ShowNotification(exchange, currentPrice, true)
	}
}

type ExchangeList struct {
	Exchanges []Exchange `json:"exchanges"`
}

type Exchange struct {
	Low      float64 `json:"low"`
	High     float64 `json:"high"`
	Coin     string  `json:"coin"`
	Exchange string  `json:"exchange"`
}
