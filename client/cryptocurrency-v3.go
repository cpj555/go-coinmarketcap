package client

import (
	"context"
	"github.com/cpj555/go-coinmarketcap/tools"
	"github.com/cpj555/go-coinmarketcap/types"
)

type cryptocurrencyV3 struct {
	cli *Client
}

func newCryptocurrencyV3(client *Client) CryptocurrencyV3API {
	return &cryptocurrencyV3{cli: client}
}

func (c *cryptocurrencyV3) GetQuotesHistorical(ctx context.Context, req *types.GetCryptocurrencyQuotesHistoricalReq) (*types.GetCryptocurrencyQuoteHistoricalResp, error) {
	values := tools.ToUrlValues(req)

	resp, err := c.cli.request(ctx).SetQueryParamsFromValues(values).Get(c.cli.getApi(getCryptocurrencyQuotesHistoricalUri))
	if err != nil {
		return nil, err
	}

	result := &types.GetCryptocurrencyQuoteHistoricalResp{}
	if err = tools.JSON.Unmarshal(c.cli.unmarshalResult(resp).Data, &result); err != nil {
		return nil, err
	}
	return result, nil
}
