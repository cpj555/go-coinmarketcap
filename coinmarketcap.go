package go_coinmarketcap

import "github.com/cpj555/go-coinmarketcap/client"

func NewClient(options ...client.OptionHandler) (*client.Client, error) {
	return client.New(options...)
}
