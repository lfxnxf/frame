package code

import "github.com/lfxnxf/frame/BackendPlatform/golang/ecode"

var (
	InvalidParam      = ecode.New(10000)
)

func init() {
	ecode.Register(map[int]string{
		10000: "xxxx错误",
	})
}

