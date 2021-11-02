package main

import (
	"context"
	"time"

	log "github.com/lfxnxf/frame/BackendPlatform/golang/logging"
	"github.com/gogo/protobuf/proto"

	rpc "github.com/lfxnxf/frame/logic/rpc-go"
	"github.com/lfxnxf/frame/logic/rpc-go/demo/echo/echo"
	"github.com/lfxnxf/frame/logic/rpc-go/discovery/consul"
)

func main() {
	log.SetOutputByName("./client.log")
	rpc.SetBusinessLogFile("./business.log")
	resolver := consul.NewConsulResolver(
		[]string{"127.0.0.1:8500"}, "")
	//resolver := file.NewResolver("./address", "")
	balancer := rpc.RoundRobin(resolver)

	client, err := rpc.Dial("echo.EchoService", rpc.WithBalancer(balancer))
	if err != nil {
		log.Fatalf("Dial client error %+v", err)
	}
	time.Sleep(1 * time.Second)
	c := echo.NewEchoServiceClient(client)

	for {
		for i := 0; i < 100; i++ {
			r := echo.EchoRequest{Message: proto.String("echo")}
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
			//resp, err := c.Echo(ctx, &r)
			c.Echo(ctx, &r)
			cancel()
			//log.Debugf("echo resp `%v`, error %v", resp, err)
			//fmt.Println(client, err)
		}
		time.Sleep(1 * time.Second)
	}
}
