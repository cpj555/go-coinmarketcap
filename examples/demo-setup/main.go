package main

import (
	"context"
	"fmt"
	"github.com/cpj555/go-coinmarketcap/client"
	"github.com/cpj555/go-coinmarketcap/types"
)

func main() {

	c, err := client.New(client.WithSandBox(true))
	if err != nil {
		fmt.Printf("创建实例失败：%v\n", err)
		return
	}

	getMap, err := c.GetMap(context.Background(), &types.GetMapReq{Start: 1, Limit: 10})
	if err != nil {
		fmt.Printf("getMap err：%v\n", err)
		return
	}

	fmt.Printf("%+v\n\n", getMap)
}
