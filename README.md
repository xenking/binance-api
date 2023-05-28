# Golang Binance API

[![PkgGoDev](https://pkg.go.dev/badge/github.com/xenking/binance-api)](https://pkg.go.dev/github.com/xenking/binance-api)
[![Go Report Card](https://goreportcard.com/badge/github.com/xenking/binance-api)](https://goreportcard.com/report/github.com/xenking/binance-api)
![Build Status](https://github.com/xenking/binance-api/actions/workflows/build.yml/badge.svg)
[![codecov](https://codecov.io/gh/xenking/binance-api/branch/master/graph/badge.svg)](https://codecov.io/gh/xenking/binance-api)

binance-api is a fast and lightweight Golang implementation for [Binance API](https://github.com/binance/binance-spot-api-docs), providing complete API coverage, and supports both REST API and websockets API

This library created to help you interact with the Binance API, streaming candlestick charts data, market depth, or use other advanced features binance exposes via API. 

## Quickstart
```golang
package main

import (
    "log"

    "github.com/xenking/binance-api"
)

func main() {
    client := binance.NewClient("API-KEY", "SECRET")

    err := client.Ping()
    if err != nil {
        panic(err)
    }

    prices, err := client.Prices()
    if err != nil {
        panic(err)
    }

    for _, p := range prices {
        log.Printf("symbol: %s, price: %s", p.Symbol, p.Price)
    }
}
```


## Installation
```
go get -u github.com/xenking/binance-api
```

## Getting started
```golang
// Create default client
client := binance.NewClient("API-KEY", "SECRET")

// Send ping request
err := client.Ping()

// Create client with custom request window size
client := binance.NewClient("API-KEY", "SECRET").ReqWindow(5000)

// Create websocket client
wsClient := ws.NewClient()

// Connect to Klines websocket
ws, err := wsClient.Klines("ETHBTC", binance.KlineInterval1m)

// Read ws
msg, err := ws.Read()
```

Full documentation on [pkg.go.dev](https://pkg.go.dev/github.com/xenking/binance-api)

## License
This library is under the [MIT License](https://opensource.org/licenses/MIT). See the [LICENSE](LICENSE.md) file for more info.

