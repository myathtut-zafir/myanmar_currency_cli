package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type ExchangeRate struct {
	Data      []CurrencyData `json:"data"`
	Epoch     float64        `json:"epoch"`
	Timestamp string         `json:"timestamp"`
}

type CurrencyData struct {
	Currency string `json:"currency"`
	Buy      string `json:"buy"`
	Sell     string `json:"sell"`
}

func main() {
	res, err := http.Get("https://myanmar-currency-api.github.io/api/latest.json")
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	//fmt.Println(res)
	if res.StatusCode != 200 {
		panic("API not good!")
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	//fmt.Println(string(body))
	var exchange ExchangeRate
	err = json.Unmarshal(body, &exchange)
	if err != nil {
		panic(err)
	}
	//currencyType,sell,buy:=.
	//fmt.Println(exchange.Data)
	people := []CurrencyData{}
	//currencyMap := make(map[string]string)
	for _, currencyData := range exchange.Data {
		if currencyData.Currency == "USD" || currencyData.Currency == "THB" || currencyData.Currency == "SGD" {

			newData := CurrencyData{
				Currency: "Currency : " + currencyData.Currency + ",",
				Buy:      "Sell Price  : " + currencyData.Sell + ",",
				Sell:     "Buy Price  : " + currencyData.Buy,
			}
			people = append(people, newData)
		}

	}
	fmt.Println("Date : ", exchange.Timestamp)
	fmt.Println("information : ", people)

}
