package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gogo/protobuf/proto"
	context "golang.org/x/net/context"

	"github.com/lfxnxf/frame/BackendPlatform/golang/logging"
	"github.com/lfxnxf/frame/BackendPlatform/golang/sql"
	rpc "github.com/lfxnxf/frame/logic/rpc-go"
	"github.com/lfxnxf/frame/logic/rpc-go/demo/echo/echo"
)

//业务配置toml文件
type BussinessConfig struct {
	Uidlist []int `toml:"uidlist"`
}

type Response struct {
	DMErr    int         `json:"dm_error"`
	ErrorMsg string      `json:"error_msg"`
	Response interface{} `json:"response"`
}

// 数据库对应的model
type AllServerPolicy struct {
	ServerName      string `gorm:"server_name"`
	ServerType      int    `gorm:"server_type"`
	ServerURL       string `gorm:"server_url"`
	SecureServerURL string `gorm:"secure_server_url"`
}

type HelloHandler struct {
}
type TestHandler struct {
}

type TextHandler struct {
}

type SQLHandler struct {
}

type PanicHandler struct {
}

func (*PanicHandler) Serve(ctx context.Context, req *http.Request) (interface{}, int) {
	var pbytes []byte
	fmt.Printf("%q", pbytes[10])
	return nil, 0
	///return &Response{
	///	DMErr:    0,
	///	ErrorMsg: "操作成功",
	///	Response: nil,
	///}, 0
}

func (*HelloHandler) Serve(c context.Context, req *http.Request) (interface{}, int) {
	body, _ := ioutil.ReadAll(req.Body)
	fmt.Printf("request body %q\n", body)
	ctx, cancel := context.WithTimeout(c, 200*time.Millisecond)
	defer cancel()
	r, _ := http.NewRequest("GET", "http://127.0.0.1:6270"+"/test", nil)
	body, err := rpc.CallHTTP(ctx, r)
	r, _ = http.NewRequest("GET", "http://127.0.0.1:6270"+"/echo", nil)
	body, err = rpc.CallHTTP(ctx, r)
	logging.Debug("request %v response %q", err, body)
	return &Response{
		DMErr:    0,
		ErrorMsg: "操作成功",
		Response: body,
	}, 0
}

func (*TestHandler) Serve(ctx context.Context, req *http.Request) (interface{}, int) {
	data, _ := rpc.HttpGet(ctx, "long_link_proxy", "/echo", nil)
	fmt.Printf("simple response %q\n", data)
	return &Response{
		DMErr:    0,
		ErrorMsg: "操作成功",
		Response: "test",
	}, 0
}

func (*TextHandler) Text() bool {
	return false
}

func (*TextHandler) Serve(ctx context.Context, req *http.Request) (interface{}, int) {
	httpData, _ := rpc.HttpPost(ctx, "echo.EchoService", "/echo/EchoService/Echo", nil, bytes.NewReader([]byte(`{"message": "echo-message"}`)))
	client, _ := rpc.DialService("echo.EchoService", nil)
	cc := echo.NewEchoServiceClient(client)
	request := &echo.EchoRequest{
		Message: proto.String("echo message"),
	}
	data, _ := cc.Echo(ctx, request)
	return map[string]interface{}{"echo_http": string(httpData), "echo_rpc": data}, 0
}

func (*SQLHandler) Serve(context.Context, *http.Request) (interface{}, int) {
	var all []AllServerPolicy
	err := sql.Get("test1").Slave().Table("all_server_policy").Find(&all).Error
	if err != nil {
		return &Response{
			DMErr:    500,
			ErrorMsg: "内部系统错误",
		}, 500
	}
	return &Response{
		DMErr:    0,
		ErrorMsg: "操作成功",
		Response: all,
	}, 0
}

func main() {

	path := flag.String("config", "http.toml", "config file")
	flag.Parse()
	//path := "./http.toml"

	var bc BussinessConfig
	conf, _ := rpc.NewConfigToml(*path, &bc)

	fmt.Println(conf.HTTPServeLocation(), ",", conf.HTTPServeLogBody(), ",bussConfig:", bc)

	server := rpc.NewHTTPServerWithConfig(conf)
	server.Register(&HelloHandler{}, &TestHandler{}, &TextHandler{}, &SQLHandler{}, &PanicHandler{})

	server.Serve(conf.Port())
}
