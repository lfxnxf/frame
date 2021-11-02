package concurrent

import (
	"errors"

	"golang.org/x/net/context"
)

var (
	emptyTaskError = errors.New("task empty")
)

type TaskFunc func(ctx context.Context) error

func Parnnel(ctx context.Context, fs []TaskFunc) error {
	l := len(fs)
	if l == 0 {
		return emptyTaskError
	}
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	ch := make(chan error)
	for _, t := range fs {
		go func(t TaskFunc) {
			ch <- t(ctx)
		}(t)
	}
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case e := <-ch:
			l--
			if e != nil {
				return e
			}
			if l <= 0 {
				return nil
			}
		}
	}
}
