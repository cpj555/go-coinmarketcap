package client

import (
	"context"
	"net/url"

	"github.com/cpj555/go-coinmarketcap/tools"
	"github.com/cpj555/go-coinmarketcap/types"
	"github.com/gorilla/schema"
	_ "github.com/mitchellh/mapstructure"
)

// GetInfo Returns all static metadata available for one or more cryptocurrencies
func (c *client) GetInfo(ctx context.Context, req *types.GetInfoReq) (*types.GetInfoResp, error) {
	if err := req.ValidParams(); err != nil {
		return nil, err
	}
	encoder := schema.NewEncoder()
	values := url.Values{}
	encoder.Encode(req, values)

	resp, err := c.request(ctx).SetQueryParamsFromValues(values).Get(c.getApi(getInfoUri))
	if err != nil {
		return nil, err
	}

	result := &types.GetInfoResp{}
	if err = tools.JSON.Unmarshal(c.unmarshalResult(resp).Data, &result); err != nil {
		return nil, err
	}
	return result, nil
}
