package main

import (
	// "bytes"
	"context"
	// "flag"
	"fmt"
	rpc "github.com/lfxnxf/frame/logic/rpc-go"
)

type Config struct {
}

func main() {

	var scs []rpc.ServerClient

	var sc rpc.ServerClient
	sc.ServiceName = "service1"
	sc.Ipport = "116.211.167.106:80"
	sc.ProtoType = "http"
	sc.Balancetype = "random"
	sc.ReadTimeout = 3000
	sc.RetryTimes = 3
	scs = append(scs, sc)
	//需要自行的进行init,loadbalance
	errInit := rpc.InitLoadBalance(scs)

	if errInit != nil {
		fmt.Println("init error ")
		return
	}

	for i := 0; i < 2; i++ {

		//get 请求
		serviceName := "service1"
		uri := "/api/live/simpleall?uid=106839"
		data, err := rpc.HttpGet(context.TODO(), serviceName, uri, nil)

		if err != nil {
			fmt.Printf("fail response %q, %v\n", data, err)
		} else {
			fmt.Println("succ ")
		}
	}

}
