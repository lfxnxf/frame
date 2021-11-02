package tracing

import (
	"testing"

	"github.com/lfxnxf/frame/logic/inits/internal/kit/retry"

	"golang.org/x/net/context"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/mocktracer"

	"github.com/lfxnxf/frame/logic/inits/internal/core"
	"github.com/stretchr/testify/assert"
)

func TestTraceServer(t *testing.T) {
	tracer := mocktracer.New()

	// Initialize the ctx with a nameless Span.
	contextSpan := tracer.StartSpan("").(*mocktracer.MockSpan)
	ctx := opentracing.ContextWithSpan(context.Background(), contextSpan)

	c := core.New([]core.Plugin{TraceServer(tracer, "testOp")})
	c.Next(ctx)
	assert.Nil(t, c.Err())

	finishedSpans := tracer.FinishedSpans()
	assert.Equal(t, 1, len(finishedSpans))

	// Test that the op name is updated
	endpointSpan := finishedSpans[0]

	assert.Equal(t, "testOp", endpointSpan.OperationName)
	contextContext := contextSpan.Context().(mocktracer.MockSpanContext)
	endpointContext := endpointSpan.Context().(mocktracer.MockSpanContext)
	// ...and that the ID is unmodified.
	assert.Equal(t, contextContext.SpanID, endpointContext.SpanID)
}

func TestTraceClient(t *testing.T) {
	tracer := mocktracer.New()

	// Initialize the ctx with a nameless Span.
	contextSpan := tracer.StartSpan("").(*mocktracer.MockSpan)
	ctx := opentracing.ContextWithSpan(context.Background(), contextSpan)

	c := core.New([]core.Plugin{TraceClient(tracer, "testOp", true), retry.Retry(0)})
	c.Next(ctx)
	assert.Nil(t, c.Err())

	finishedSpans := tracer.FinishedSpans()
	assert.Equal(t, 1, len(finishedSpans))

	// Test that the op name is updated
	endpointSpan := finishedSpans[0]

	assert.Equal(t, "testOp", endpointSpan.OperationName)
	contextContext := contextSpan.Context().(mocktracer.MockSpanContext)
	endpointContext := endpointSpan.Context().(mocktracer.MockSpanContext)
	// ...and that the ID is unmodified.
	assert.NotEqual(t, contextContext.SpanID, endpointContext.SpanID)
}
