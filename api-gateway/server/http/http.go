package http

import (
	"github.com/lfxnxf/frame/api-gateway/conf"
	"github.com/lfxnxf/frame/api-gateway/service"
	"github.com/lfxnxf/frame/BackendPlatform/golang/logging"
	"github.com/lfxnxf/frame/logic/inits"
	httpserver "github.com/lfxnxf/frame/logic/inits/http/server"
	httpplugin "github.com/lfxnxf/frame/logic/inits/plugins/http"
)

var (
	svc *service.Service

	httpServer httpserver.Server
)

// Init create a rpc server and run it
func Init(s *service.Service, conf *conf.Config) {
	svc = s

	// new http server
	httpServer = inits.HTTPServer()

	// add namespace plugin
	httpServer.Use(httpplugin.Namespace)

	// register handler with http route
	initRoute(httpServer)

	// start a http server
	go func() {
		if err := httpServer.Run(); err != nil {
			logging.Fatalf("http server start failed, err %v", err)
		}
	}()

}

func Shutdown() {
	if httpServer != nil {
		httpServer.Stop()
	}
	if svc != nil {
		svc.Close()
	}
}
