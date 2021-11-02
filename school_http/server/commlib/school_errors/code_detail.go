package school_errors

type errorDetailInfo struct {
	code  Code
	toast string
	err   error
}

// Error return Code in string form
func (e *errorDetailInfo) Error() string {
	if e.err == nil {
		//return strconv.FormatInt(int64(e.code), 10)+":"+e.toast
		return e.code.Error()
	}
	return e.err.Error()
}

// Code get error code.
func (e *errorDetailInfo) Code() int {
	return e.code.Code()
}

// Message get code message.
func (e *errorDetailInfo) Message() string {
	return e.toast
}

// Equal for compatible.
func (e *errorDetailInfo) Equal(err error) bool {
	if e2, ok := err.(*errorDetailInfo); ok {
		return *e2 == *e
	}
	return false
}

// Unwrap error
func (e *errorDetailInfo) Unwrap() error {
	return e.err
}
