package main

import (
	"flag"
	"fmt"
	rpc "github.com/lfxnxf/frame/logic/rpc-go"
	"github.com/gogo/protobuf/proto"
	context "golang.org/x/net/context"
	"time"

	"github.com/lfxnxf/frame/logic/rpc-go/demo/echo/echo"
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

	serviceName := "servicerpc"
	requestOptional := rpc.NewRequestOptional()
	requestOptional.SetTimeOut(1000)
	requestOptional.SetRetryTimes(1)
	client, errf := rpc.DialService(serviceName, requestOptional)
	if errf != nil {
		fmt.Println("rpc dialService error ,err:", errf)
		return
	}

	c := echo.NewEchoServiceClient(client)
	for i := 0; i < 1000; i++ {

		r := echo.EchoRequest{Message: proto.String("echo")}
		rsp, er := c.Echo(context.Background(), &r)
		if er != nil {
			fmt.Println("call echo error ,err:", er, rsp.GetMessage())
			continue
		}
		time.Sleep(1 * time.Second)
	}
}
