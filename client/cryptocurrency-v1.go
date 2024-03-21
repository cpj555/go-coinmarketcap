package client

import (
	"context"
	"github.com/cpj555/go-coinmarketcap/tools"
	"github.com/cpj555/go-coinmarketcap/types"
)

type cryptocurrencyV1 struct {
	cli *Client
}

func newCryptocurrencyV1(client *Client) CryptocurrencyV1API {
	return &cryptocurrencyV1{cli: client}
}

// GetMap Returns a mapping of all cryptocurrencies to unique CoinMarketCap ids
func (c *cryptocurrencyV1) GetMap(ctx context.Context, req *types.GetCryptocurrencyMapReq) (*types.GetCryptocurrencyMapResp, error) {
	if err := req.ValidParams(); err != nil {
		return nil, err
	}

	values := tools.ToUrlValues(req)

	resp, err := c.cli.request(ctx).SetQueryParamsFromValues(values).Get(c.cli.getApi(getCryptocurrencyMapUri))
	if err != nil {
		return nil, err
	}

	result := &types.GetCryptocurrencyMapResp{}
	if err = tools.JSON.Unmarshal(c.cli.unmarshalResult(resp).Data, &result); err != nil {
		return nil, err
	}
	return result, nil
}
