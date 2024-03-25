package types

import (
	"errors"
	"time"
)

type GetExchangeMapReq struct {
	ListingStatus ListingStatus `schema:"listing_status,omitempty,default:active"`
	Start         int           `schema:"start,omitempty,default:1"`
	Limit         int           `schema:"limit,omitempty"`
	Sort          Sort          `schema:"sort,omitempty"`
	Symbol        string        `schema:"symbol,omitempty"`
	Aux           string        `schema:"aux,omitempty"`
}

func (p *GetExchangeMapReq) ValidParams() error {
	if p.Start <= 0 {
		return errors.New("invalid parameter Start (Start >= 1)")
	}
	if p.Limit <= 0 || p.Limit > 5000 {
		return errors.New("invalid parameter Limit (0 < Limit <= 5000)")
	}
	return nil
}

type GetExchangeMapResp []struct {
	Id                  int       `json:"id"`
	Name                string    `json:"name"`
	Slug                string    `json:"slug"`
	IsActive            int       `json:"is_active"`
	Status              string    `json:"status"`
	FirstHistoricalData time.Time `json:"first_historical_data"`
	LastHistoricalData  time.Time `json:"last_historical_data"`
}

type GetExchangeInfoReq struct {
	Id   string `schema:"id,omitempty"`
	Slug string `schema:"slug,omitempty"`
	Aux  string `schema:"aux,omitempty"`
}

type GetExchangeInfoResp map[string]struct {
	Id           int       `json:"id"`
	Name         string    `json:"name"`
	Slug         string    `json:"slug"`
	Logo         string    `json:"logo"`
	Description  string    `json:"description"`
	DateLaunched time.Time `json:"date_launched"`
	Notice       *string   `json:"notice"`
	Countries    *[]string `json:"countries"`
	Fiats        []string  `json:"fiats"`
	Tags         *[]struct {
		Name string `json:"name"`
		Slug string `json:"slug"`
	} `json:"tags"`
	Type                  string    `json:"type"`
	MakerFee              float64   `json:"maker_fee"`
	TakerFee              float64   `json:"taker_fee"`
	WeeklyVisits          int       `json:"weekly_visits"`
	SpotVolumeUsd         float64   `json:"spot_volume_usd"`
	SpotVolumeLastUpdated time.Time `json:"spot_volume_last_updated"`
	Urls                  struct {
		Website []string      `json:"website"`
		Twitter []string      `json:"twitter"`
		Blog    []interface{} `json:"blog"`
		Chat    []string      `json:"chat"`
		Fee     []string      `json:"fee"`
	} `json:"urls"`
}

type GetExchangeQuotesReq struct {
	Id        string `schema:"id,omitempty"`
	Slug      string `schema:"slug,omitempty"`
	Convert   string `schema:"convert,omitempty"`
	ConvertId string `schema:"convert_id,omitempty"`
	Aux       string `schema:"aux,omitempty"`
}

type GetExchangeQuotesResp map[string]struct {
	Id             int       `json:"id"`
	Name           string    `json:"name"`
	Slug           string    `json:"slug"`
	NumCoins       int       `json:"num_coins"`
	NumMarketPairs int       `json:"num_market_pairs"`
	LastUpdated    time.Time `json:"last_updated"`
	TrafficScore   int       `json:"traffic_score"`
	Rank           int       `json:"rank"`
	ExchangeScore  float64   `json:"exchange_score"`
	LiquidityScore float64   `json:"liquidity_score"`
	Quote          map[string]struct {
		Volume24H              float64 `json:"volume_24h"`
		Volume24HAdjusted      float64 `json:"volume_24h_adjusted"`
		Volume7D               int64   `json:"volume_7d"`
		Volume30D              int64   `json:"volume_30d"`
		PercentChangeVolume24H float64 `json:"percent_change_volume_24h"`
		PercentChangeVolume7D  float64 `json:"percent_change_volume_7d"`
		PercentChangeVolume30D float64 `json:"percent_change_volume_30d"`
		EffectiveLiquidity24H  float64 `json:"effective_liquidity_24h"`
	} `json:"quote"`
}
