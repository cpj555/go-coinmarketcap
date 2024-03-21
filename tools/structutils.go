package tools

import (
	"net/url"

	"github.com/gorilla/schema"
)

func ToUrlValues(r interface{}) url.Values {
	values := url.Values{}
	encoder := schema.NewEncoder()
	encoder.Encode(r, values)
	return values
}
