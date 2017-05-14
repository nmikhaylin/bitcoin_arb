package main

import (
  "fmt"
  "api_calls"
)

func main() {
  nice_hash_market, err := api_calls.GetNiceHashMarket()
  if (err != nil) {
    panic(err)
  }
  fmt.Printf("%+v\n", nice_hash_market)
  profit, err := api_calls.GetProfitabilityInfo()
  if (err != nil) {
    panic(err)
  }
  fmt.Printf("%+v\n", profit)
  fmt.Printf("Current Bitcoin Price: %f\n", current_bitcoin_price)
}
