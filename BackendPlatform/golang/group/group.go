package group

import (
	"context"
	"fmt"
	"log"
	"os"
	"runtime"
	"sync"

	"github.com/lfxnxf/frame/BackendPlatform/golang/common"
	"github.com/lfxnxf/frame/BackendPlatform/golang/logging"
	"github.com/lfxnxf/frame/tpc/inf/go-tls"
	"github.com/lfxnxf/frame/tpc/inf/metrics"
)

// A Group is a collection of goroutines working on subtasks that are part of
// the same overall task. go function will run in tls.For with panic recover and report.
// group can use GOMAXPROCS function to set max goroutine to work.

type call struct {
	ctx context.Context
	f   func(ctx context.Context) error
}

type Group struct {
	err     error
	wg      sync.WaitGroup
	errOnce sync.Once

	workerOnce sync.Once
	ch         chan call
	chs        []call
}

func (g *Group) do(ctx context.Context, f func(ctx context.Context) error) {
	var err error
	defer func() {
		if r := recover(); r != nil {
			buf := make([]byte, 64<<10)
			buf = buf[:runtime.Stack(buf, false)]
			err = fmt.Errorf("errgroup: panic recovered: %s\n%s", r, buf)
			// error log
			logging.For(ctx).Errorf("service panic: %s", err)
			// crash log
			logging.CrashLog(err)
			// stdout log
			log.Printf("service panic: %s\n", err)
			// stderr log
			_, _ = fmt.Fprintf(os.Stderr, "service panic: %s\n", err)
			// inke base 监控模板中有该点位监控
			metrics.Meter(common.ServicePanicMeterName, 1, "name", "go.group")
		}
		if err != nil {
			g.errOnce.Do(func() {
				g.err = err
			})
		}
		g.wg.Done()
	}()
	tls.For(ctx, func() {
		err = f(ctx)
	})()
}

// GOMAXPROCS set max goroutine to work.
func (g *Group) GOMAXPROCS(n int) *Group {
	if n <= 0 {
		n = 1
	}
	g.workerOnce.Do(func() {
		g.ch = make(chan call, n)
		for i := 0; i < n; i++ {
			go func() {
				for c := range g.ch {
					g.do(c.ctx, c.f)
				}
			}()
		}
	})
	return g
}

// Go calls the given function in a new goroutine within tls.For.
// Go will recover if any panic occurs.
// The _ctx will passed to f.
// The First error will be returned by Wait.
func (g *Group) Go(ctx context.Context, f func(ctx context.Context) error) {
	g.wg.Add(1)
	if g.ch != nil {
		select {
		case g.ch <- call{
			ctx: ctx,
			f:   f,
		}:
		default:
			g.chs = append(g.chs, call{
				ctx: ctx,
				f:   f,
			})
		}
		return
	}
	go g.do(ctx, f)
}

// Wait blocks until all function calls from the Go method have returned, then
// returns the first non-nil error (if any) from them.
func (g *Group) Wait() error {
	if g.ch != nil {
		for _, f := range g.chs {
			g.ch <- f
		}
	}
	g.wg.Wait()
	if g.ch != nil {
		close(g.ch) // let all receiver exit
	}
	return g.err
}
