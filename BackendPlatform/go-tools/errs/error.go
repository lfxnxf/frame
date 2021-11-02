package errs

import "fmt"

var (
	// ErrIsNil represents success
	ErrIsNil = New(0, "操作成功")
	// ErrIllegal represents param is dangerous
	ErrIllegal = New(498, "非法参数")
	// ErrParam represents param is wrong
	ErrParam = New(499, "参数错误")
	// ErrInternal represents service occur some error
	ErrInternal = New(500, "内部系统错误")
	// ErrSession represents check identify failed
	ErrSession = New(604, "session检查失败")
)

// DMError represents standard response
type DMError struct {
	DmError  int    `json:"dm_error"`
	ErrorMsg string `json:"error_msg"`
}

// New creates DMError
func New(code int, msg string) DMError {
	return DMError{
		DmError:  code,
		ErrorMsg: msg,
	}
}

// Error returns error string
func (e DMError) Error() string {
	return fmt.Sprintf("dm_error(%d),error_msg(%s)", e.DmError, e.ErrorMsg)
}

// DMResp represents common response
type DMResp struct {
	DMError
	Data interface{} `json:"data,omitempty"`
}

// NewResp returns a common response.
// if error is standard error,the error_code should be 500.
func NewResp(e error, data interface{}) DMResp {
	var err DMError
	var ok bool
	if e == nil {
		err = ErrIsNil
	} else {
		err, ok = e.(DMError)
		if !ok {
			err = New(500, e.Error())
		}
	}
	return DMResp{
		DMError: err,
		Data:    data,
	}
}

// Resp returns a common response
func Resp(e error, data ...interface{}) (DMResp, int) {
	var err DMError
	var ok bool
	if e == nil {
		err = ErrIsNil
	} else {
		err, ok = e.(DMError)
		if !ok {
			err = New(500, e.Error())
		}
	}

	resp := DMResp{DMError: err}
	if len(data) > 0 {
		resp.Data = data[0]
	}
	return resp, resp.DmError
}
func (e DMResp) Error() string {
	return fmt.Sprintf("dm_error(%d),error_msg(%s),data(%v)", e.DmError, e.ErrorMsg, e.Data)
}
