package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	ch := make(chan float64, 2)

	go func() {
		resp, err := http.Get("https://api.coinpaprika.com/v1/tickers/btc-bitcoin")
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		b, _ := ioutil.ReadAll(resp.Body)
		result := map[string]interface{}{}
		json.Unmarshal(b, &result)
		elem := result["quotes"].(map[string]interface{})
		price := elem["USD"].(map[string]interface{})["price"]
		ch <- price.(float64)
	}()

	go func() {
		resp, err := http.Get("https://api.coingecko.com/api/v3/simple/price?ids=bitcoin&vs_currencies=usd")
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		b, _ := ioutil.ReadAll(resp.Body)
		result := map[string]interface{}{}
		json.Unmarshal(b, &result)
		price := result["bitcoin"].(map[string]interface{})["usd"]
		ch <- price.(float64)
	}()

	fmt.Println(<-ch)
}
