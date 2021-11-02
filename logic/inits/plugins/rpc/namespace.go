package rpcplugin

import (
	"github.com/lfxnxf/frame/logic/inits"
	rpc "github.com/lfxnxf/frame/logic/inits/rpc/server"
)

func Namespace(c *rpc.Context) {
	if c.Namespace != "" {
		c.Ctx = inits.WithAPPKey(c.Ctx, c.Namespace)
	}
	c.Next()
}
