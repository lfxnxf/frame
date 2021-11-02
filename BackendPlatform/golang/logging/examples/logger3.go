package main

import (
	"github.com/lfxnxf/frame/BackendPlatform/golang/logging"
)

var (
	log *logging.Logger
)

func init() {
	log = logging.NewLogger(&logging.Options{
		DisableTimestamp: true,
		DisableLevel:     true,
	})
}

func main() {
	log.Infof("This message will print into stdout")
}
