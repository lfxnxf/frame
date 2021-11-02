package trace

import (
	"context"

	"github.com/lfxnxf/frame/tpc/inf/go-tls"
	"github.com/opentracing/opentracing-go"
)

const operationLimit = 512

// Context returns a ctx with root span
// only for backend task service
func Context(operationName ...string) context.Context {
	var operation = "Backend root"
	if len(operationName) > 0 && len(operationName[0]) < operationLimit {
		operation = operationName[0]
	}
	tracer := opentracing.GlobalTracer()
	span := tracer.StartSpan(operation)
	ctx := opentracing.ContextWithSpan(context.Background(), span)
	tls.SetContext(ctx)
	return ctx
}
