package client

import (
	"errors"
	"time"
)

type OptionHandler func(*Config) error

// WithApiKey Customize DoDoBot Api base host
func WithApiKey(apiKey string) OptionHandler {
	return func(config *Config) error {
		config.ApiKey = apiKey
		return nil
	}
}

// WithBaseApi Customize DoDoBot Api base host
func WithBaseApi(baseApi string) OptionHandler {
	return func(config *Config) error {
		if baseApi == "" {
			return errors.New("invalid BaseApi (empty string detected)")
		}
		config.BaseApi = baseApi
		return nil
	}
}

// WithTimeout Customize RestyClient request timeout
func WithTimeout(duration time.Duration) OptionHandler {
	return func(config *Config) error {
		config.Timeout = duration
		return nil
	}
}

func WithProxyUrl(proxyUrl string) OptionHandler {
	return func(config *Config) error {
		config.ProxyUrl = proxyUrl
		return nil
	}
}

// WithSandBox Toggle sandbox mode
func WithSandBox(flag bool) OptionHandler {
	return func(config *Config) error {
		config.IsSandBox = flag
		return nil
	}
}

// WithDebugMode Toggle debug mode
func WithDebugMode(flag bool) OptionHandler {
	return func(config *Config) error {
		config.IsDebug = flag
		return nil
	}
}
