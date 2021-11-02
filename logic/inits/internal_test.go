package inits

import (
	"context"
	"fmt"
	"path/filepath"
	"strings"
	"testing"
	"time"

	jaeger "github.com/lfxnxf/frame/BackendPlatform/jaeger-client-go"
	jaegerconfig "github.com/lfxnxf/frame/BackendPlatform/jaeger-client-go/config"
	"github.com/lfxnxf/frame/logic/inits/config/encoder/json"
	httpserver "github.com/lfxnxf/frame/logic/inits/http/server"
	"github.com/lfxnxf/frame/logic/inits/internal/kit/sd"
	"github.com/lfxnxf/frame/logic/inits/internal/kit/tracing"
	"github.com/BurntSushi/toml"
	"github.com/stretchr/testify/assert"
)

var testdata = `
[server]
service_name="inf.sms.base"
port=10000
[server.breaker]
"/api/sms/send"={break=false, error_percent_threshold=50, connsecutive_error_threshold=50, minsamples=100} # 特定接口
"*"={break=false, error_percent_threshold=50, connsecutive_error_threshold=50, minsamples=100} # 所有接口
[server.limiter]
"*@/api/sms/send"={open=true, limits=10000} # 上游所有服务
"nvwa.usercenter.account.logic./api/sms/send"={open=true, limits=10000} # 上游特定服务

[log]
level="debug"
logpath="logs"
rotate="hour"

[[server_client]]
service_name="inf.secret.base"
proto="http"
endpoints_from="consul"
balancetype="roundrobin"
read_timeout=10000
retry_times=3
slow_time=1000
[server_client.resource]
"/api/sms/send"={open=true, limits=10000, break=false, error_percent_threshold=50, connsecutive_error_threshold=50, minsamples=100}
"*"={open=true, limits=10000, break=false, error_percent_threshold=50, connsecutive_error_threshold=50, minsamples=100}
`

func testInitInitsConfig(data string) initsConfig {
	c := initsConfig{}
	if err := toml.Unmarshal([]byte(data), &c); err != nil {
		panic(err)
	}
	return c
}

func Test_getServerLimiterConfig(t *testing.T) {
	c := testInitInitsConfig(testdata)
	serverM := getServerLimiterConfig("", c)
	for _, v := range serverM {
		fmt.Printf("server limiter key:%s, val:%+v\n", v.Name, v)
	}
}

func Test_getServerBreakerConfig(t *testing.T) {
	c := testInitInitsConfig(testdata)
	serverM := getServerBreakerConfig("", c)
	for _, v := range serverM {
		fmt.Printf("server breaker key:%s, val:%+v\n", v.Name, v)
	}
}

func Test_getLimiterConfig(t *testing.T) {
	c := testInitInitsConfig(testdata)
	for _, v := range c.ServerClient {
		clientM := getClientLimiterConfig("", v)
		for _, v := range clientM {
			fmt.Printf("client limiter key:%s, val:%+v\n", v.Name, v)
		}
	}
}

func Test_getBreakerConfig(t *testing.T) {
	c := testInitInitsConfig(testdata)
	for _, v := range c.ServerClient {
		clientM := getClientBreakerConfig("", v)
		for _, v := range clientM {
			fmt.Printf("client breaker key:%s, val:%+v\n", v.Name, v)
		}
	}
}

func Test_tracingKVPut(t *testing.T) {

	handlers := map[string]httpserver.HandlerFunc{
		"/api/kv/put": func(c *httpserver.Context) {

			c.Raw(map[string]interface{}{
				"dm_error":  0,
				"error_msg": "success",
			}, 0)
			return
		},
	}

	serverPort := 5778

	// init tracer
	cfg := jaegerconfig.Configuration{
		// SamplingServerURL: "http://localhost:5778/sampling"
		Sampler: &jaegerconfig.SamplerConfig{Type: jaeger.SamplerTypeRemote},
		Reporter: &jaegerconfig.ReporterConfig{
			LogSpans:            false,
			BufferFlushInterval: 1 * time.Second,
			LocalAgentHostPort:  "127.0.0.1:6831",
		},
	}
	tracer, _, err := cfg.New("danerys.test.service")
	if err != nil {
		panic(err)
	}

	// server
	s := httpserver.NewServer(httpserver.Name("danerys.test.service"), httpserver.Tracer(tracer))

	for k, v := range handlers {
		s.ANY(k, v)
	}

	go func() {
		err := s.Run(fmt.Sprintf(":%d", serverPort))
		if err != nil {
			fmt.Println(err)
		}
	}()

	var c = `[server]
service_name="inf.framebenchmark.logic"
port=60088

[[server_client]]
service_name="inf.framebenchmark.base"
proto="http"
#endpoints_from="consul"
endpoints="ali-c-inf-testing01.bj:60092"
balancetype="roundrobin"
read_timeout=60000
retry_times=0
slow_time=10000

[[redis]]
server_name="framebenchmark.logic.redis"
addr="10.100.130.14:6379"
password=""
max_idle=1000
max_active=5000
idle_timeout=60000
connect_timeout=1000
read_timeout=1000
write_timeout=1000
database=0
slow_time=10

[[kafka_producer_client]]
#produce标识 与 consumeFrom 对应
producer_to="framebenchmark.backend.kafka"
#broker地址
kafka_broken="ali-a-inf-kafka-test11.bj:9092,ali-c-inf-kafka-test12.bj:9092,ali-a-inf-kafka-test13.bj:9092,ali-c-inf-kafka-test14.bj:9092,ali-e-inf-kafka-test15.bj:9092"
use_sync=true
request_timeout=500
required_acks="WaitForAll"
get_success=true
get_error=true
retrymax=1`

	consulAddr = "test.consul.inkept.cn:8500"
	dc, err := sd.GetDatacenter(consulAddr)
	if err != nil {
		t.Errorf("error:%s", err)
		t.Fail()
	}

	body := make(map[string]interface{})
	body["content"] = c
	body["service"] = "inf.framebenchmark.logic"
	body["cluster"] = dc
	body["path"] = "/backup/config.toml"
	bodyB, _ := json.NewEncoder().Encode(body)

	respB, err := tracing.KVPut(bodyB)
	assert.Equal(t, nil, err)
	assert.Equal(t, `{"dm_error":0,"error_msg":"success"}`, strings.Trim(string(respB[:]), "\n"))
}

func TestMiddleware(t *testing.T) {
	func1 := func() error {
		time.Sleep(500 * time.Millisecond)
		return nil
	}

	func2 := func() error {
		time.Sleep(500 * time.Millisecond)
		return nil
	}

	func3 := func() error {
		time.Sleep(2 * time.Second)
		return nil
	}

	func4 := func() error {
		time.Sleep(500 * time.Millisecond)
		return nil
	}

	middlewares := map[string]func() error{
		"kafkaProducerInit": func1,
		"kafkaConsumerInit": func2,
		"redisClientInit":   func3,
		"mysqlClientInit":   func4,
	}

	for name, fn := range middlewares {
		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		fnDone := make(chan error)
		go func() {
			fnDone <- fn()
		}()
	INNER:
		for {
			select {
			case <-ctx.Done():
				cancel()
				fmt.Printf("doing %s timeout, please check your config\n", name)
				return
			case err := <-fnDone:
				if err != nil {
					cancel()
					fmt.Println(err)
					return
				}
				fmt.Println("init success", name)
				break INNER
			}
		}
		cancel()
	}
}

func Test_makeUploadPath(t *testing.T) {
	p1 := "."
	p2 := "./"
	p3 := "/"
	p4 := "../"
	p5 := "./../"
	p6 := "app"
	p7 := "./app"
	p8 := "../../../app"
	p9 := "./xxx"
	p10 := "./config/app/config"

	 list := []string{p1, p2, p3, p4, p5, p6, p7, p8, p9, p10}
	for _, p := range list {
		filename := filepath.Join(p, "inke", "config.toml")
		fmt.Println(makeUploadPath(p, filename))
	}

	fp := "config/app/config/abc-test/config.toml"
	fmt.Println(makeUploadPath(p10, fp))
}
