package _testing

import (
	newpb "github.com/lfxnxf/frame/logic/inits/rpc/_testing/newpb"
	"github.com/lfxnxf/frame/logic/inits/rpc/server"
	"github.com/golang/protobuf/proto"
	"golang.org/x/net/context"
)

type NServer struct {
	s server.Server
}

func (s *NServer) Echo(ctx context.Context, r *newpb.EchoRequest) (*newpb.EchoResponse, error) {
	ra := &newpb.EchoResponse{
		Response: proto.String(r.GetMessage()),
		Code:     newpb.ResponseCode_SUCCESS.Enum(),
	}
	return ra, nil
}

func NNew(server server.Server) *NServer {
	s := &NServer{server}
	if err := newpb.RegisterEchoServiceHandler(server, s); err != nil {
		panic(err)
	}
	return s
}

func (s *NServer) Start() error {
	return s.s.Start()
}
