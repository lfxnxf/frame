package proxy

import (
	"github.com/lfxnxf/frame/logic/inits"
	"github.com/lfxnxf/frame/logic/inits/rpc/client"
	"golang.org/x/net/context"
)

type RPC struct {
	name, endpoint string
}

func InitRPC(name, endpoint string) *RPC {
	return &RPC{name, endpoint}
}

func (r *RPC) Invoke(ctx context.Context, in interface{}, out interface{}, opts ...client.CallOption) error {
	return inits.RPCFactory(ctx, r.name).Client(r.endpoint).Invoke(ctx, in, out, opts...)
}
