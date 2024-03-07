package client

import "fmt"

type uri string

const (

	// cryptocurrency API
	getMapUri uri = "/v1/cryptocurrency/map"
)

// getApi build the full api url
func (c *client) getApi(u uri) string {
	return fmt.Sprintf("%s%s", c.conf.BaseApi, u)
}
