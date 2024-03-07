package client

import (
	"context"

	"github.com/cpj555/go-coinmarketcap/tools"
	"github.com/cpj555/go-coinmarketcap/types"
)

// GetMap Returns a mapping of all cryptocurrencies to unique CoinMarketCap ids
func (c *client) GetMap(ctx context.Context, req *types.GetMapReq) (*types.GetMapResp, error) {
	if err := req.ValidParams(); err != nil {
		return nil, err
	}

	resp, err := c.request(ctx).SetQueryParams(nil).Get(c.getApi(getMapUri))
	if err != nil {
		return nil, err
	}

	result := &types.GetMapResp{}
	if err = tools.JSON.Unmarshal(c.unmarshalResult(resp).Data, &result); err != nil {
		return nil, err
	}
	return result, nil
}
