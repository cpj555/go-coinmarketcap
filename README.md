<p align="center">
  <a href="https://coinmarketcap.com/api/documentation/v1/">
    <img src="examples/coinmarketcap.png" width="200" height="200" alt="dodo-open">
  </a>
</p>

<div align="center">


# go-coinmarketcap
  <a href="https://github.com/dodo-open/dodo-open-go/blob/main/LICENSE">
    <img src="https://img.shields.io/github/license/cpj555/go-coinmarketcap" alt="license">
  </a>

</div>



## 特性

- coinmarketcap OpenAPI
- 看github现有的sdk太久没更新,自己封装了一套(2024/03/07)
- 只封装了自己用到的api


## 起步

```shell
go get -u github.com/cpj555/go-coinmarketcap
```



```go
apiKey := "your-api-key"

// 下面的第三个参数，设定了 resty 的请求超时为 3 秒：
//
//     client.WithTimeout(time.Second*3)
//
client, err := NewClient(client.WithApiKey(apiKey), client.WithTimeout(time.Second*3),
client.WithProxyUrl("http://localhost:1080") //本地调试得开代理设置下代理请求,生产环境部署到香港或者国外服务器
	)

// 获取map列表，可以使用下面的方法
list, err := client.CryptocurrencyV1.GetMap(context.Background(),&types.GetCryptocurrencyV1MapReq{Start: 1, Limit: 10})
```

### API documentation
https://coinmarketcap.com/api/documentation

- [examples](examples/demo-setup)



* **/v1/cryptocurrency/map** (Get the current price of any cryptocurrencies in any other supported currencies that you need)
  ```go
  client.CryptocurrencyV1.GetMap(context.Background(),&types.GetCryptocurrencyMapReq{Start: 1, Limit: 10})
  ```