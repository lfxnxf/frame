package main

import (
	"flag"
	"fmt"
	// "math/rand"
	"os"
	"os/signal"

	"github.com/lfxnxf/frame/logic/rpc-go/demo/echo/echo"

	rpc "github.com/lfxnxf/frame/logic/rpc-go"

	log "github.com/lfxnxf/frame/BackendPlatform/golang/logging"
)

type Config struct {
	Server struct {
		Name string `toml:"name"`
		Port int    `toml:"port"`
		TCP  struct {
			IdleTimeout       int `toml:"idle_timeout"`
			KeepalinkInterval int `toml:"keepalink_interval"`
		} `toml:"tcp"`
	} `toml:"server"`
	Trace struct {
		Port int `toml:"port"`
	} `toml:"trace"`
	Log struct {
		Level       string `toml:"level"`
		Businesslog string `toml:"businesslog"`
		Serverlog   string `toml:"serverlog"`
		Accesslog   string `toml:"accesslog"`
		StatLog     string `toml:"statLog"`
		StatMetric  string `toml:"statMetric"`
	} `toml:"log"`
	Redis struct {
		Status struct {
			IP   string `toml:"ip"`
			Port int    `toml:"port"`
		} `toml:"status"`
	} `toml:"redis"`
}

func main() {

	conf := flag.String("config", "./config.toml", "rpc config file")
	flag.Parse()

	var cc Config
	config, err := rpc.NewConfigToml(*conf, &cc)
	if err != nil {
		fmt.Printf("New Config error %s\n", err)
	}

	// fmt.Println(config, ",,,")

	// rpc.InitLoadBalance(config.GetServiceClients())

	server := rpc.NewServerWithConfig(config)

	// Register service
	srv := &echo.EchoService{}
	echo.RegisterEchoServiceServer(server, srv)

	//Client Example
	//port := config.Port()

	//go func() {
	//        time.Sleep(1 * time.Second)
	//        client, _ := rpc.Dial(fmt.Sprintf("127.0.0.1:%d", port+1))
	//        // client, _ := rpc.Dial("echo.EchoService", rpc.WithConfigBalance(true))
	//        c := echo.NewEchoServiceClient(client)
	//        for {
	//                for i := 0; i < 1; i++ {
	//                        go func() {
	//                                r := echo.EchoRequest{Message: proto.String("echo")}
	//                                // ctx, cancel := context.WithTimeout(context.Background(), time.Duration(rand.Intn(10000000))*time.Nanosecond)
	//                                // defer cancel()
	//                                //resp, err := c.Echo(ctx, &r)
	//                                resp,err:=c.Echo(context.TODO(), &r)
	//                                log.Debugf("echo resp `%v`, error %v", resp, err)
	//                                time.Sleep(1 * time.Second)
	//                        }()
	//                }
	//                time.Sleep(1 * time.Second)
	//        }
	//}()

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
