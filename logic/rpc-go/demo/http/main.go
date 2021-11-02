package main

import (
	"fmt"
	"net/http"
	"time"

	context "golang.org/x/net/context"

	rpc "github.com/lfxnxf/frame/logic/rpc-go"
)

type Response struct {
	DMErr    int    `json:"dm_error"`
	ErrorMsg string `json:"error_msg"`
	Response string `json:"response"`
}

type HelloHandler struct {
}
type TestHandler struct {
}

type TextHandler struct {
}

func (*HelloHandler) Serve(context.Context, *http.Request) (interface{}, int) {
	return Response{
		DMErr:    0,
		ErrorMsg: "操作成功",
		Response: "HelloWorld",
	}, 0
}

func (*TestHandler) Serve(context.Context, *http.Request) (interface{}, int) {
	return &Response{
		DMErr:    0,
		ErrorMsg: "操作成功",
		Response: "test",
	}, 0
}

func (*TextHandler) Text() bool {
	return true
}

func (*TextHandler) Serve(context.Context, *http.Request) (interface{}, int) {
	time.Sleep(90 * time.Millisecond)
	return "HelloWorld10aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa0", 0
}

func main() {
	conf, _ := rpc.NewConfig("http.ini")
	server := rpc.NewHTTPServerWithConfig(conf)
	server.Register(&HelloHandler{}, &TestHandler{}, &TextHandler{})
	//go func() {
	//        for {
	//                ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	//                r, _ := http.NewRequest("GET", "http://127.0.0.1:"+strconv.Itoa(conf.Port())+"/test", nil)
	//                rpc.CallHTTP(ctx, r)
	//                cancel()
	//                //log.Debugf("response %q %v", rsp, err)
	//        }
	//}()
	//go func() {
	//        for {
	//                ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	//                r, _ := http.NewRequest("GET", "http://127.0.0.1:"+strconv.Itoa(conf.Port())+"/test", nil)
	//                rpc.CallHTTP(ctx, r)
	//                cancel()
	//                //log.Debugf("response %q %v", rsp, err)
	//        }
	//}()
	//go func() {
	//        for {
	//                ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	//                r, _ := http.NewRequest("GET", "http://127.0.0.1:"+strconv.Itoa(conf.Port())+"/test", nil)
	//                rpc.CallHTTP(ctx, r)
	//                cancel()
	//                //log.Debugf("response %q %v", rsp, err)
	//        }
	//}()
	server.Serve(conf.Port())
	fmt.Println("vim-go")
}
