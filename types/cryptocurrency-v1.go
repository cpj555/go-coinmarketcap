package types

import (
	"errors"
	"time"
)

type ListingStatus string

type Sort string

// default active
const (
	active    ListingStatus = "active"
	inactive  ListingStatus = "inactive"
	untracked ListingStatus = "untracked"

	id      Sort = "id"
	cmcRank Sort = "cmc_rank"
)

// GetMapResp CoinMarketCap ID Map
// [https://coinmarketcap.com/api/documentation/v1/#operation/getV1CryptocurrencyMap]
type GetMapResp struct {
	Data []struct {
		Id                  int       `json:"id"`
		Rank                int       `json:"rank"`
		Name                string    `json:"name"`
		Symbol              string    `json:"symbol"`
		Slug                string    `json:"slug"`
		IsActive            int       `json:"is_active"`
		FirstHistoricalData time.Time `json:"first_historical_data"`
		LastHistoricalData  time.Time `json:"last_historical_data"`
		Platform            struct {
			Id           int    `json:"id"`
			Name         string `json:"name"`
			Symbol       string `json:"symbol"`
			Slug         string `json:"slug"`
			TokenAddress string `json:"token_address"`
		} `json:"platform"`
	} `json:"data"`
}

type (
	GetMapReq struct {
		Start  int
		Limit  int
		Sort   Sort
		Symbol string
		aux    string
	}
)

func (p *GetMapReq) ValidParams() error {
	if p.Start <= 0 {
		return errors.New("invalid parameter Start (Start >= 1)")
	}
	if p.Limit <= 0 || p.Limit > 5000 {
		return errors.New("invalid parameter Limit (0 < Limit <= 5000)")
	}
	return nil
}
