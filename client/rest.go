package client

import (
	"context"
	"net"
	"net/http"
	"time"

	"github.com/cpj555/go-coinmarketcap/errs"
	"github.com/cpj555/go-coinmarketcap/network"
	"github.com/cpj555/go-coinmarketcap/tools"
	"github.com/cpj555/go-coinmarketcap/types"
	"github.com/go-resty/resty/v2"
)

// setupResty setup resty client
func (c *client) setupResty() {
	c.r = resty.New().
		SetTransport(createTransport(nil, 500)).
		SetDebug(c.conf.IsDebug).
		SetTimeout(c.conf.Timeout).
		SetHeader("X-CMC_PRO_API_KEY", c.conf.ApiKey).
		OnAfterResponse(
			func(r *resty.Client, resp *resty.Response) error {
				if !network.IsSuccessResponse(resp.StatusCode()) {
					return errs.New(resp.StatusCode(), string(resp.Body()))
				}
				rsp := c.unmarshalResult(resp)
				if rsp.Status.ErrorCode != network.OpenAPIStatusOK {
					return errs.New(rsp.Status.ErrorCode, rsp.Status.ErrorMessage)
				}
				return nil
			},
		)
	if c.conf.ProxyUrl != "" {
		c.r.SetProxy(c.conf.ProxyUrl)
	}

	c.r.JSONMarshal = tools.JSON.Marshal
	c.r.JSONUnmarshal = tools.JSON.Unmarshal
}

// request you should create a request object before doing each HTTP request
func (c *client) request(ctx context.Context) *resty.Request {
	return c.r.R().
		SetContext(ctx).
		SetResult(types.OpenAPIRsp{})
}

// unmarshalResult get model.OpenAPIRsp result from the response
func (c *client) unmarshalResult(resp *resty.Response) *types.OpenAPIRsp {
	return resp.Result().(*types.OpenAPIRsp)
}

// createTransport customize transport
func createTransport(addr net.Addr, idleConnections int) *http.Transport {
	dialer := &net.Dialer{
		Timeout:   60 * time.Second,
		KeepAlive: 60 * time.Second,
	}
	if addr != nil {
		dialer.LocalAddr = addr
	}
	return &http.Transport{
		Proxy:                 http.ProxyFromEnvironment,
		DialContext:           dialer.DialContext,
		ForceAttemptHTTP2:     false,
		MaxIdleConns:          idleConnections,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
		MaxIdleConnsPerHost:   idleConnections,
		MaxConnsPerHost:       idleConnections,
	}
}
