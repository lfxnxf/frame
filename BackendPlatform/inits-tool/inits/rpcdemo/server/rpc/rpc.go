package rpc

import (
	api "github.com/lfxnxf/frame/BackendPlatform/inits-tool/inits/rpcdemo/api/game_center_base"
	"github.com/lfxnxf/frame/BackendPlatform/inits-tool/inits/rpcdemo/conf"
	"github.com/lfxnxf/frame/BackendPlatform/inits-tool/inits/rpcdemo/service"
	"github.com/lfxnxf/frame/BackendPlatform/golang/logging"
	"github.com/lfxnxf/frame/logic/inits"
	rpcserver "github.com/lfxnxf/frame/logic/inits/rpc/server"
	rpcplugin "github.com/lfxnxf/frame/logic/inits/plugins/rpc"
)


var (
	svc *service.Service

	rpcServer rpcserver.Server
)

// Init create a rpc server and run it
func Init(s *service.Service, conf *conf.Config) {
	svc = s

	rpcServer = inits.RPCServer()

	// add namespace plugin
	rpcServer.Use(rpcplugin.Namespace)

	api.RegisterGameCenterServiceHandler(rpcServer, svc)

	// start a rpc server
	if err := rpcServer.Start(); err != nil {
		logging.Fatalf("rpc server start failed, err %v", err)
	}
}

// Close close the resource
func Shutdown() {
	if rpcServer != nil {
		rpcServer.Stop()
	}
	if svc != nil {
		svc.Close()
	}
}

