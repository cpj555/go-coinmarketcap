package types

import (
	"time"
)

// [https://coinmarketcap.com/api/documentation/v1/#operation/getV2CryptocurrencyInfo]
type GetCryptocurrencyInfoReq struct {
	Id          string `schema:"id,omitempty"`
	Slug        string `schema:"slug,omitempty"`
	Symbol      string `schema:"symbol,omitempty"`
	Address     string `schema:"address,omitempty"`
	SkipInvalid bool   `schema:"skip_invalid,omitempty"`
	Aux         string `schema:"aux,omitempty"`
}

// GetInfoResp CoinMarketCap ID Map
type GetCryptocurrencyInfoResp map[string]struct {
	Urls struct {
		Website      []string `json:"website"`
		TechnicalDoc []string `json:"technical_doc"`
		Twitter      []string `json:"twitter"`
		Reddit       []string `json:"reddit"`
		MessageBoard []string `json:"message_board"`
		Announcement []string `json:"announcement"`
		Chat         []string `json:"chat"`
		Explorer     []string `json:"explorer"`
		SourceCode   []string `json:"source_code"`
	} `json:"urls"`
	Logo                          string    `json:"logo"`
	Id                            int       `json:"id"`
	Name                          string    `json:"name"`
	Symbol                        string    `json:"symbol"`
	Slug                          string    `json:"slug"`
	Description                   string    `json:"description"`
	Notice                        *string   `json:"notice"`
	DateAdded                     time.Time `json:"date_added"`
	DateLaunched                  time.Time `json:"date_launched"`
	Tags                          []string  `json:"tags"`
	Platform                      *Platform `json:"platform"`
	Category                      string    `json:"category"`
	SelfReportedCirculatingSupply *int      `json:"self_reported_circulating_supply"`
	SelfReportedMarketCap         *int      `json:"self_reported_market_cap"`
	SelfReportedTags              *[]string `json:"self_reported_tags"`
	InfiniteSupply                bool      `json:"infinite_supply"`
}

// Returns the latest market quote
// [https://coinmarketcap.com/api/documentation/v1/#operation/getV2CryptocurrencyQuotesLatest]
type GetCryptocurrencyQuotesReq struct {
	Id          string `schema:"id,omitempty"`
	Slug        string `schema:"slug,omitempty"`
	Symbol      string `schema:"symbol,omitempty"`
	Convert     string `schema:"convert,omitempty"`
	ConvertId   string `schema:"convert_id,omitempty"`
	Aux         string `schema:"aux,omitempty"`
	SkipInvalid bool   `schema:"skip_invalid,omitempty,default:true"`
}

type GetCryptocurrencyQuoteResp map[string]struct {
	Id                            int       `json:"id"`
	Name                          string    `json:"name"`
	Symbol                        string    `json:"symbol"`
	Slug                          string    `json:"slug"`
	IsActive                      int       `json:"is_active"`
	IsFiat                        int       `json:"is_fiat"`
	CirculatingSupply             int       `json:"circulating_supply"`
	TotalSupply                   int       `json:"total_supply"`
	MaxSupply                     int       `json:"max_supply"`
	DateAdded                     time.Time `json:"date_added"`
	NumMarketPairs                int       `json:"num_market_pairs"`
	CmcRank                       int       `json:"cmc_rank"`
	LastUpdated                   time.Time `json:"last_updated"`
	Tags                          []string  `json:"tags"`
	Platform                      *Platform `json:"platform"`
	SelfReportedCirculatingSupply *int      `json:"self_reported_circulating_supply"`
	SelfReportedMarketCap         *int      `json:"self_reported_market_cap"`
	Quote                         map[string]struct {
		Price                 float64   `json:"price"`
		Volume24H             float64   `json:"volume_24h"`
		VolumeChange24H       float64   `json:"volume_change_24h"`
		PercentChange1H       float64   `json:"percent_change_1h"`
		PercentChange24H      float64   `json:"percent_change_24h"`
		PercentChange7D       float64   `json:"percent_change_7d"`
		PercentChange30D      float64   `json:"percent_change_30d"`
		MarketCap             float64   `json:"market_cap"`
		MarketCapDominance    int       `json:"market_cap_dominance"`
		FullyDilutedMarketCap float64   `json:"fully_diluted_market_cap"`
		LastUpdated           time.Time `json:"last_updated"`
	} `json:"quote"`
}

// Returns price performance statistics
// [https://coinmarketcap.com/api/documentation/v1/#operation/getV2CryptocurrencyOhlcvLatest]
type GetCryptocurrencyPricePerformanceStatsReq struct {
	Id          string `schema:"id,omitempty"`
	Slug        string `schema:"slug,omitempty"`
	Symbol      string `schema:"symbol,omitempty"`
	TimePeriod  string `schema:"time_period,omitempty"`
	Convert     string `schema:"convert,omitempty"`
	ConvertId   string `schema:"convert_id,omitempty"`
	SkipInvalid bool   `schema:"skip_invalid,omitempty,default:true"`
}

type GetCryptocurrencyPricePerformanceStatsResp struct {
	Id          int       `json:"id"`
	Name        string    `json:"name"`
	Symbol      string    `json:"symbol"`
	Slug        string    `json:"slug"`
	LastUpdated time.Time `json:"last_updated"`
	Periods     map[string]struct {
		OpenTimestamp  time.Time `json:"open_timestamp"`
		HighTimestamp  time.Time `json:"high_timestamp"`
		LowTimestamp   time.Time `json:"low_timestamp"`
		CloseTimestamp time.Time `json:"close_timestamp"`
		Quote          map[string]struct {
			Open           float64   `json:"open"`
			OpenTimestamp  time.Time `json:"open_timestamp"`
			High           float64   `json:"high"`
			HighTimestamp  time.Time `json:"high_timestamp"`
			Low            float64   `json:"low"`
			LowTimestamp   time.Time `json:"low_timestamp"`
			Close          float64   `json:"close"`
			CloseTimestamp time.Time `json:"close_timestamp"`
			PercentChange  float64   `json:"percent_change"`
			PriceChange    float64   `json:"price_change"`
		} `json:"quote"`
	} `json:"periods"`
}