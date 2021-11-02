package proxy

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/lfxnxf/frame/logic/inits/http/client"

	"github.com/lfxnxf/frame/logic/inits"
	httpserver "github.com/lfxnxf/frame/logic/inits/http/server"

	"github.com/lfxnxf/frame/logic/inits/internal/kit/tracing"
	"github.com/opentracing/opentracing-go"
	"github.com/stretchr/testify/assert"
)

// 启动一个服务
var testServerData = `
[server]
service_name="inf.server.base"
port=20206
[server.breaker]
"/server/phone/decode"={break=false, error_percent_threshold=50, consecutive_error_threshold=5, minsamples=10}
[server.limiter]
"unknown@/server/phone/encode"={open=true, limits=10}
[log]
level="debug"
logpath="logs"
rotate="hour"
`

var testClientData = `
[server]
service_name="inf.client.base"
port=20207
[log]
level="debug"
logpath="logs"
rotate="hour"

[[server_client]]
service_name="inf.client.base"
proto="http"
endpoints="127.0.0.1:20207"
balancetype="roundrobin"
read_timeout=500
slow_time=500
namespace="ClientLimiterAndBreaker"
[server_client.resource]
"/phone/encode"={open=true, limits=20, break=false, error_percent_threshold=50, consecutive_error_threshold=5, minsamples=10}
"/phone/decode"={open=true, limits=20, break=false, error_percent_threshold=50, consecutive_error_threshold=5, minsamples=10}
`

var testNamespaceKeyData = `
[server]
service_name="inf.namespace.base"
port=20209
[log]
level="debug"
logpath="logs"
rotate="hour"

[[server_client]]
service_name="inf.namespace.base"
proto="http"
endpoints="127.0.0.1:20209"
balancetype="roundrobin"
read_timeout=500
slow_time=500
namespace="unknown_namespace_key"
[server_client.resource]
"/phone/encode"={open=true, limits=20, break=false, error_percent_threshold=50, consecutive_error_threshold=5, minsamples=10}
"/phone/decode"={open=true, limits=20, break=false, error_percent_threshold=50, consecutive_error_threshold=5, minsamples=10}

[[server_client]]
service_name="inf.namespace.base2"
proto="http"
endpoints="127.0.0.1:20209"
balancetype="roundrobin"
read_timeout=500
slow_time=500
namespace="namespace_key"
[server_client.resource]
"/phone/encode"={open=true, limits=20, break=false, error_percent_threshold=50, consecutive_error_threshold=5, minsamples=10}
"/phone/decode"={open=true, limits=20, break=false, error_percent_threshold=50, consecutive_error_threshold=5, minsamples=10}
`

func TestServerLimiterAndBreaker(t *testing.T) {
	_, dae := testInitInits("ServerLimiterAndBreaker", "./testconfig2/", "./testconfig2/config.toml", testServerData)

	srv := dae.HTTPServer()
	srv.ANY("/server/phone/encode", func(c *httpserver.Context) {
		c.JSON(c.Request.URL.String(), nil)
	})
	srv.ANY("/server/phone/decode", func(c *httpserver.Context) {
		num, _ := strconv.Atoi(c.Request.URL.Query().Get("num"))
		if num > 5 && num <= 15 {
			c.AbortErr(fmt.Errorf("bad resquest"))
		}
		c.JSON(c.Request.URL.String(), nil)
	})

	go func() {
		if err := srv.Run(); err != nil {
			fmt.Println(err)
		}
	}()

	fmt.Println("====================== ServerLimiterAndBreaker start http server ======================")
	time.Sleep(1 * time.Second)

	cc := http.Client{Timeout: 10 * time.Second}

	for i := 0; i < 20; i++ {
		req, _ := http.NewRequest("GET", fmt.Sprintf("http://127.0.0.1:20206/server/phone/encode?num=%d", i), nil)
		tracer := opentracing.GlobalTracer()
		ctx := opentracing.ContextWithSpan(context.Background(), tracer.StartSpan("HTTP Do"))
		span := opentracing.SpanFromContext(ctx)
		span.SetBaggageItem("_namespace_appkey_", "ServerLimiterAndBreaker")
		nReq := tracing.ContextToHTTP(ctx, tracer, req)
		resp, err := cc.Do(nReq)
		if err != nil {
			fmt.Printf("===>>> ServerLimiterAndBreaker error:%s\n", err.Error())
			t.Fail()
		}
		if resp != nil {
			bodyB, err := ioutil.ReadAll(resp.Body)
			assert.Equal(t, err, nil)
			fmt.Printf("===>>> ServerLimiterAndBreaker resp:%s\n", string(bodyB))
			_ = resp.Body.Close()
			if i >= 10 {
				assert.Equal(t, "rate limit exceeded", string(bodyB))
			}
		}
	}

	for i := 0; i < 20; i++ {
		req, _ := http.NewRequest("GET", fmt.Sprintf("http://127.0.0.1:20206/server/phone/decode?num=%d", i), nil)
		tracer := opentracing.GlobalTracer()
		ctx := opentracing.ContextWithSpan(context.Background(), tracer.StartSpan("HTTP Do"))
		span := opentracing.SpanFromContext(ctx)
		span.SetBaggageItem("_namespace_appkey_", "ServerLimiterAndBreaker")
		nReq := tracing.ContextToHTTP(ctx, tracer, req)
		resp, err := cc.Do(nReq)
		if err == nil {
			fmt.Printf("===>>> ServerLimiterAndBreaker http status:%s, num:%d\n", resp.Status, i)
			bodyB, err := ioutil.ReadAll(resp.Body)
			_ = resp.Body.Close()
			assert.Equal(t, err, nil)
			bodyS := string(bodyB)
			fmt.Printf("===>>> ServerLimiterAndBreaker resp:%s\n", bodyS)
			if i >= 11 {
				assert.Equal(t, "breaker: consecutive error threshold", bodyS)
			}
		} else {
			fmt.Printf("===>>> ServerLimiterAndBreaker error:%s\n", err.Error())
		}
	}
}

func TestClientLimiterAndBreaker(t *testing.T) {
	ns, dae := testInitInits("ClientLimiterAndBreaker", "./testconfig3/", "./testconfig3/config.toml", testClientData)
	ctx := inits.WithAPPKey(context.Background(), ns)

	srv := dae.HTTPServer()
	srv.ANY("/phone/encode", func(c *httpserver.Context) {
		c.JSON(c.Request.URL.String(), nil)
	})

	srv.ANY("/phone/decode", func(c *httpserver.Context) {
		num, _ := strconv.Atoi(c.Request.URL.Query().Get("num"))
		if num < 10 {
			time.Sleep(2 * time.Second)
		} else {
			c.JSON(c.Request.URL.String(), nil)
		}
	})

	go func() {
		if err := srv.Run(); err != nil {
			fmt.Println(err)
		}
	}()

	fmt.Println("====================== ClientLimiterAndBreaker start http server ======================")
	time.Sleep(1 * time.Second)

	cc := dae.HTTPClient("inf.client.base")
	// limiter
	for i := 0; i < 30; i++ {
		req := client.NewRequest(ctx).
			WithMethod("GET").
			WithPath("/phone/encode").
			WithQueryParam("num", fmt.Sprintf("%d", i))
		resp, err := cc.Call(req)
		if err == nil {
			fmt.Printf("===>>> ClientLimiterAndBreaker num:%d, resp:%s\n", i, resp.String())
		} else {
			fmt.Printf("===>>> ClientLimiterAndBreaker num:%d, error:%s\n", i, err.Error())
			assert.Equal(t, "rate limit exceeded", err.Error())
		}
	}

	fmt.Println("====================== ClientLimiterAndBreaker sleep ======================")
	time.Sleep(1 * time.Second)

	// breaker
	for i := 0; i < 30; i++ {
		req := client.NewRequest(ctx).
			WithMethod("GET").
			WithPath("/phone/decode").
			WithQueryParam("num", fmt.Sprintf("%d", i))
		resp, err := cc.Call(req)
		if err == nil {
			fmt.Printf("===>>> ClientLimiterAndBreaker num:%d, resp:%s\n", i, resp.String())
		} else {
			fmt.Printf("===>>> ClientLimiterAndBreaker num:%d, error:%s\n", i, err.Error())
			if i < 5 {
				assert.Equal(t, "context deadline exceeded", err.Error())
			}
			if i >= 5 && i <= 24 {
				assert.Equal(t, "breaker: consecutive error threshold", err.Error())
			}
			if i > 24 {
				assert.Equal(t, "rate limit exceeded", err.Error())
			}
		}
	}
}

func TestUnknownNamespaceKey(t *testing.T) {
	ns, dae := testInitInits("namespace_key", "./testconfig4/", "./testconfig4/config.toml", testNamespaceKeyData)
	ctx := inits.WithAPPKey(context.Background(), ns)

	srv := dae.HTTPServer()
	srv.ANY("/namespace/key", func(c *httpserver.Context) {
		c.JSON(c.Request.URL.String(), nil)
	})

	go func() {
		if err := srv.Run(); err != nil {
			fmt.Println(err)
		}
	}()

	fmt.Println("====================== TestUnknownNamespaceKey start http server ======================")
	time.Sleep(1 * time.Second)

	cc := dae.HTTPClient("inf.namespace.base")

	req := client.NewRequest(ctx).
		WithMethod("GET").
		WithPath("/namespace/key")
	resp, err := cc.Call(req)
	if err == nil {
		fmt.Printf("===>>> TestUnknownNamespaceKey service_name:inf.namespace.base, uri:/namespace, resp:%s, code:%d\n", resp.String(), resp.Code())
		assert.Equal(t, 501, resp.Code())
		assert.Equal(t, "unknown namespace", resp.String())
	} else {
		fmt.Printf("===>>> TestUnknownNamespaceKey service_name:inf.namespace.base, uri:/namespace, error:%s\n", err.Error())
	}

	cc2 := dae.HTTPClient("inf.namespace.base2")

	req2 := client.NewRequest(ctx).
		WithMethod("GET").
		WithPath("/namespace/key2")
	resp2, err := cc2.Call(req2)
	if err == nil {
		fmt.Printf("===>>> TestUnknownNamespaceKey service_name:inf.namespace.base2, uri:/namespace/key2, resp:%s, code:%d\n", resp2.String(), resp2.Code())
		assert.Equal(t, 404, resp2.Code())
		assert.Equal(t, "Not Found", resp2.String())
	} else {
		fmt.Printf("===>>> TestUnknownNamespaceKey service_name:inf.namespace.base2, uri:/namespace/key2, error:%s\n", err.Error())
	}

	req3 := client.NewRequest(ctx).
		WithMethod("GET").
		WithPath("/namespace/key")

	resp3, err := cc2.Call(req3)
	if err == nil {
		fmt.Printf("===>>> TestUnknownNamespaceKey service_name:inf.namespace.base2, uri:/namespace/key, resp:%s, code:%d\n", resp3.String(), resp3.Code())
		assert.Equal(t, 200, resp3.Code())
		assert.Equal(t, `{"dm_error":0,"error_msg":"0","data":"/namespace/key"}`, strings.Trim(resp3.String(), "\n"))
	} else {
		fmt.Printf("===>>> TestUnknownNamespaceKey service_name:inf.namespace.base2, uri:/namespace/key, error:%s\n", err.Error())
	}
}
