package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/lfxnxf/frame/BackendPlatform/inits-tool/inits/rpcdemo/conf"
	"github.com/lfxnxf/frame/BackendPlatform/inits-tool/inits/rpcdemo/server/rpc"
	"github.com/lfxnxf/frame/BackendPlatform/inits-tool/inits/rpcdemo/service"
	"github.com/lfxnxf/frame/BackendPlatform/golang/logging"
	"github.com/lfxnxf/frame/logic/inits"
)

func init() {
	configS := flag.String("config", "config/config.toml", "Configuration file")
	appS := flag.String("app", "", "App dir")
	flag.Parse()
	
	inits.Init(
		inits.ConfigPath(*configS),
	)
	
	if *appS != "" {
		inits.InitNamespace(*appS)
	}

}

func main() {
	log.Println("rpcdemo start")

	defer inits.Shutdown()

	// init local config
	cfg, err := conf.Init()
	if err != nil {
		logging.Fatalf("service config init error %s", err)
	}

	// create a service instance
	svc := service.New(cfg)

	// init and start http server
	rpc.Init(svc, cfg)

	defer rpc.Shutdown()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-sigChan
		log.Printf("get a signal %s\n", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			log.Println("rpcdemo server exit now...")
			return
		case syscall.SIGHUP:
		default:
		}
	}

}