package main

import (
	"fmt"
	"github.com/lfxnxf/frame/logic/rpc-go/naming/manager"
	"time"
)

func doTest() {

	ncm := manager.NewConsuleManager()

	ncm.Start()

	target := "link.business.unreadrpc"
	proto := "http"
	ncm.Watch(target, proto, "", "")

	go func() {
		for {
			msg, _ := ncm.Next()
			fmt.Println(msg)
		}
	}()

}

func main() {

	doTest()

	fmt.Print("d")
	time.Sleep(1 * time.Hour)

}
