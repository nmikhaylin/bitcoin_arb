package main

import (
  "fmt"
  "api_calls"
  "errors"
)

const kTurnoverMin float32 = .5

func FindMarketRates(market *api_calls.NiceHashMarket) (map[string]api_calls.MarketRate, error){
  ret := map[string]api_calls.MarketRate{}
  for _, v := range market.AlgoMarkets {
    var speed_sum float32 = 0.0
    var i int
    for i = len(v.Orders) - 1; i >= 0 ; i-- {
      speed_sum += (v.Orders[i].AcceptedSpeed * v.Orders[i].Price)
      if speed_sum > kTurnoverMin {
        ret[v.Algo] = api_calls.MarketRate{v.Algo, v.Orders[i].Price}
        break
      }
    }
    if (i < 0) {
      return nil, errors.New(fmt.Sprintf(
          "Not enough Volumn for Algo: %s", v.Algo))
    }
  }
  return ret, nil
}

func main() {
  nice_hash_market, err := api_calls.GetNiceHashMarket()
  if (err != nil) {
    panic(err)
  }
  rates, err := FindMarketRates(nice_hash_market)
  if (err != nil) {
    panic(err)
  }
  fmt.Printf("%+v", rates)
  profit, err := api_calls.GetProfitabilityInfo(rates)
  if (err != nil) {
    panic(err)
  }
  for _, v := range profit {
    fmt.Printf("%s: %f\n", v.CoinName, v.ProfitRatio)

  }
}
