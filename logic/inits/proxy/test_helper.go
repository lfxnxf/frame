package proxy

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/schema"

	"github.com/lfxnxf/frame/BackendPlatform/jaeger-client-go"
	jaegerconfig "github.com/lfxnxf/frame/BackendPlatform/jaeger-client-go/config"
	"github.com/lfxnxf/frame/logic/inits"
	httpserver "github.com/lfxnxf/frame/logic/inits/http/server"
)

//nolint:unused
func testInitInits(app, configPath, configFile, configContent string) (string, *inits.Inits) {
	// 创建目录
	err := os.MkdirAll(configPath, 0750)
	if err != nil {
		panic(err)
	}
	// 删除文件
	defer func() {
		_ = os.RemoveAll(configPath)
	}()

	// 创建配置文件
	fd, err := os.OpenFile(configFile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		panic(err)
	}
	// 关闭文件
	defer func() {
		_ = fd.Close()
	}()

	if configContent != "" {
		_, _ = fd.WriteString(configContent)
		inits.GlobalNamespace.Add(app, inits.ConfigPath(configFile))
	} else {
		inits.GlobalNamespace.Add(app)
	}

	return app, inits.GlobalNamespace.Get(app)
}

//nolint:unused
func testInitServer(serviceName string, handlers map[string]httpserver.HandlerFunc, serverPort int) {
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
	tracer, _, err := cfg.New(serviceName)
	if err != nil {
		panic(err)
	}

	// server
	s := httpserver.NewServer(httpserver.Name(serviceName), httpserver.Tracer(tracer))

	for k, v := range handlers {
		s.ANY(k, v)
	}

	go func() {
		err := s.Run(fmt.Sprintf(":%d", serverPort))
		if err != nil {
			fmt.Println(err)
		}
	}()

	time.Sleep(1 * time.Second)
}

func TestParseArguements(request *http.Request, v interface{}) error {
	schemaDecoder := schema.NewDecoder()
	schemaDecoder.IgnoreUnknownKeys(true)
	err := schemaDecoder.Decode(v, request.URL.Query())
	if err != nil {
		return err
	}
	return nil
}
