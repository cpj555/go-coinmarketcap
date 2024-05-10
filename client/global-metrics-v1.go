package client

import (
	"context"
	"github.com/cpj555/go-coinmarketcap/tools"
	"github.com/cpj555/go-coinmarketcap/types"
)

type globalMetricsV1 struct {
	cli *Client
}

func newGlobalMetricsV1(client *Client) GlobalMetricsV1API {
	return &globalMetricsV1{
		cli: client,
	}
}

func (g *globalMetricsV1) GetQuotesLatest(ctx context.Context, req *types.GetGlobalMetricsQuotesReq) (*types.GetGlobalMetricsQuotesResp, error) {
	values := tools.ToUrlValues(req)

	resp, err := g.cli.request(ctx).SetQueryParamsFromValues(values).Get(g.cli.getApi(getGlobalMetricsQuotesLatestUri))
	if err != nil {
		return nil, err
	}

	result := &types.GetGlobalMetricsQuotesResp{}
	if err = tools.JSON.Unmarshal(g.cli.unmarshalResult(resp).Data, &result); err != nil {
		return nil, err
	}
	return result, nil
}
