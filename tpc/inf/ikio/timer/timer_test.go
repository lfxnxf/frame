package timer

import (
	//"fmt"
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"

	"golang.org/x/net/context"
)

type a struct {
	i int
}

var r = rand.New(rand.NewSource(time.Now().Unix()))

func (a *a) Id() TimerId {
	return TimerId(r.Intn(10000000))
}

func (a *a) Value() interface{} {
	return nil
}

func TestTimerAdd2(t *testing.T) {
	tm := NewTimingWheel(context.TODO())
	m := map[*Context]context.CancelFunc{}
	for i := 0; i < 10; i++ {
		ctx, _ := tm.AddTimer(5*time.Second, time.Second, WithValue(nil, i))
		m[ctx] = nil
	}
	//start := time.Now()
	wg := sync.WaitGroup{}
	for ctx := range m {
		ctx := ctx
		wg.Add(1)
		if (ctx.Value(nil).(int) % 2) == 0 {
			tm.CancelTimer(ctx)
		}
	}

	wg.Add(1)
	go func() {
		for ctx := range tm.TimeOutChannel() {
			fmt.Printf("ch %+v\n", ctx.Value(nil))
		}
		wg.Done()
	}()

	wg.Wait()
}

func TestTimerAdd(t *testing.T) {
	tm := NewTimingWheel(context.TODO())
	wg := &sync.WaitGroup{}
	wg.Add(1)

	m := make(map[int]struct{})
	go func() {
		for ctx := range tm.TimeOutChannel() {
			m[ctx.Value(nil).(int)] = struct{}{}
			if len(m) == 100 {
				wg.Done()
			}
		}
	}()

	//_, cancle,  _ := tm.AddTimer(5 * time.Second, true, WithValue(nil, &a{0}))
	for i := 0; i < 100; i++ {
		tm.AddTimer(5*time.Second, time.Second)
	}
	//time.Sleep(2 * time.Second)
	//cancle()
	wg.Wait()
	tm.Stop()

	for i := 0; i < 100; i++ {
		_, ok := m[i]
		if !ok {
			panic(i)
		}
	}
}

func BenchmarkTimerAdd(t *testing.B) {
	t.N = 150000
	tm := NewTimingWheel(context.TODO())

	go func() {
		for range tm.TimeOutChannel() {
			//cb.Callback(time.Now(), nil)
			//println("??")
		}
	}()

	for n := 0; n < t.N; n++ {
		tm.AddTimer(2*time.Second, time.Second, WithValue(nil, &a{0}))
	}
	t.ReportAllocs()
}

func testShuffle(vals []*Context) []*Context {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	ret := make([]*Context, len(vals))
	n := len(vals)
	for i := 0; i < n; i++ {
		randIndex := r.Intn(len(vals))
		ret[i] = vals[randIndex]
		vals = append(vals[:randIndex], vals[randIndex+1:]...)
	}
	return ret
}

func BenchmarkTimerDelete(t *testing.B) {
	t.N = 150000

	tm := NewTimingWheel(context.TODO())
	wg := &sync.WaitGroup{}
	wg.Add(1)

	go func() {
		for range tm.TimeOutChannel() {
			//cb.Callback(time.Now(), nil)
		}
	}()

	indexs := make([]*Context, t.N)

	for n := 0; n < t.N; n++ {
		ctx, _ := tm.AddTimer(time.Second, time.Second, WithValue(nil, &a{}))
		indexs[n] = ctx
	}

	indexs = testShuffle(indexs)

	for n := 0; n < t.N; n++ {
		indexs[n] = indexs[t.N-1-n]
	}

	t.ResetTimer()
	for n := 0; n < t.N; n++ {
		tm.CancelTimer(indexs[n])
	}
	t.ReportAllocs()
}
