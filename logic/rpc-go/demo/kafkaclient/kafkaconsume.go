package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/lfxnxf/frame/BackendPlatform/golang/kafka"
	rpc "github.com/lfxnxf/frame/logic/rpc-go"
)

var closeChan chan bool

func init() {
	c := make(chan os.Signal, 1)
	closeChan = make(chan bool)
	go func() {
		<-c
		rpc.CloseAllClient()
		close(closeChan)
	}()
	signal.Notify(c, os.Interrupt)
}

func testConsume() {

	client, err := rpc.GetKafkaConsumeClient("kafka-test")
	if err != nil {
		fmt.Println("get kafka consume client error ,", err)
		return
	}

	go func() {
		for k := range client.Errors() {
			fmt.Println("sserrr::::", k)
		}
	}()

	cnt := 0
	for message := range client.GetMessages() {
		cnt++
		fmt.Printf("recv mesage id=%s, value=%q, key=%q, Count=%d\n", message.MessageID, message.Value, message.Key, cnt)
		for _, h := range message.Headers {
			fmt.Printf("key %q value %q", h.Key, h.Value)
		}
		client.CommitUpto(message)
	}
}

type Config struct {
}

func main() {

	conf := flag.String("config", "./config.toml", "rpc config file")
	flag.Parse()

	// 读取配置文件
	var cc Config
	_, err := rpc.NewConfigToml(*conf, &cc)
	if err != nil {
		fmt.Printf("New Config error %s\n", err)
	}

	testConsume()
	testProducter()
	time.Sleep(10 * time.Second)
}

func testProducter() {

	client, err := rpc.GetSyncProducerClient("kafka-test")

	if err != nil {
		fmt.Println("get producter client error ,err:", err)
		return
	}
	successCnt := 0
	for i := 0; i < 120000; i++ {
		m := &kafka.ProducerMessage{
			Topic:     "test",
			Key:       "key",
			Value:     []byte(fmt.Sprintf("msg value %d", i)),
			MessageID: fmt.Sprintf("%d", i),
		}
		_, _, err := client.Send(context.TODO(), m)
		if err == nil {
			successCnt++
		}
		//fmt.Printf("send message %q, err %v\n", m.MessageID, err)
		time.Sleep(time.Duration(1000) * time.Millisecond)
	}
	fmt.Printf("send success cnt %d\n", successCnt)
}
