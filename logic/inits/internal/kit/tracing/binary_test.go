package tracing

import (
	"github.com/lfxnxf/frame/logic/inits/log"
	opentracing "github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/mocktracer"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBinaryToHTTP(t *testing.T) {
	tracer := mocktracer.New()
	tracer.RegisterExtractor(
		opentracing.HTTPHeaders, &mocktracer.TextMapPropagator{true},
	)

	h := make(map[string]string)
	ctx := BinaryToContext(tracer, h, "testOp", log.Stdout())

	span := opentracing.SpanFromContext(ctx)
	span.Finish()
	assert.Equal(t, len(h), 0)

	finishedSpans := tracer.FinishedSpans()
	assert.Equal(t, 1, len(finishedSpans))

	endpointSpan := finishedSpans[0]
	assert.Equal(t, "testOp", endpointSpan.OperationName)
}
