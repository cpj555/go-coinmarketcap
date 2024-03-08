package main

import (
	"context"
	"fmt"
	"github.com/cpj555/go-coinmarketcap/client"
	"github.com/cpj555/go-coinmarketcap/types"
	"time"
)

func main() {

	c, err := client.New(client.WithSandBox(true), client.WithTimeout(10*time.Second), client.WithProxyUrl("http://localhost:1080"))
	if err != nil {
		fmt.Printf("创建实例失败：%v\n", err)
		return
	}

	//getMap, err := c.GetMap(context.Background(), &types.GetMapReq{Start: 1, Limit: 1})
	//if err != nil {
	//	fmt.Printf("getMap err：%v\n", err)
	//	return
	//}
	//fmt.Printf("%+v\n", getMap)

	getInfo, err := c.GetInfo(context.Background(), &types.GetInfoReq{Id: "1"})
	if err != nil {
		fmt.Printf("getMap err：%v\n", err)
		return
	}
	fmt.Printf("%+v\n", getInfo)

}
