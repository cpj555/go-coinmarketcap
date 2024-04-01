package types

import (
	"errors"
	"time"
)

type ListingStatus string

type Sort string

// GetMapResp CoinMarketCap ID Map
// [https://coinmarketcap.com/api/documentation/v1/#operation/getV1CryptocurrencyMap]
type GetCryptocurrencyMapReq struct {
	ListingStatus ListingStatus `schema:"listing_status,omitempty,default:active"`
	Start         int           `schema:"start"`
	Limit         int           `schema:"limit"`
	Sort          Sort          `schema:"sort,omitempty"`
	Symbol        string        `schema:"symbol,omitempty"`
	Aux           string        `schema:"aux,omitempty"`
}

type GetCryptocurrencyMapResp []struct {
	Id                  int       `json:"id"`
	Rank                int       `json:"rank"`
	Name                string    `json:"name"`
	Symbol              string    `json:"symbol"`
	Slug                string    `json:"slug"`
	IsActive            int       `json:"is_active"`
	Status              string    `json:"status"`
	FirstHistoricalData time.Time `json:"first_historical_data"`
	LastHistoricalData  time.Time `json:"last_historical_data"`
	Platform            *struct {
		Id           int    `json:"id"`
		Name         string `json:"name"`
		Symbol       string `json:"symbol"`
		Slug         string `json:"slug"`
		TokenAddress string `json:"token_address"`
	} `json:"platform"`
}

func (p *GetCryptocurrencyMapReq) ValidParams() error {
	if p.Start <= 0 {
		return errors.New("invalid parameter Start (Start >= 1)")
	}
	if p.Limit <= 0 || p.Limit > 5000 {
		return errors.New("invalid parameter Limit (0 < Limit <= 5000)")
	}
	return nil
}
