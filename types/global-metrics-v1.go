package types

import (
	"encoding/json"
	"time"
)

type GetGlobalMetricsQuotesReq struct {
	Convert   string `schema:"convert,omitempty"`
	ConvertId string `schema:"convert_id,omitempty"`
}

type GetGlobalMetricsQuotesResp struct {
	EthDominanceYesterday           json.Number `json:"eth_dominance_yesterday"`
	BtcDominanceYesterday           json.Number `json:"btc_dominance_yesterday"`
	EthDominance24HPercentageChange json.Number `json:"eth_dominance_24h_percentage_change"`
	BtcDominance24HPercentageChange json.Number `json:"btc_dominance_24h_percentage_change"`
	DefiVolume24H                   json.Number `json:"defi_volume_24h"`
	DefiVolume24HReported           json.Number `json:"defi_volume_24h_reported"`
	DefiMarketCap                   json.Number `json:"defi_market_cap"`
	Defi24HPercentageChange         json.Number `json:"defi_24h_percentage_change"`
	StablecoinVolume24H             json.Number `json:"stablecoin_volume_24h"`
	StablecoinVolume24HReported     json.Number `json:"stablecoin_volume_24h_reported"`
	StablecoinMarketCap             json.Number `json:"stablecoin_market_cap"`
	Stablecoin24HPercentageChange   json.Number `json:"stablecoin_24h_percentage_change"`
	DerivativesVolume24H            json.Number `json:"derivatives_volume_24h"`
	DerivativesVolume24HReported    json.Number `json:"derivatives_volume_24h_reported"`
	Derivatives24HPercentageChange  json.Number `json:"derivatives_24h_percentage_change"`
	BtcDominance                    json.Number `json:"btc_dominance"`           // Bitcoin's market dominance percentage by market cap.
	EthDominance                    json.Number `json:"eth_dominance"`           // Ethereum's market dominance percentage by market cap.
	ActiveCryptocurrencies          int         `json:"active_cryptocurrencies"` // Count of active cryptocurrencies tracked by CoinMarketCap. This includes all cryptocurrencies with a listing_status of "active" or "listed" as returned from our /cryptocurrency/map call.
	TotalCryptocurrencies           int         `json:"total_cryptocurrencies"`  // Count of all cryptocurrencies tracked by CoinMarketCap. This includes "inactive" listing_status cryptocurrencies.
	ActiveMarketPairs               int         `json:"active_market_pairs"`     // Count of active market pairs tracked by CoinMarketCap across all exchanges.
	ActiveExchanges                 int         `json:"active_exchanges"`        // Count of active exchanges tracked by CoinMarketCap. This includes all exchanges with a listing_status of "active" or "listed" as returned by our /exchange/map call.
	TotalExchanges                  int         `json:"total_exchanges"`         // Count of all exchanges tracked by CoinMarketCap. This includes "inactive" listing_status exchanges.
	LastUpdated                     time.Time   `json:"last_updated"`            // Timestamp of when this record was last updated.
	Quote                           map[string]struct {
		TotalMarketCap                          json.Number `json:"total_market_cap"`            // The sum of all individual cryptocurrency market capitalizations in the requested currency.
		TotalVolume24H                          json.Number `json:"total_volume_24h"`            // The sum of rolling 24 hour adjusted volume (as outlined in our methodology) for all cryptocurrencies in the requested currency.
		TotalVolume24HReported                  json.Number `json:"total_volume_24h_reported"`   // The sum of rolling 24 hour reported volume for all cryptocurrencies in the requested currency.
		AltcoinVolume24H                        json.Number `json:"altcoin_volume_24h"`          // The sum of rolling 24 hour adjusted volume (as outlined in our methodology) for all cryptocurrencies excluding Bitcoin in the requested currency.
		AltcoinVolume24HReported                json.Number `json:"altcoin_volume_24h_reported"` // The sum of rolling 24 hour reported volume for all cryptocurrencies excluding Bitcoin in the requested currency.
		AltcoinMarketCap                        json.Number `json:"altcoin_market_cap"`          // The sum of all individual cryptocurrency market capitalizations excluding Bitcoin in the requested currency.
		DefiVolume24H                           json.Number `json:"defi_volume_24h"`
		DefiVolume24HReported                   json.Number `json:"defi_volume_24h_reported"`
		Defi24HPercentageChange                 json.Number `json:"defi_24h_percentage_change"`
		DefiMarketCap                           json.Number `json:"defi_market_cap"`
		StablecoinVolume24H                     json.Number `json:"stablecoin_volume_24h"`
		StablecoinVolume24HReported             json.Number `json:"stablecoin_volume_24h_reported"`
		Stablecoin24HPercentageChange           json.Number `json:"stablecoin_24h_percentage_change"`
		StablecoinMarketCap                     json.Number `json:"stablecoin_market_cap"`
		DerivativesVolume24H                    json.Number `json:"derivatives_volume_24h"`
		DerivativesVolume24HReported            json.Number `json:"derivatives_volume_24h_reported"`
		Derivatives24HPercentageChange          json.Number `json:"derivatives_24h_percentage_change"`
		LastUpdated                             time.Time   `json:"last_updated"` // Timestamp (ISO 8601) of when the conversion currency's current value was referenced.
		TotalMarketCapYesterday                 json.Number `json:"total_market_cap_yesterday"`
		TotalVolume24HYesterday                 json.Number `json:"total_volume_24h_yesterday"`
		TotalMarketCapYesterdayPercentageChange json.Number `json:"total_market_cap_yesterday_percentage_change"`
		TotalVolume24HYesterdayPercentageChange json.Number `json:"total_volume_24h_yesterday_percentage_change"`
	} `json:"quote"`
}
