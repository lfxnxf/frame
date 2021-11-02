# minisql
在服务内启动一个mysql实例，用于单元测试


# 使用
```text
package main

import (
	"fmt"
	"time"

	"github.com/lfxnxf/frame/BackendPlatform/minisql"
)

func main() {
	// 启动实例
	server := minisql.Run()
	fmt.Println("地  址: 127.0.0.1")
	fmt.Printf("端  口: %v\n", server.Port)
	fmt.Println("用户名: root")
	fmt.Println("密  码: (空)")
	fmt.Println("默认库: test")
	time.Sleep(time.Second * 2)
	// 停止实例
	server.Stop()
}

// 地  址: 127.0.0.1
// 端  口: 61891
// 用户名: root
// 密  码: (空)
// 默认库: test

```