package main

import (
	"flag"
	rpc "github.com/lfxnxf/frame/logic/rpc-go"
	"log"
)

type Config struct {
}

func main() {
	conf := flag.String("config", "./config.toml", "rpc config file")
	flag.Parse()

	var cc Config
	_, err := rpc.NewConfigToml(*conf, &cc)
	if err != nil {
		log.Fatalf("New Config error %s\n", err)
	}

	r1, _ := rpc.GetRedis("redis1")
	if r1 == nil {
		log.Fatalf("redis is nil\n")
	}

	r2, _ := rpc.GetRedis("redis22")
	if r2 == nil {
		log.Fatalf("redis is nil\n")
	}
	r1.Set("hahaha", 123)
	r2.Set("hahaha", 123)
	r1.Set("hahaha1", 1234)
	r2.Set("hahaha1", 1234)
}
