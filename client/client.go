package client

import (
	"context"
	"errors"
	"time"

	"github.com/cpj555/go-coinmarketcap/types"
	"github.com/go-resty/resty/v2"
)

type (

	// Base client basic interface
	Base interface {
		GetConfig() *Config
	}

	// 币种接口
	CryptocurrencyV1API interface {
		GetMap(ctx context.Context, req *types.GetCryptocurrencyMapReq) (*types.GetCryptocurrencyMapResp, error)
	}

	// 币种接口
	CryptocurrencyV2API interface {
		GetInfo(ctx context.Context, req *types.GetCryptocurrencyInfoReq) (*types.GetCryptocurrencyInfoResp, error)
		GetQuotesLatest(ctx context.Context, req *types.GetCryptocurrencyQuotesReq) (*types.GetCryptocurrencyQuoteResp, error)
		GetPricePerformanceStatsLatest(ctx context.Context, req *types.GetCryptocurrencyPricePerformanceStatsReq) (*types.GetCryptocurrencyPricePerformanceStatsResp, error)
		GetMarketPairLatest(ctx context.Context, req *types.GetCryptocurrencyMarketPairReq) (*types.GetCryptocurrencyMarketPairResp, error)
	}

	// 交易所接口
	ExchangeV1API interface {
		GetMap(ctx context.Context, req *types.GetExchangeMapReq) (*types.GetExchangeMapResp, error)
		GetInfo(ctx context.Context, req *types.GetExchangeInfoReq) (*types.GetExchangeInfoResp, error)
		GetQuotesLatest(ctx context.Context, req *types.GetExchangeQuotesReq) (*types.GetExchangeQuotesResp, error)
		GetMarketPairLatest(ctx context.Context, req *types.GetExchangeMarketPairReq) (*types.GetExchangeMarketPairResp, error)
	}
)

type (
	Client struct {
		conf             *Config
		CryptocurrencyV1 CryptocurrencyV1API
		CryptocurrencyV2 CryptocurrencyV2API
		ExchangeV1       ExchangeV1API
		r                *resty.Client
	}

	// Config coinmarketcap client configuration
	Config struct {
		BaseApi     string // coinmarketcap OpenAPI Server Domain
		RestyClient *resty.Client
		ApiKey      string // coinmarketcap ApiKey
		IsDebug     bool   // debug mode
		IsSandBox   bool
		Timeout     time.Duration // resty client request timeout
		ProxyUrl    string        // 国内访问不了设置下代理
	}
)

// New create a new coinmarket instance
func New(options ...OptionHandler) (*Client, error) {
	config := getDefaultConfig()

	// handle custom options
	for _, optionHandler := range options {
		if optionHandler == nil {
			return nil, errors.New("invalid OptionHandler (nil detected)")
		}
		if err := optionHandler(config); err != nil {
			return nil, err
		}
	}
	if config.IsSandBox {
		config.BaseApi = "https://sandbox-api.coinmarketcap.com"
		config.ApiKey = "b54bcf4d-1bca-4e8e-9a24-22ff2c3d462c"
	}

	instance := &Client{conf: config}
	if config.RestyClient != nil {
		instance.r = config.RestyClient
	}
	instance.setupResty()

	initAPI(instance)

	return instance, nil
}

func initAPI(client *Client) {
	client.CryptocurrencyV1 = newCryptocurrencyV1(client)
	client.CryptocurrencyV2 = newCryptocurrencyV2(client)
	client.ExchangeV1 = newExchangeV1(client)
}

// getDefaultConfig Get the default configuration
func getDefaultConfig() *Config {
	return &Config{
		BaseApi: "https://pro-api.coinmarketcap.com",
		IsDebug: false,
		Timeout: time.Second * 5,
	}
}

// GetConfig get instance configuration
func (c *Client) GetConfig() *Config {
	return c.conf
}
