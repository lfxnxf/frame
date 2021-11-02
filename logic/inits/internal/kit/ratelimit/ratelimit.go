package ratelimit

import (
	"github.com/lfxnxf/frame/logic/inits/internal/core"
	"github.com/lfxnxf/frame/logic/inits/ratelimit"
	"github.com/pkg/errors"
	"golang.org/x/net/context"
)

func Limiter(limType int, namespace, clientname, resource string, config *ratelimit.Config) core.Plugin {
	return core.Function(func(ctx context.Context, flow core.Core) {
		if config == nil {
			return
		}
		lim := config.GetLimiter(limType, namespace, clientname, resource)
		if lim != nil && !lim.Allow() {
			err := flow.Err()
			if err != nil {
				err = errors.Wrap(err, ratelimit.ErrLimited.Error())
			} else {
				err = ratelimit.ErrLimited
			}
			flow.AbortErr(err)
		}
	})
}
