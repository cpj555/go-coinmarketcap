package client

import (
	"context"
	"github.com/cpj555/go-coinmarketcap/tools"
	"github.com/cpj555/go-coinmarketcap/types"
)

type fiatV1 struct {
	cli *Client
}

func newFiatV1(client *Client) FiatV1API {
	return &fiatV1{cli: client}
}

func (e *fiatV1) GetMap(ctx context.Context, req *types.GetFiatMapReq) (*types.GetFiatResp, error) {

	values := tools.ToUrlValues(req)

	resp, err := e.cli.request(ctx).SetQueryParamsFromValues(values).Get(e.cli.getApi(getFiatMapUri))
	if err != nil {
		return nil, err
	}

	result := &types.GetFiatResp{}
	if err = tools.JSON.Unmarshal(e.cli.unmarshalResult(resp).Data, &result); err != nil {
		return nil, err
	}
	return result, nil
}
