package errs

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	err := New(0, "操作成功")
	buf, _ := json.Marshal(err)
	assert.Equal(t, string(buf), `{"dm_error":0,"error_msg":"操作成功"}`)
}
func TestNewResp(t *testing.T) {
	resp := NewResp(ErrInternal, nil)
	buf, _ := json.Marshal(resp)
	assert.Equal(t, string(buf), `{"dm_error":500,"error_msg":"内部系统错误"}`)
}

func TestResp(t *testing.T) {
	resp, code := Resp(ErrInternal, "test")
	assert.Equal(t, code, 500)
	buf, _ := json.Marshal(resp)
	assert.Equal(t, string(buf), `{"dm_error":500,"error_msg":"内部系统错误","data":"test"}`)
}
