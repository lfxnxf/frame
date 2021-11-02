package group

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

var _ctx = context.Background()

func TestGroup_GOMAXPROCS(t *testing.T) {
	ns := []int{2, 5, 10, 20, 80, 100, 1, 8}
	for _, n := range ns {
		group := Group{}
		group.GOMAXPROCS(n)
		count := int64(0)

		for i := 0; i < n+10; i++ {
			group.Go(_ctx, func(c context.Context) error {
				atomic.AddInt64(&count, 1)
				time.Sleep(200 * time.Millisecond)
				return nil
			})
		}
		time.Sleep(20 * time.Millisecond)
		if atomic.LoadInt64(&count) != int64(n) {
			t.Fatalf("GOMAXPROCS not works: need=%d, got=%d", n, atomic.LoadInt64(&count))
		}
		if err := group.Wait(); err != nil {
			t.Fatal(err)
		}
	}
}

func TestGroup_Go(t *testing.T) {

	k := "key"
	ctxVal := context.WithValue(_ctx, k, "1")
	group := Group{}
	group.GOMAXPROCS(10)
	for i := 0; i < 10; i++ {
		group.Go(ctxVal, func(c context.Context) error {
			v := ctxVal.Value(k)
			if v == nil {
				t.Fatal("context val nil")
			}
			vs, ok := v.(string)
			if !ok {
				t.Fatal("context val type error")
			}
			if vs != "1" {
				t.Fatal("context val error")
			}
			return nil
		})
	}
	if err := group.Wait(); err != nil {
		t.Fatal(err)
	}
}

func TestGroup_Go2(t *testing.T) {
	ctxVal, cancel := context.WithCancel(_ctx)
	group := Group{}
	group.GOMAXPROCS(10)
	for i := 0; i < 10; i++ {
		group.Go(ctxVal, func(c context.Context) error {
			for {
				select {
				case <-ctxVal.Done():
					return nil // returning not to leak the goroutine
				default:
					fmt.Println("alive")
				}
				time.Sleep(1 * time.Second)
			}
		})
	}

	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		defer wg.Done()
		time.Sleep(10 * time.Millisecond)
		cancel()
		fmt.Println("done")
	}()
	if err := group.Wait(); err != nil {
		t.Fatal(err)
	}
	wg.Wait()
}
