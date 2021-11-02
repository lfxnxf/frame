package redis_test

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/lfxnxf/frame/BackendPlatform/golang/redis"
	"github.com/lfxnxf/frame/BackendPlatform/miniredis"
)

var (
	r               *redis.Redis
	miniredisServer *miniredis.Miniredis
)

func init() {
	var err error
	miniredisServer, err = miniredis.Run()
	r, err = redis.NewRedis(&redis.RedisConfig{
		ServerName:     "test",
		Addr:           fmt.Sprintf("localhost:%s", miniredisServer.Port()),
		MaxIdle:        100,
		MaxActive:      100,
		IdleTimeout:    0,
		ConnectTimeout: 200,
		ReadTimeout:    100, //ms
		WriteTimeout:   100, //ms
		Database:       0,
	})
	if err != nil {
		log.Fatalf("init: %s\n", err)
	}
}

func ExampleBool() {
	miniredisServer.FlushAll()
	r.Do("SET", "foo", 1)
	exists, err := redis.Bool(r.Do("EXISTS", "foo"))
	if err != nil {
		log.Fatalf("err=%+v\n", err)
	}
	fmt.Printf("%#v\n", exists)
	// Output:
	// true
}

func ExampleInt() {
	miniredisServer.FlushAll()
	r.Do("SET", "k1", 1)
	n, _ := redis.Int(r.Do("GET", "k1"))
	fmt.Printf("%#v\n", n)
	n, _ = redis.Int(r.Do("INCR", "k1"))
	fmt.Printf("%#v\n", n)
	// Output:
	// 1
	// 2
}

func ExampleInts() {
	miniredisServer.FlushAll()
	r.Do("SADD", "set_with_integers", 4, 5, 6)
	ints, _ := redis.Ints(r.Do("SMEMBERS", "set_with_integers"))
	fmt.Printf("%#v\n", ints)
	// Output:
	// []int{4, 5, 6}
}

func ExampleString() {
	miniredisServer.FlushAll()
	r.Do("SET", "hello", "world")
	s, err := redis.String(r.Do("GET", "hello"))
	if err != nil {
		log.Fatalf("get string err=%+v\n", err)
	}
	fmt.Printf("%#v\n", s)
	// Output:
	// "world"
}

func ExampleRedis_Set() {
	miniredisServer.FlushAll()
	if ok, err := r.Set("TestSet", "this is value"); err != nil {
		fmt.Println(err)
	} else if !ok {
		log.Fatalf("not ok")
	}
}

func ExampleRedis_Do() {
	miniredisServer.FlushAll()
	reply, err := r.Do("config", "get", "*")
	if err != nil {
		log.Fatalf("Do: %s\n", err)
	}
	m, _ := redis.StringMap(reply, err)
	for k, v := range m {
		fmt.Printf("%s:%s", k, v)
	}
}

func ExampleRedis_For() {
	// 这个例子里的ctx使用了context.TODO(), 仅仅为了说明如何使用
	// 在实际代码中需要使用从基础库里传递进来的context
	miniredisServer.FlushAll()
	reply, err := r.For(context.TODO()).Do("config", "get", "*")
	if err != nil {
		log.Fatalf("Do: %s\n", err)
	}
	m, _ := redis.StringMap(reply, err)
	for k, v := range m {
		fmt.Printf("%s:%s", k, v)
	}
}

func ExampleRedis_DoCtx() {
	miniredisServer.FlushAll()
	timeout := 20 * time.Millisecond
	ctx, _ := context.WithTimeout(context.Background(), timeout)
	reply, err := r.DoCtx(ctx, "BLPOP", "NotExistKey", 0)
	if err != context.DeadlineExceeded {
		log.Fatalf("TestTimeOut err: %s %v\n", err, reply)
	}
}

func ExampleRedis_Send() {
	miniredisServer.FlushAll()
	closes := make(chan struct{})
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	go func() {
		<-c
		close(closes)
	}()

	go func() {
		for i := 0; ; i++ {
			if err := r.Send("queue_name", "this is value", i); err != nil {
				log.Fatalf("err:%s", err)
			}
			time.Sleep(time.Second * 1)
		}
	}()
}

func ExampleRedis_Receive() {
	bufSize := 10
	// "queue_name" is redis queue, "closes" is used to close queue, see Send example
	// bufSize is the size of channal buffer size
	closes := make(chan struct{})
	for b := range r.Receive("queue_name", closes, bufSize) {
		b := b
		go func() {
			fmt.Printf("receive=%s\n", b)
		}()
	}
}

func ExampleRedis_Do_timeout() {
	_, err := r.Do("BLPOP", "NotExistKey", 1)
	if err != redis.ErrTimeout {
		log.Fatalf("Do: %s\n", err)
	}
}

func ExampleRedis_ZAdd() {
	if res, err := r.ZAdd("TestZSet", 1, "m1", 2, "m2"); err != nil {
		log.Fatalf("TestZSet %s\n", err)
	} else if res != 1 {
		log.Fatalf("result not match res=%v\n", res)
	}
	if _, err := r.ZAdd("TestZSet", 1, "m1", 2, "m2", 4); err == nil {
		log.Fatalf("TestZSet err\n")
	}
}

func ExamplePipelining() {
	p, err := r.NewPipelining(context.TODO())
	if err != nil {
		log.Fatalf("err=%+v\n", err)
	}

	// 不要忘记关闭piplining
	defer p.Close()

	// 把命令写入缓冲区
	if err := p.Send("SET", "key", "value"); err != nil {
		log.Fatalf("err=%+v\n", err)
	}
	if err := p.Send("GET", "key2", "value2"); err != nil {
		log.Fatalf("err=%+v\n", err)
	}

	// 把缓冲区中的内容写入到网络
	if err := p.Flush(); err != nil {
		log.Fatalf("err=%+v\n", err)
	}

	// 处理SET命令
	if ok, err := redis.String(p.Receive()); err != nil || ok != "OK" {
		log.Fatalf("err=%+v\n", err)
	}

	// 处理GET命令
	if value2, err := redis.String(p.Receive()); err != nil {
		log.Fatalf("err=%+v\n", err)
	} else {
		// 处理values2 ....
		fmt.Printf("%+v\n", value2)
	}
}
