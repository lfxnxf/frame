package school_errors

import (
	"fmt"
	"github.com/lfxnxf/frame/BackendPlatform/golang/ecode"
	"github.com/lfxnxf/frame/BackendPlatform/golang/utils"
	"go.uber.org/multierr"
	"strings"
	"sync"
)

// 重定义错误码，以便增加新的支持
type Code struct {
	code ecode.Code
	err  error
}

func (c Code) Unwrap() error {
	return c.err
}


var (
	toasts = make(map[int]error)
	mux    sync.RWMutex
)

func (c Code) AddError(err error) Code {
	newC := c
	newC.err = multierr.Append(newC.err, err)
	return newC
}

func (c Code) Error() string {
	if c.err == nil {
		return c.code.Error()
	}
	return c.code.Error() + ";" + c.err.Error()
}

// Code return error code
func (c Code) Code() int {
	return c.code.Code()
}

// Message return error message
func (c Code) Message() string {
	return c.code.Message()
}

// Equal for compatible.
func (c Code) Equal(err error) bool {
	return c.code.Equal(err)
}

// 设置用户提示文案
func (c Code) SetToast(t string) Code {
	mux.Lock()
	defer mux.Unlock()

	if _, ok := toasts[c.Code()]; ok {
		panic(fmt.Sprintf("ecode: %d for toast already exist", c))
	}
	toasts[c.Code()] = &errorDetailInfo{
		code:  c,
		toast: t,
		err:   nil,
	}
	return c
}

func (c Code) HasToast() bool {
	mux.RLock()
	_, ok := toasts[c.Code()]
	mux.RUnlock()
	return ok
}

// 返回拥有用户提示文案的错误码
func (c Code) Toast() error {
	mux.RLock()
	if err, ok := toasts[c.Code()]; ok {
		mux.RUnlock()
		return err
	}
	mux.RUnlock()

	// 错误码不在错误集合内，则使用默认逻辑
	code := c.Code()
	//if code <= 399 {
	//	return &errorDetailInfo{
	//		code:  c,
	//		toast: c.Message(),
	//		err:   nil,
	//	}
	//}
	if code >= 400 && code <= 499 {
		return &errorDetailInfo{
			code:  c,
			toast: "网络异常，请重试",
			err:   nil,
		}
	}
	if code >= 500 && code <= 599 {
		return &errorDetailInfo{
			code:  c,
			toast: "系统开了小差，请稍后再试",
			err:   nil,
		}
	}
	//return &errorDetailInfo{
	//	code:  c,
	//	toast: "网络出现问题，请稍后再试",
	//	err:   nil,
	//}
	// 其他业务错误码使用原有提示信息
	return c
}

// 错误码拼接详细错误信息
func (c Code) DetailF(f string, args ...interface{}) error {
	return &errorDetailInfo{
		code:  c,
		toast: c.Message() + "(" + fmt.Sprintf(f, args...) + ")",
		err:   nil,
	}
}

// 错误码吗拼接详细错误信息
func (c Code) DetailW(params ...string) error {
	return &errorDetailInfo{
		code:  c,
		toast: c.Message() + "(" + strings.Join(params, ",") + ")",
		err:   nil,
	}
}

func genError(code int, msg string) Code {
	if code <= 499 || code >= 600 {
		// 框架bug，不支持，改成AddSBatchuccCode
		//utils.AddSuccCode(code)
		successCodes := make(map[int]int)
		successCodes[code] = 1
		utils.AddSBatchuccCode(successCodes)
	}
	return Code{
		code: ecode.Error(code, msg),
	}
}

// 添加错误码，不可重复添加同一错误码
func AddError(code int, msg string) Code {
	return genError(code, msg)
}

// 通过错误码查询错误
func Get(code int) Code {
	return Code{
		code: ecode.Int(code),
	}
}

func NewTmpError(code int, msg string) error {
	return &errorDetailInfo{
		code:  Get(code),
		toast: msg,
		err:   nil,
	}
}

func Toast(err error) error {
	switch err.(type) {
	case Code:
		//fmt.Println("Code")
		return err.(Code).Toast()
	case *errorDetailInfo:
		//fmt.Println("*errorDetailInfo")
		// 注释原因：部分 toast struct 是用于技术之间联调提示 detail message, examples Code.DetailF or Code.DetailW，本Toast函数用于返回给用户提示的信息，须强制使用设定的Toast
		//return err
		c := Get(err.(*errorDetailInfo).Code())
		// 错误码仓库中有对应的toast信息，则以仓库toast为准，否则使用原生
		if c.HasToast() {
			return c.Toast()
		}
		return err
	case ecode.Code:
		//fmt.Println("ecode.Code")
		return Get(err.(ecode.Code).Code()).Toast()
	}
	if c, ok := err.(ecode.Codes); ok {
		//fmt.Println("ecode.Codes")
		return Get(c.Code()).Toast()
	}
	//fmt.Println("unknown")
	return Codes.ServerError.Toast()
}

func DMError(err error) ecode.Codes {
	switch err.(type) {
	case Code:
		return err.(Code)
	case *errorDetailInfo:
		return err.(*errorDetailInfo)
	case ecode.Code:
		return err.(ecode.Code)
	}
	if c, ok := err.(ecode.Codes); ok {
		return c
	}
	return Codes.ServerError
}
