package echo

import (
	"fmt"

	rpc "github.com/lfxnxf/frame/logic/rpc-go"
	context "golang.org/x/net/context"
)

type EchoService struct {
}

func (e *EchoService) Echo(c context.Context, echo *EchoRequest) (*EchoResponse, error) {
	resp := &EchoResponse{Message: echo.Message}
	//time.Sleep(5 * time.Millisecond)
	//fmt.Printf("echo payload %s\n", rpc.GetPayload(c))
	data, err := rpc.HttpGet(c, "test", "/api/fuck", nil)
	fmt.Printf("fuck response %q, err %v trace id %s\n", data, err, rpc.GetRequestID(c))
	return resp, nil
}
