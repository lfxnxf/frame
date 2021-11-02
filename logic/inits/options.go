package inits

import (
	"github.com/lfxnxf/frame/logic/inits/log"
	"github.com/opentracing/opentracing-go"
)

type Option func(*Inits)

// Mode TODO
type Mode int

const (
	// Deprecated: Development func should not use anymore.
	Development Mode = iota // 0
	// Deprecated: Production func should not use anymore.
	Production // 1
)

// Deprecated: String func should not use anymore.
func (m *Mode) String() string {
	switch *m {
	case Development:
		return "development"
	case Production:
		return "production"
	default:
		return "unknown"
	}
}

// Deprecated: Kit func should not use anymore.
func Kit(kit log.Kit) Option {
	return func(o *Inits) {
		//o.Kit = kit
	}
}

// Deprecated: RunMode func should not use anymore.
func RunMode(mode Mode) Option {
	return func(o *Inits) {
		//o.RunMode = mode
	}
}

func Namespace(namespace string) Option {
	return func(o *Inits) {
		o.Namespace = namespace
	}
}

func Name(name string) Option {
	return func(o *Inits) {
		o.Name = name
	}
}

func App(app string) Option {
	return func(o *Inits) {
		o.App = app
	}
}

func Version(ver string) Option {
	return func(o *Inits) {
		o.Version = ver
	}
}

func Deps(deps string) Option {
	return func(o *Inits) {
		o.Deps = deps
	}
}

// Deprecated: Tracer func should not use anymore.
func Tracer(tracer opentracing.Tracer) Option {
	return func(o *Inits) {
		//o.Tracer = tracer
	}
}

func ConfigPath(path string) Option {
	return func(o *Inits) {
		o.ConfigPath = path
	}
}

func NamespaceDir(dir string) Option {
	return func(o *Inits) {
		o.namespaceDir = dir
	}
}

func ConfigMemory(mem []byte) Option {
	return func(o *Inits) {
		o.ConfigMemory = mem
	}
}

// Deprecated: ConsulAddr func should not use anymore.
// 可以使用用户输入: -consul-addr="127.0.0.1:8500" 或者环境变量: "CONSUL_ADDR"
func ConsulAddr(addr string) Option {
	return func(o *Inits) {
		//o.ConsulAddr = addr
	}
}

// Deprecated: TraceReportAddr func should not use anymore.
// 可以使用用户输入: -trace-addr="127.0.0.1:6831" 或者环境变量: "TRACE_ADDR"
func TraceReportAddr(addr string) Option {
	return func(o *Inits) {
		//o.TraceReportAddr = addr
	}
}
