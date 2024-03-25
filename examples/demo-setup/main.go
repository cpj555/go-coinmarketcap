package main

import (
	"context"
	"fmt"
	go_coinmarketcap "github.com/cpj555/go-coinmarketcap"
	"github.com/cpj555/go-coinmarketcap/client"
	"github.com/cpj555/go-coinmarketcap/types"
	"time"
)

func main() {

	c, err := client.New(client.WithSandBox(true), client.WithTimeout(10*time.Second), client.WithProxyUrl("http://localhost:1087"))
	if err != nil {
		fmt.Printf("创建实例失败：%v\n", err)
		return
	}

	//getMap, err := c.GetMap(context.Background(), &types.GetMapReq{Start: 1, Limit: 1})
	//if err != nil {
	//	fmt.Printf("getMap err：%v\n", err)
	//	return
	//}
	//fmt.Printf("%+v\n", getMap)

	//m, err := c.ExchangeV1.GetMap(context.Background(), &types.GetExchangeMapReq{Start: 1, Limit: 10})
	//fmt.Println(err)
	//fmt.Println(m)

	getExchangeInfo(c)
	//getInfo, err := c.CryptocurrencyV1.GetMap(context.Background(), &types.GetCryptocurrencyMapReq{Start: 1, Limit: 10})
	//if err != nil {
	//	fmt.Printf("getMap err：%v\n", err)
	//	return
	//}
	//fmt.Printf("%+v\n", getInfo)

}

func getExchangeInfo(c *client.Client) {
	m, err := c.ExchangeV1.GetInfo(context.Background(), &types.GetExchangeInfoReq{Slug: "thruster-v2-0-3-fee"})
	fmt.Println(err)
	fmt.Println(m)

}

func getQuotes(c *client.Client) {
	m, err := c.ExchangeV1.GetQuotesLatest(context.Background(), &types.GetExchangeQuotesReq{Slug: "poloniex", Convert: "USD,CNY"})
	fmt.Println(err)
	fmt.Println(m)
}

func getCryptMap(c *client.Client) {

	m, err := c.CryptocurrencyV1.GetMap(context.Background(), &types.GetCryptocurrencyMapReq{Start: 1, Limit: 1})
	fmt.Println(err)
	fmt.Println(m)
}
