package types

import "time"

// [https://coinmarketcap.com/api/documentation/v1/#operation/getV3CryptocurrencyQuotesHistorical]
type GetCryptocurrencyQuotesHistoricalReq struct {
	GetCryptocurrencyQuotesReq
	TimeStart string `schema:"time_start,omitempty"`
	TimeEnd   string `schema:"time_end,omitempty"`
	Count     int    `schema:"count,omitempty"`
	Interval  string `schema:"interval,omitempty"`
}

type GetCryptocurrencyQuoteHistoricalResp map[string]GetCryptocurrencyQuoteHistoricalItem

type GetCryptocurrencyQuoteHistoricalItem struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Symbol   string `json:"symbol"`
	IsActive int    `json:"is_active"`
	IsFiat   int    `json:"is_fiat"`
	Quotes   []struct {
		TimeStamp time.Time `json:"timestamp"`
		Quote     map[string]struct {
			PercentChange1H   float64   `json:"percent_change_1h"`
			PercentChange24H  float64   `json:"percent_change_24h"`
			PercentChange7D   float64   `json:"percent_change_7d"`
			PercentChange30D  float64   `json:"percent_change_30d"`
			Price             float64   `json:"price"`
			Volume24H         float64   `json:"volume_24h"`
			MarketCap         float64   `json:"market_cap"`
			CirculatingSupply float64   `json:"circulating_supply"`
			TotalSupply       float64   `json:"total_supply"`
			TimeStamp         time.Time `json:"timestamp"`
		} `json:"quote"`
	} `json:"quotes"`
}
