package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"time"

	context "golang.org/x/net/context"

	"github.com/gogo/protobuf/proto"

	"github.com/lfxnxf/frame/logic/rpc-go/demo/echo/echo"

	rpc "github.com/lfxnxf/frame/logic/rpc-go"

	log "github.com/lfxnxf/frame/BackendPlatform/golang/logging"
)

func main() {
	conf := flag.String("config", "./rpc.ini", "rpc config file")
	flag.Parse()

	config, err := rpc.NewConfig(*conf)
	if err != nil {
		fmt.Printf("New Config error %s\n", err)
	}
	server := rpc.NewServerWithConfig(config)

	// Register service
	srv := &echo.EchoService{}
	echo.RegisterEchoServiceServer(server, srv)

	//Client Example
	port := config.Port()
	go func() {
		time.Sleep(1 * time.Second)
		client, _ := rpc.Dial(fmt.Sprintf("127.0.0.1:%d", port+1))
		c := echo.NewEchoServiceClient(client)
		for {
			for i := 0; i < 1; i++ {
				go func() {
					r := echo.EchoRequest{Message: proto.String("echo")}
					ctx, cancel := context.WithTimeout(context.Background(), time.Duration(rand.Intn(10000000))*time.Nanosecond)
					defer cancel()
					//resp, err := c.Echo(ctx, &r)
					resp, err := c.Echo(ctx, &r)
					log.Debugf("echo resp %v, error %v", resp.GetMessage(), err)
					time.Sleep(1 * time.Second)
				}()
			}
			time.Sleep(1 * time.Second)
		}
	}()

	//Signal to gracefull shutdown
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	go func() {
		<-quit
		log.Info("Shutdown....")
		server.Stop()
	}()

	//Start Serving
	err = server.Serve(config.Port())
	if err != nil {
		fmt.Printf("server error %v", err)
	}
}
