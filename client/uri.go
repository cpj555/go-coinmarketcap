package client

import "fmt"

type uri string

const (

	// cryptocurrency API
	getCryptocurrencyMapUri                         uri = "/v1/cryptocurrency/map"
	getCryptocurrencyInfoUri                        uri = "/v2/cryptocurrency/info"
	getCryptocurrencyQuotesLatest                   uri = "/v2/cryptocurrency/quotes/latest"
	getCryptocurrencyPricePerformanceStatsLatestUri uri = "/v2/cryptocurrency/price-performance-stats/latest"
	getCryptocurrencyMarketPairUri                  uri = "/v2/cryptocurrency/market-pairs/latest"
	// exchange API
	getExchangeMapUri          uri = "/v1/exchange/map"
	getExchangeInfoUri         uri = "/v1/exchange/info"
	getExchangeQuotesLatestUri uri = "/v1/exchange/quotes/latest"
	getExchangeMarketPairUri   uri = "/v1/exchange/market-pairs/latest"

	getFiatMapUri uri = "/v1/fiat/map"

	getGlobalMetricsQuotesLatestUri uri = "/v1/global-metrics/quotes/latest"
)

// getApi build the full api url
func (c *Client) getApi(u uri) string {
	return fmt.Sprintf("%s%s", c.conf.BaseApi, u)
}
