package main

import (
	"bytes"
	"context"
	"flag"
	"fmt"
	rpc "github.com/lfxnxf/frame/logic/rpc-go"
)

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

	for i := 0; i < 1; i++ {

		// post请求
		value := "{\"logid\":\"249,258,214,226,231\",\"b\":{\"id\":\"15048402117809736\",\"pass-through\":{\"token\":\"\"},\"ev\":\"c.jr\"},\"liveid\":\"1504835704167589\",\"gid\":\"Csj8cryJFIbrjqiR\",\"from\":\"game\",\"atom\":{\"ua\":\"iPhone9_2\",\"imei\":\"\",\"imsi\":\"\",\"uid\":\"397161827\",\"proto\":\"8\",\"idfa\":\"8EB71229-57DA-4BE3-AB2C-48CEE2C871A9\",\"devicetoken\":\"8ae219c746b3a3a1f058bbfeb3506b787bcee1cbd50a9d3ed7f039757ee9a2fc\",\"lc\":\"0000000000000071\",\"cc\":\"TG0001\",\"sid\":\"20Gf6C6FkqQNqKozei2mvS2jP9MXFHo9dP77mQBo2WLf7UKGTF1\",\"mtxid\":\"12696cbf4997\",\"cv\":\"IK5.0.00_Iphone\",\"mtid\":\"5c5ff1093757fe628465f1961b6436b2\",\"idfv\":\"4148D4E8-6902-49AC-905A-D1CD372457E5\",\"conn\":\"wifi\",\"devi\":\"48b257df13515cfd2790ec9599306808479bcf32\",\"logid\":\"249,258,214,226,231\",\"osversion\":\"ios_11.000000\",\"userip\":\"10.111.13.57\"},\"userid\":397161827,\"server_ip\":\"10.111.13.68\"}"

		body := bytes.NewReader([]byte(value))

		serviceName := "roombuss"
		uri := "/dispatcher/DispatcherService/Business"
		data, err := rpc.HttpPost(context.TODO(), serviceName, uri, nil, body)

		if err != nil {
			fmt.Printf("fail,http post response %q, %v\n", data, err)
		} else {
			fmt.Println("succ,http post retry ,%q", string(data))
		}
	}

	for i := 0; i < 1; i++ {

		serviceName := "service1"
		uri := "/api/live/simpleall?uid=106839"

		//自定义重试次数或者超时时间
		requestOptional := rpc.NewRequestOptional()
		requestOptional.SetTimeOut(10)
		requestOptional.SetRetryTimes(10)
		requestOptional.SetHeader("key", "value")

		data, err := rpc.HttpGet(context.TODO(), serviceName, uri, requestOptional)

		if err != nil {
			fmt.Printf("fail ,http with retry response %q, %v\n", data, err)
		} else {
			fmt.Println("succ ,http with reti")
		}
	}
}
