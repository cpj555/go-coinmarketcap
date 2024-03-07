package client

import (
	"context"
	"errors"
	"github.com/cpj555/go-coinmarketcap/types"
	"github.com/go-resty/resty/v2"
	"time"
)

type (
	Client interface {
		Base
		CryptocurrencyV1API
	}

	// Base client basic interface
	Base interface {
		GetConfig() *Config
	}

	CryptocurrencyV1API interface {
		GetMap(ctx context.Context, req *types.GetMapReq) (*types.GetMapResp, error)
	}
)
type (
	client struct {
		conf *Config
		r    *resty.Client
	}

	// Config coinmarketcap client configuration
	Config struct {
		BaseApi   string // coinmarketcap OpenAPI Server Domain
		ApiKey    string // coinmarketcap ApiKey
		IsDebug   bool   // debug mode
		IsSandBox bool
		Timeout   time.Duration // resty client request timeout
	}
)

// New create a new coinmarket instance
func New(options ...OptionHandler) (Client, error) {
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

	instance := &client{conf: config}
	instance.setupResty()

	return instance, nil
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
func (c *client) GetConfig() *Config {
	return c.conf
}
