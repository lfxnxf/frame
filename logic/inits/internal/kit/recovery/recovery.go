package recovery

import (
	"fmt"
	"runtime/debug"

	"github.com/lfxnxf/frame/BackendPlatform/golang/logging"
	"github.com/lfxnxf/frame/logic/inits/internal/core"
	"github.com/lfxnxf/frame/tpc/inf/metrics"
	"golang.org/x/net/context"
)

func Recovery(recoverPanic bool) core.Plugin {
	return core.Function(func(ctx context.Context, c core.Core) {
		defer func() {
			if rc := recover(); rc != nil {
				logging.CrashLogf("recover panic info: %q, stacks info:\n%s", rc, debug.Stack())
				debug.PrintStack()
				metrics.Meter("plugin.recovery", 1)
				c.AbortErr(fmt.Errorf("%q", rc))
				if !recoverPanic {
					panic(rc)
				}
			}
		}()
		c.Next(ctx)
	})
}
