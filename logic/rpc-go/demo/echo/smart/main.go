package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/lfxnxf/frame/logic/rpc-go/demo/echo/echo"

	rpc "github.com/lfxnxf/frame/logic/rpc-go"
	util "github.com/lfxnxf/frame/logic/rpc-go/util"

	log "github.com/lfxnxf/frame/BackendPlatform/golang/logging"
)

func main() {
	port := flag.Int("port", 4041, "port ...")
	flag.Parse()
	l := log.GetLogger()
	server := rpc.NewServer(
		util.Timeout(120*time.Second),
		util.Keepalive(10*time.Second),
		util.Logger(l),
		//util.RegisterAddr([]string{"127.0.0.1:8500"}))
		util.RegisterAddr([]string{"10.26.47.163:8500"}))

	// Register service
	srv := &echo.EchoService{}
	echo.RegisterEchoServiceServer(server, srv)

	//Signal to gracefull shutdown
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt, syscall.SIGUSR2, syscall.SIGTERM)
	go func() {
		for {
			sig := <-quit
			fmt.Printf("accept signal %d\n", sig)
			if sig == syscall.SIGUSR2 {
				log.Infof("Unregister....")
				server.Unregister()
				continue
			}
			log.Infof("Shutdown....")
			server.Stop()
			break
		}
	}()

	//Start Serving
	err := server.Serve(*port)
	if err != nil {
		fmt.Printf("server error %v", err)
	}
}
