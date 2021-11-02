package inits

import (
	"os"
	"runtime"
	"strconv"
	"sync"

	"github.com/lfxnxf/frame/BackendPlatform/golang/logging"
	"github.com/lfxnxf/frame/logic/inits/config"
	"github.com/lfxnxf/frame/logic/inits/internal/kit/tracing"
	clusterconfig "github.com/lfxnxf/frame/tpc/inf/go-upstream/config"
	"github.com/lfxnxf/frame/tpc/inf/go-upstream/registry"
	"github.com/lfxnxf/frame/tpc/inf/go-upstream/registry/consul"
)

var (
	consulAddr      string
	traceReportAddr string
	initOnce        sync.Once
)

const (
	_inkeApp         = "inke"
	_pprofURI        = "/debug/pprof/port"
	LOG_ROTATE_HOUR  = "hour"
	LOG_ROTATE_DAY   = "day"
	LOG_ROTATE_MONTH = "month"
)

func init() {
	var (
		fallbackConsulAddr      = "127.0.0.1:8500"
		fallbackTraceReportAddr = "127.0.0.1:6831"
		fallbackTraceAPIAddr    = "127.0.0.1:5778"
	)

	if addr, ok := os.LookupEnv("CONSUL_ADDR"); ok {
		fallbackConsulAddr = addr
	}
	if addr, ok := os.LookupEnv("TRACE_ADDR"); ok {
		fallbackTraceReportAddr = addr
	}
	if addr, ok := os.LookupEnv("TRACE_API_ADDR"); ok {
		fallbackTraceAPIAddr = addr
	}
	if cores, ok := os.LookupEnv("ALLOCATE_CPU_MILLICORES"); ok {
		n, _ := strconv.Atoi(cores)
		if n < 200 {
			runtime.GOMAXPROCS(2)
		} else if n < 300 {
			runtime.GOMAXPROCS(4)
		} else {
			runtime.GOMAXPROCS(8)
		}
	}

	consulAddr = fallbackConsulAddr
	traceReportAddr = fallbackTraceReportAddr
	tracing.InitTraceAPIAddr(fallbackTraceAPIAddr)

	config.ConsulAddr = consulAddr

	registry.Default, _ = consul.NewBackend(&clusterconfig.Consul{Addr: consulAddr, Scheme: "http", Logger: logging.Log(logging.GenLoggerName)})

}
