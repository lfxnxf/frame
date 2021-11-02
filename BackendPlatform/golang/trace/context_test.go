package trace

import (
	"io"
	"testing"

	"github.com/lfxnxf/frame/BackendPlatform/jaeger-client-go"
	"github.com/lfxnxf/frame/BackendPlatform/jaeger-client-go/config"
	"github.com/opentracing/opentracing-go"
	"github.com/stretchr/testify/assert"
)

func TestContext(t *testing.T) {
	_, closer := Init("room.backend.service")
	defer closer.Close()
	taskLen := 100
	for i := 0; i < taskLen; i++ {
		// Notice: Every iteration need create new trace context
		ctx := Context()
		span := opentracing.SpanFromContext(ctx)
		sp, _ := span.(*jaeger.Span)
		assert.Equal(t, "Backend root", sp.OperationName())
	}
	ctx := Context("HTTP Client Get /api/get")
	span := opentracing.SpanFromContext(ctx)
	sp, _ := span.(*jaeger.Span)
	assert.Equal(t, "HTTP Client Get /api/get", sp.OperationName())
}

// init global tracer
func Init(service string) (opentracing.Tracer, io.Closer) {
	cfg := &config.Configuration{
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: 1.0,
		},
		Reporter: &config.ReporterConfig{
			LogSpans: true,
		},
	}
	tracer, closer, err := cfg.New(service, config.Logger(jaeger.StdLogger))
	opentracing.SetGlobalTracer(tracer)
	if err != nil {
		panic(err)
	}
	return tracer, closer
}
