package client

import (
	"context"
	"net/url"

	"github.com/cpj555/go-coinmarketcap/tools"
	"github.com/cpj555/go-coinmarketcap/types"
	"github.com/gorilla/schema"
	_ "github.com/mitchellh/mapstructure"
)

// GetMap Returns a mapping of all cryptocurrencies to unique CoinMarketCap ids
func (c *client) GetMap(ctx context.Context, req *types.GetMapReq) (*types.GetMapResp, error) {
	if err := req.ValidParams(); err != nil {
		return nil, err
	}
	encoder := schema.NewEncoder()
	values := url.Values{}
	encoder.Encode(req, values)

	resp, err := c.request(ctx).SetQueryParamsFromValues(values).Get(c.getApi(getMapUri))
	if err != nil {
		return nil, err
	}

	result := &types.GetMapResp{}
	if err = tools.JSON.Unmarshal(c.unmarshalResult(resp).Data, &result); err != nil {
		return nil, err
	}
	return result, nil
}
