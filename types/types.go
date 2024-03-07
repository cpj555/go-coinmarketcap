package types

import (
	jsoniter "github.com/json-iterator/go"
	"time"
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
