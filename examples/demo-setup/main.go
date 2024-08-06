package main

import (
	"context"
	"encoding/json"
	"fmt"
	go_coinmarketcap "github.com/cpj555/go-coinmarketcap"
	"github.com/cpj555/go-coinmarketcap/client"
	"github.com/cpj555/go-coinmarketcap/tools"
	"github.com/cpj555/go-coinmarketcap/types"
	"time"
)

func main() {

	opts := []client.OptionHandler{
		client.WithProxyUrl("http://localhost:1087"),
		client.WithApiKey("b354b6f0-53a6-4fb4-8496-76657015d736"),
		client.WithTimeout(30 * time.Second),
		client.WithRequestPerSecond(1),
	}

	c, err := go_coinmarketcap.NewClient(opts...)
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

	//getExchangeInfo(c)
	//getQuotes(c)
	//getCryptMap(c)
	//getCryptInfo(c)
	//getCryptQuotesHistorical(c)
	getExchangeQuotesHistorical(c)
	//getCryptQuotes(c)
	//getExchangeMarketPair(c)
	//getCryptocurrencyMarketPair(c)
	//getGlobalMetricsQuotes(c)
	//getPricePerformance(c)
	//go getFiatMap(c)
	//go getFiatMap(c)
	//go getFiatMap(c)

	//time.Sleep(time.Second * 5)
	//getInfo, err := c.CryptocurrencyV1.GetMap(context.Background(), &types.GetCryptocurrencyMapReq{Start: 1, Limit: 10})
	//if err != nil {
	//	fmt.Printf("getMap err：%v\n", err)
	//	return
	//}
	//fmt.Printf("%+v\n", getInfo)

}

func getExchangeInfo(c *client.Client) {
	m, err := c.ExchangeV1.GetInfo(context.Background(), &types.GetExchangeInfoReq{Slug: "exmarkets"})
	fmt.Println(err)
	fmt.Println(m)

}

func getQuotes(c *client.Client) {
	m, err := c.ExchangeV1.GetQuotesLatest(context.Background(), &types.GetExchangeQuotesReq{Slug: "binance", Convert: "CNY"})
	fmt.Println(err)
	fmt.Println(m)
}

func getCryptMap(c *client.Client) {

	m, err := c.CryptocurrencyV1.GetMap(context.Background(), &types.GetCryptocurrencyMapReq{Start: 1, Limit: 1, Aux: "first_historical_data,last_historical_data,is_active,status"})
	fmt.Println(err)
	fmt.Println(m)
}

func getCryptInfo(c *client.Client) {
	m, err := c.CryptocurrencyV2.GetInfo(context.Background(), &types.GetCryptocurrencyInfoReq{Id: "9337"})
	fmt.Println(err)
	fmt.Println(m)
}

func getCryptQuotes(c *client.Client) {
	m, err := c.CryptocurrencyV2.GetQuotesLatest(context.Background(), &types.GetCryptocurrencyQuotesReq{Id: "1,2,3,5,6,8,10,13,16,18,25,35,41,42,43,45,52,53,56,61,66,67,69,72,74,77,78,83,87,90,93,99,109,118,122,128,131,132,145,148,154,162,168,170,184,213,215,217,234,258,260,263,268,276,278,290,291,293,298,316,328,333,360,362,367,372,377,389,405,416,448,460,463,470,501,502,506,512,541,551,558,572,576,584,597,601,638,644,656,659,693,702,703,707,720,733,760,764,788,815", Convert: "USD,CNY,BTC"})
	fmt.Println(err)
	fmt.Println(m)
}

func getCryptQuotesHistorical(c *client.Client) {
	m, err := c.CryptocurrencyV3.GetQuotesHistorical(context.Background(), &types.GetCryptocurrencyQuotesHistoricalReq{
		GetCryptocurrencyQuotesReq: types.GetCryptocurrencyQuotesReq{Id: "1", Convert: "USD,BTC,CNY"}, Interval: "daily"})
	fmt.Println(err)
	fmt.Println(m)
}

func getExchangeQuotesHistorical(c *client.Client) {
	mm, err := c.ExchangeV1.GetQuotesHistorical(context.Background(), &types.GetExchangeQuotesHistoricalReq{
		GetExchangeQuotesReq: types.GetExchangeQuotesReq{Id: "16,21", Convert: "CNY"}, TimeStart: "2019-08-06T07:35:06.849Z", Count: 1, Interval: "daily"})
	fmt.Println(err)
	fmt.Printf("%+v\n", (*mm)["16"].Quotes[0].Quote["CNY"].Volume24H)

	test := struct {
		X float64
	}{
		X: (*mm)["16"].Quotes[0].Quote["CNY"].Volume24H,
	}
	ss, _ := tools.JSON.Marshal(test)
	fmt.Println(string(ss))

}

func getPricePerformance(c *client.Client) {
	m, err := c.CryptocurrencyV2.GetPricePerformanceStatsLatest(context.Background(), &types.GetCryptocurrencyPricePerformanceStatsReq{TimePeriod: "all_time,24h,7d", Id: "1,2,3,5,6,8,10,13,16,18,25,35,41,42,43,45,52,53,56,61,66,67,69,72,74,77,78,83,87,90,93,99,109,118,122,128,131,132,145,148,154,162,168,170,184,213,215,217,234,258,260,263,268,276,278,290,291,293,298,316,328,333,360,362,367,372,377,389,405,416,448,460,463,470,501,502,506,512,541,551,558,572,576,584,597,601,638,644,656,659,693,702,703,707,720,733,760,764,788,815,819,825,831,837,853,857,859,873,894,895,898,916,918,921,934,938,945,948,951,978,986,993,1004,1019,1020,1027,1032,1033,1035,1037,1038,1042,1044,1052,1053,1066,1070,1082,1085,1090,1104,1106,1107,1120,1135,1136,1146,1154,1155,1156,1159,1164,1165,1168,1169,1175,1185,1191,1194,1200,1209,1210,1212,1214,1216,1218,1223,1229,1230,1241,1244,1247,1248,1250,1252,1254,1257,1259,1273,1274,1279,1281,1282,1285,1291,1297,1298,1299,1306,1312,1320,1321,1343,1351,1353,1368,1376,1382,1389,1392,1395,1396,1414,1437,1439,1447,1455,1466,1468,1473,1474,1492,1495,1500,1503,1505,1511,1514,1515,1518,1521,1522,1528,1546,1552,1556,1558,1562,1567,1578,1582,1586,1588,1596,1609,1612,1619,1623,1624,1629,1630,1632,1636,1637,1638,1651,1654,1657,1658,1659,1660,1669,1674,1678,1680,1681,1684,1693,1694,1697,1698,1700,1703,1706,1708,1710,1712,1720,1721", Convert: "USD,CNY"})
	fmt.Println(err)
	fmt.Println(m)
}

func getExchangeMarketPair(c *client.Client) {
	m, err := c.ExchangeV1.GetMarketPairLatest(context.Background(), &types.GetExchangeMarketPairReq{Slug: "binance", Convert: "USD", Start: 1, Limit: 10})
	fmt.Println(err)
	fmt.Println(m)
	fmt.Println(len(m.MarketPairs))
}

type FloatTest struct {
	Price json.Number `json:"price"`
}

func getCryptocurrencyMarketPair(c *client.Client) {
	//tools.JSON.NewDecoder(bytes.NewReader(s))

	m, err := c.CryptocurrencyV2.GetMarketPairLatest(context.Background(), &types.GetCryptocurrencyMarketPairReq{
		Slug: "bitcoin", Start: 1,
		//Aux: "num_market_pairs,category,fee_type,market_url,currency_name,currency_slug,price_quote,notice,cmc_rank,effective_liquidity,market_score,market_reputation",
		Limit: 20, Sort: "cmc_rank_advanced",
	})
	fmt.Println(err)
	fmt.Println(m)
	fmt.Println(len(m.MarketPairs))
}

func getFiatMap(c *client.Client) {
	c.FiatV1.GetMap(context.Background(), &types.GetFiatMapReq{Limit: 1})
	//fmt.Println(err)
	//fmt.Println(m)
}

func getGlobalMetricsQuotes(c *client.Client) {
	m, err := c.GlobalMetricsV1.GetQuotesLatest(context.Background(), &types.GetGlobalMetricsQuotesReq{})
	fmt.Println(err)
	fmt.Printf("%+v\n", m)
	fmt.Println(m.DefiVolume24H)
}
