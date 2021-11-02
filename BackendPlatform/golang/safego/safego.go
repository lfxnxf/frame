package safego

import (
	"context"
	"fmt"
	"os"
	"runtime"

	"github.com/lfxnxf/frame/BackendPlatform/golang/common"
	"github.com/lfxnxf/frame/BackendPlatform/golang/logging"
	"github.com/lfxnxf/frame/tpc/inf/go-tls"
	"github.com/lfxnxf/frame/tpc/inf/metrics"
)

func Go(ctx context.Context, f func(context.Context)) {
	go tls.For(ctx, func() {
		defer func() {
			if r := recover(); r != nil {
				buf := make([]byte, 64<<10)
				buf = buf[:runtime.Stack(buf, false)]
				err := fmt.Errorf("errgroup: panic recovered: %s\n%s", r, buf)
				// error log
				logging.For(ctx).Errorf("service panic: %s", err)
				// crash log
				logging.CrashLog(err)
				// stderr log
				_, _ = fmt.Fprintf(os.Stderr, "service panic: %s\n", err)
				// inke base 监控模板中有该点位监控
				metrics.Meter(common.ServicePanicMeterName, 1, "name", "safe.go")
			}
		}()
		f(ctx)
	})()
}
