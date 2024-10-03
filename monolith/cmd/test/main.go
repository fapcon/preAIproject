//go:build !test

package main

import (
	"context"
	"fmt"
	"github.com/shopspring/decimal"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/modules/exchange_client/service"
	"time"

	"github.com/adshao/go-binance/v2"
)

var (
	apiKey    = "asdf"
	secretKey = "fasdfasdf"
)

var (
	apiKeyJulustan    = "asdfsadf"
	secretKeyJulustan = "sadfasdf"
)

func main() {
	client := binance.NewClient(apiKeyJulustan, secretKeyJulustan)
	symbol := "LTCBUSD"
	interval := "15m"
	limit := 100
	startTime := int64(0)
	endTime := time.Now().Add(time.Duration(-3) * time.Hour).Unix()
	klines, err := client.NewKlinesService().Symbol(symbol).
		Interval(interval).Limit(limit).StartTime(startTime).
		EndTime(endTime).Do(context.Background())
	fmt.Println(klines, err)
}

func addBNB() {
	exchange := service.NewBinanceWithKey(1, nil, "apiKeyJulustan", "secretKeyJulustan")
	exchange.BuyMarket(service.MarketIn{
		Pair:     "BNBBUSD",
		Quantity: decimal.NewFromInt(1),
	})
	acc := exchange.GetAccount(context.Background())
	fmt.Println(acc)
}

func test() {
	exchange := service.NewBinanceWithKey(2, nil, "apiKey", "secretKey")
	orderOut := exchange.GetOrder(context.Background(), service.GetOrderIn{
		Pair:    "LTCBUSD",
		OrderID: "563860452",
	})
	fmt.Println(orderOut)
}

func getAccount() {
	exchange := service.NewBinanceWithKey(3, nil, apiKeyJulustan, secretKeyJulustan)
	account := exchange.GetAccount(context.Background())
	fmt.Println(account)
}

func getAccountRaw(client *binance.Client) {
	res, err := client.NewGetAccountService().Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res)
}

func Buy(client *binance.Client, amount string) {
	order, err := client.NewCreateOrderService().Symbol("BTCBUSD").
		Side(binance.SideTypeSell).Type(binance.OrderTypeLimit).
		TimeInForce(binance.TimeInForceTypeGTC).Quantity("0.002").
		Price("24600.91").Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(order)
}

func trade(client *binance.Client) {
	var sideType binance.SideType
	var (
		buy, sell int
	)
	sell = 1
	for i := 0; i < 45; i++ {
		if buy > sell {
			sideType = binance.SideTypeSell
			sell++
		} else {
			sideType = binance.SideTypeBuy
			buy++
		}
		n := i
		go func() {
			order, err := client.NewCreateOrderService().Symbol("ETHBTC").
				Side(sideType).Type(binance.OrderTypeMarket).Quantity("0.002").Do(context.Background())
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(n, order)
		}()
	}
}
