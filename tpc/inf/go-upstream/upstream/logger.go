package upstream

import (
	log "github.com/lfxnxf/frame/BackendPlatform/golang/logging"
)

var (
	logging *log.Logger
)

func init() {
	logging = log.New()
}

func SetLogger(l *log.Logger) {
	logging = l
}
