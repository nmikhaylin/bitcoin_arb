package main

import (
  "fmt"
  "net/http"
  "log"
  "os"
  "io/ioutil"
  "encoding/json"
)

const kApiKey string = "20aeeaae2f8341d2842ef67af9ab44dd"
const kApiKeyInfoEndpoint string = "http://www.coinwarz.com/v1/api/apikeyinfo"

type ApiKeyInfo struct {
  success bool  `json:Success`
  message string `json:Message`

}
func GetApiKeyInfo() {
  req, err := http.NewRequest("GET", kApiKeyInfoEndpoint, nil)
  if err != nil {
      log.Print(err)
      os.Exit(1)
  }

  q := req.URL.Query()
  q.Add("apikey", kApiKey)
  req.URL.RawQuery = q.Encode()
  fmt.Println(req.URL.String())
  resp, err := http.Get(req.URL.String())
  if (err != nil) {
    fmt.Print(err)
  }
  fmt.Println(resp.Status)
  fmt.Println(resp.Body)
  info := ApiKeyInfo{}
  read_body, err := ioutil.ReadAll(resp.Body)
  json.Unmarshal(read_body, &info)
  fmt.Println(info)
}

func main() {
  //GetApiKeyInfo()
  response, _ := http.Get("http://www.coinwarz.com/v1/api/apikeyinfo?apikey=20aeeaae2f8341d2842ef67af9ab44dd")
  buf, _ := ioutil.ReadAll(response.Body)
  read_data := ApiKeyInfo{}
  fmt.Println(response)
  json.Unmarshal(buf, &read_data)
  fmt.Println(read_data)
}
