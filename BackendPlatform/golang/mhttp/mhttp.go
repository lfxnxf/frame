package mhttp

import (
	"errors"
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/lfxnxf/frame/BackendPlatform/golang/concurrent"
	rpc "github.com/lfxnxf/frame/logic/rpc-go"
)

var (
	emptyRequestError = errors.New("request execute empty")
)

type DecodeFunc func([]byte, error) error

type RequestOpts struct {
	Req     *http.Request
	Dec     DecodeFunc
	Timeout time.Duration
}

func Execute(ctx context.Context, opts []RequestOpts) error {
	ts := []concurrent.TaskFunc{}
	for _, opt := range opts {
		opt := opt
		ts = append(ts, func(context.Context) error {
			return execute(ctx, opt)
		})
	}
	return concurrent.Parnnel(ctx, ts)
}

func execute(ctx context.Context, opts RequestOpts) error {
	ctx, cancel := context.WithTimeout(ctx, opts.Timeout)
	defer cancel()
	data, err := rpc.CallHTTP(ctx, opts.Req)
	if opts.Dec == nil {
		return nil
	}
	return opts.Dec(data, err)
}
