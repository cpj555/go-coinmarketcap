package client

import (
	"context"
	"github.com/cpj555/go-coinmarketcap/tools"
	"github.com/cpj555/go-coinmarketcap/types"
)

type exchangeV1 struct {
	cli *Client
}

func newExchangeV1(client *Client) ExchangeV1API {
	return &exchangeV1{cli: client}
}

func (e exchangeV1) GetMap(ctx context.Context, req *types.GetExchangeMapReq) (*types.GetExchangeMapResp, error) {

	if err := req.ValidParams(); err != nil {
		return nil, err
	}
	values := tools.ToUrlValues(req)

	resp, err := e.cli.request(ctx).SetQueryParamsFromValues(values).Get(e.cli.getApi(getExchangeMapUri))
	if err != nil {
		return nil, err
	}

	result := &types.GetExchangeMapResp{}
	if err = tools.JSON.Unmarshal(e.cli.unmarshalResult(resp).Data, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (e exchangeV1) GetInfo(ctx context.Context, req *types.GetExchangeInfoReq) (*types.GetExchangeInfoResp, error) {
	values := tools.ToUrlValues(req)

	resp, err := e.cli.request(ctx).SetQueryParamsFromValues(values).Get(e.cli.getApi(getExchangeInfoUri))
	if err != nil {
		return nil, err
	}

	result := &types.GetExchangeInfoResp{}
	if err = tools.JSON.Unmarshal(e.cli.unmarshalResult(resp).Data, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (e exchangeV1) GetQuotesLatest(ctx context.Context, req *types.GetExchangeQuotesReq) (*types.GetExchangeQuotesResp, error) {
	values := tools.ToUrlValues(req)

	resp, err := e.cli.request(ctx).SetQueryParamsFromValues(values).Get(e.cli.getApi(getExchangeQuotesLatestUri))
	if err != nil {
		return nil, err
	}

	result := &types.GetExchangeQuotesResp{}
	if err = tools.JSON.Unmarshal(e.cli.unmarshalResult(resp).Data, &result); err != nil {
		return nil, err
	}
	return result, nil
}
