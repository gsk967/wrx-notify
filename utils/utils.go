package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"time"
)

// StartWatchingExchanges it will watch the exchanges
func StartWatchingExchanges(exchangeFilePath string, interval time.Duration, tgAlert *TelegramAlert) {
	exchanges, err := ReadExchangesListJson(exchangeFilePath)
	if err != nil {
		log.Fatalf("Error while reading the json file %s : %v\n", exchangeFilePath, err)
	}

	ticker := time.NewTicker(interval * time.Second)
	done := make(chan bool)

	for {
		select {
		case <-done:
			log.Println("Stop the ticker....")
			return
		case <-ticker.C:
			for _, exchange := range exchanges.Exchanges {
				WatchingExchange(exchange, done, tgAlert)
			}
		}
	}
}

// WatchingExchange
func WatchingExchange(exchange Exchange, done chan<- bool, tgAlert *TelegramAlert) {
	log.Println("Watching exchange ", exchange.Exchange, " At ", GetTimeStamp(time.Now().Unix()))

	tickerUrl := TICKER_URL + exchange.Exchange
	resp, err := http.Get(tickerUrl)
	if err != nil {
		log.Printf("Error while getting the infos %v\n", err)
	}
	if err != nil {
		log.Fatalf("Error while getting the infos %v\n", err)
	}
	defer resp.Body.Close()

	var tickerResponse ExchangeResponse

	if err := json.NewDecoder(resp.Body).Decode(&tickerResponse); err != nil {
		done <- true
		log.Fatalf("Error while json decoding the resp : %v\n", err)
	}
	tickerResponse.PrintLastValue(exchange.Exchange)
	if runtime.GOOS == "linux" {
		tickerResponse.ShowPriceNotification(exchange.Exchange, exchange.High, exchange.Low)
	}

	if tgAlert != nil {
		currentPrice, _ := strconv.ParseFloat(tickerResponse.Ticker.Last, 10)
		if currentPrice < exchange.Low {
			message := fmt.Sprintf("Exchange %s price is Below this price : %f\nPrice is %f", exchange.Exchange, exchange.Low, currentPrice)
			err := tgAlert.SendTelegramMessage(message)
			if err != nil {
				log.Fatalln("Error while sending the notification to telegram bot", "Error", err)
			}
		}

		if currentPrice > exchange.High {
			message := fmt.Sprintf("Exchange %s price is Higher than this price : %f\nPrice is %f %s", exchange.Exchange, exchange.High, currentPrice, strings.ToUpper(exchange.Exchange[len(exchange.Exchange)-3:]))
			err := tgAlert.SendTelegramMessage(message)
			if err != nil {
				log.Fatalln("Error while sending the notification to telegram bot", "Error", err)
			}
		}
	}
}

func GetTimeStamp(at int64) string {
	loc, _ := time.LoadLocation("Asia/Kolkata")
	t := time.Unix(at, 0).In(loc)
	return t.Local().String()
}

// ReadExchangesListJson will read json format of exchanges
func ReadExchangesListJson(filePath string) (ExchangeList, error) {
	var exchangesList ExchangeList
	bodyBytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return ExchangeList{}, err
	}

	if err := json.Unmarshal(bodyBytes, &exchangesList); err != nil {
		return ExchangeList{}, err
	}

	return exchangesList, nil
}

// ShowNotification will show the notification about exchange price
func ShowNotification(exchange string, price float64, high bool) {
	var commandArgs []string
	if high {
		commandArgs = []string{"-u", "normal", "-a", "WRX-NOTIFY",
			fmt.Sprintf("Exchange %s Price is HIGH", exchange),
			fmt.Sprintf("Price is %f", price)}
	} else {
		commandArgs = []string{"-u", "normal", "-a", "WRX-NOTIFY",
			fmt.Sprintf("Exchange %s Price is Low", exchange),
			fmt.Sprintf("Price is %d", price)}
	}
	_, _ = exec.Command("notify-send", commandArgs...).Output()
}
