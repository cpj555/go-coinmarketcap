package client

import (
	"context"
)

type (
	Client interface {
		CryptocurrencyV1API
	}

	CryptocurrencyV1API interface {
		Map(ctx context.Context)
	}
)
