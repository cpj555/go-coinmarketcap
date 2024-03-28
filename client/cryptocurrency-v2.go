package client

import (
	"context"
	"github.com/cpj555/go-coinmarketcap/tools"
	"github.com/cpj555/go-coinmarketcap/types"
)

type cryptocurrencyV2 struct {
	cli *Client
}

func newCryptocurrencyV2(client *Client) CryptocurrencyV2API {
	return &cryptocurrencyV2{cli: client}
}

// GetInfo Returns all static metadata available for one or more cryptocurrencies
func (c *cryptocurrencyV2) GetInfo(ctx context.Context, req *types.GetCryptocurrencyInfoReq) (*types.GetCryptocurrencyInfoResp, error) {

	values := tools.ToUrlValues(req)

	resp, err := c.cli.request(ctx).SetQueryParamsFromValues(values).Get(c.cli.getApi(getCryptocurrencyInfoUri))
	if err != nil {
		return nil, err
	}

	result := &types.GetCryptocurrencyInfoResp{}
	if err = tools.JSON.Unmarshal(c.cli.unmarshalResult(resp).Data, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *cryptocurrencyV2) GetQuotesLatest(ctx context.Context, req *types.GetCryptocurrencyQuotesReq) (*types.GetCryptocurrencyQuoteResp, error) {

	values := tools.ToUrlValues(req)

	resp, err := c.cli.request(ctx).SetQueryParamsFromValues(values).Get(c.cli.getApi(getCryptocurrencyQuotesLatest))
	if err != nil {
		return nil, err
	}

	result := &types.GetCryptocurrencyQuoteResp{}
	if err = tools.JSON.Unmarshal(c.cli.unmarshalResult(resp).Data, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *cryptocurrencyV2) GetPricePerformanceStatsLatest(ctx context.Context, req *types.GetCryptocurrencyPricePerformanceStatsReq) (*types.GetCryptocurrencyPricePerformanceStatsResp, error) {
	values := tools.ToUrlValues(req)

	resp, err := c.cli.request(ctx).SetQueryParamsFromValues(values).Get(c.cli.getApi(getCryptocurrencyPricePerformanceStatsLatestUri))
	if err != nil {
		return nil, err
	}

	result := &types.GetCryptocurrencyPricePerformanceStatsResp{}
	if err = tools.JSON.Unmarshal(c.cli.unmarshalResult(resp).Data, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *cryptocurrencyV2) GetMarketPairLatest(ctx context.Context, req *types.GetCryptocurrencyMarketPairReq) (*types.GetCryptocurrencyMarketPairResp, error) {
	values := tools.ToUrlValues(req)

	resp, err := c.cli.request(ctx).SetQueryParamsFromValues(values).Get(c.cli.getApi(getCryptocurrencyMarketPairUri))
	if err != nil {
		return nil, err
	}

	result := &types.GetCryptocurrencyMarketPairResp{}
	if err = tools.JSON.Unmarshal(c.cli.unmarshalResult(resp).Data, &result); err != nil {
		return nil, err
	}
	return result, nil
}
