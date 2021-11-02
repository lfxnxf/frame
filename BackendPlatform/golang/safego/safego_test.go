package safego

import (
	"context"
	"fmt"
	"log"
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGo(t *testing.T) {
	type args struct {
		ctx context.Context
		f   func(ctx context.Context)
	}

	k := "key"
	ctx := context.WithValue(context.Background(), k, "1")

	var count int32
	f := func(ctxVal context.Context) {
		atomic.AddInt32(&count, 1)
		v := ctxVal.Value(k)
		if v == nil {
			t.Fatal("context val nil")
		}
		vs, ok := v.(string)
		if !ok {
			t.Fatal("context val type error")
		}
		if vs != "1" {
			t.Fatal("not 1")
		}

		log.Println("func done")
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "go1", args: args{
			ctx: ctx,
			f:   f},
		},
		{
			name: "go2", args: args{
			ctx: ctx,
			f:   f},
		},
		{
			name: "go3", args: args{
			ctx: ctx,
			f:   f},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Go(tt.args.ctx, tt.args.f)
		})
	}
	time.Sleep(20 * time.Millisecond)
	assert.Equal(t, int32(len(tests)), atomic.LoadInt32(&count))
}

func TestGo_isSafe(t *testing.T) {
	f := func(ctx context.Context) {
		panic("this is a testing print")
	}
	Go(context.Background(), f)
	time.Sleep(10 * time.Millisecond)
	fmt.Println("function done")
}
