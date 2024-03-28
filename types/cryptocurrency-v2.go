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

type GetCryptocurrencyInfoResp map[string]InfoItem

type InfoItem struct {
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
	SelfReportedCirculatingSupply *float64  `json:"self_reported_circulating_supply"`
	SelfReportedMarketCap         *float64  `json:"self_reported_market_cap"`
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

type GetCryptocurrencyQuoteResp map[string]QuoteItem

type QuoteItem struct {
	Id                int       `json:"id"`
	Name              string    `json:"name"`
	Symbol            string    `json:"symbol"`
	Slug              string    `json:"slug"`
	IsActive          int       `json:"is_active"`
	IsFiat            int       `json:"is_fiat"`
	CirculatingSupply float64   `json:"circulating_supply"`
	TotalSupply       float64   `json:"total_supply"`
	MaxSupply         float64   `json:"max_supply"`
	DateAdded         time.Time `json:"date_added"`
	NumMarketPairs    int       `json:"num_market_pairs"`
	CmcRank           int       `json:"cmc_rank"`
	LastUpdated       time.Time `json:"last_updated"`
	Tags              *[]struct {
		Name     string `json:"name"`
		Slug     string `json:"slug"`
		Category string `json:"category"`
	} `json:"tags"`
	Platform *struct {
		Id           int    `json:"id"`
		Name         string `json:"name"`
		Symbol       string `json:"symbol"`
		Slug         string `json:"slug"`
		TokenAddress string `json:"token_address"`
	} `json:"platform"`
	SelfReportedCirculatingSupply *float64 `json:"self_reported_circulating_supply"`
	SelfReportedMarketCap         *float64 `json:"self_reported_market_cap"`
	Quote                         map[string]struct {
		Price                 float64   `json:"price"`
		Volume24H             float64   `json:"volume_24h"`
		VolumeChange24H       float64   `json:"volume_change_24h"`
		PercentChange1H       float64   `json:"percent_change_1h"`
		PercentChange24H      float64   `json:"percent_change_24h"`
		PercentChange7D       float64   `json:"percent_change_7d"`
		PercentChange30D      float64   `json:"percent_change_30d"`
		MarketCap             float64   `json:"market_cap"`
		MarketCapDominance    float64   `json:"market_cap_dominance"`
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

type GetCryptocurrencyPricePerformanceStatsResp map[string]CryptocurrencyPricePerformanceStatsItem

type CryptocurrencyPricePerformanceStatsItem struct {
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

type GetCryptocurrencyMarketPairReq struct {
	Id            string `schema:"id,omitempty"`
	Slug          string `schema:"slug,omitempty"`
	Symbol        string `schema:"symbol,omitempty"`
	Start         int    `schema:"start,omitempty,default:1"`
	Limit         int    `schema:"limit,omitempty"`
	SortDir       string `schema:"sort_dir,omitempty"`
	Sort          string `schema:"sort,omitempty"`
	Aux           string `schema:"aux,omitempty"`
	MatchedId     string `schema:"matched_id,omitempty"`
	MatchedSymbol string `schema:"matched_symbol,omitempty"`
	Category      string `schema:"category,omitempty"`
	FeeType       string `schema:"fee_type,omitempty"`
	Convert       string `schema:"convert,omitempty"`
	ConvertId     string `schema:"convert_id,omitempty"`
}

type GetCryptocurrencyMarketPairResp struct {
	Id             int    `json:"id"`
	Name           string `json:"name"`
	Symbol         string `json:"symbol"`
	NumMarketPairs int    `json:"num_market_pairs"`
	MarketPairs    []struct {
		Exchange struct {
			Id   int    `json:"id"`
			Name string `json:"name"`
			Slug string `json:"slug"`
		} `json:"exchange"`
		MarketId       int    `json:"market_id"`
		MarketPair     string `json:"market_pair"`
		Category       string `json:"category"`
		FeeType        string `json:"fee_type"`
		MarketPairBase struct {
			CurrencyId     int    `json:"currency_id"`
			CurrencySymbol string `json:"currency_symbol"`
			ExchangeSymbol string `json:"exchange_symbol"`
			CurrencyType   string `json:"currency_type"`
		} `json:"market_pair_base"`
		MarketPairQuote struct {
			CurrencyId     int    `json:"currency_id"`
			CurrencySymbol string `json:"currency_symbol"`
			ExchangeSymbol string `json:"exchange_symbol"`
			CurrencyType   string `json:"currency_type"`
		} `json:"market_pair_quote"`
		Quote map[string]struct {
			Price          float64   `json:"price"`
			Volume24HBase  float64   `json:"volume_24h_base"`
			Volume24HQuote float64   `json:"volume_24h_quote"`
			LastUpdated    time.Time `json:"last_updated"`
		} `json:"quote"`
	} `json:"market_pairs"`
}
