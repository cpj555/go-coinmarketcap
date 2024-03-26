package types

import (
	"time"

	jsoniter "github.com/json-iterator/go"
)

// default active
const (
	active    ListingStatus = "active"
	inactive  ListingStatus = "inactive"
	untracked ListingStatus = "untracked"

	id      Sort = "id"
	cmcRank Sort = "cmc_rank"
)

type OpenAPIRsp struct {
	Status struct {
		Timestamp    time.Time `json:"timestamp"`
		ErrorCode    int       `json:"error_code"`
		ErrorMessage string    `json:"error_message"`
		Elapsed      int       `json:"elapsed"`
		CreditCount  int       `json:"credit_count"`
	} `json:"status"`
	Data jsoniter.RawMessage `json:"data"`
}

type Platform struct {
	Id           string `json:"id"`
	Name         string `json:"name"`
	Symbol       string `json:"symbol"`
	Slug         string `json:"slug"`
	TokenAddress string `json:"token_address"`
}
