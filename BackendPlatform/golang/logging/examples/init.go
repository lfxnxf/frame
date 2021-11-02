package main

import (
	"github.com/lfxnxf/frame/BackendPlatform/golang/logging"
)

var (
	log *logging.Logger
)

func init() {
	log = logging.NewLogger(&logging.Options{})
}

func main() {
	log.Debugf("This is debug")
}
