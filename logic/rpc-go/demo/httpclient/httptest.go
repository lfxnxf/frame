package main

import (
	"context"
	"flag"
	"fmt"
	rpc "github.com/lfxnxf/frame/logic/rpc-go"
)

type Config struct {
}

func main() {

	conf := flag.String("config", "./config.toml", "rpc config file")
	flag.Parse()

	// 读取配置文件
	var cc Config
	_, err := rpc.NewConfigToml(*conf, &cc)
	if err != nil {
		fmt.Printf("New Config error %s\n", err)
	}

	var sum = 1
	for sum < 100 {
		// 获取httpclient,注意url的使用

		ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
		req, ef := rpc.NewRequest("echo.EchoService", "GET", "http://service.inke.cn/api/live/simpleall?uid=106839", nil)
		if ef != nil {
			fmt.Println("eff:", ef)
			return
		}
		data, err := rpc.CallHTTP(context.TODO(), req)
		cancel()
		if err == nil {
			fmt.Printf("response %q, %v\n", data, err)
		}
	}
}

// req, ef := rpc.NewRequest("echo.EchoService", "GET", "http://service.inke.cn/api/live/simpleall?uid=106839", nil)

// req, ef := rpc.NewRequest("echo.EchoService", "GET", "http://{ipport}/api/live/simpleall?uid=106839", nil)
