package req

import (
	"testing"
	"net/http"
	"net/url"
	"github.com/stretchr/testify/assert"
)

type req struct {
	A string `json:"a"`
	C int    `json:"c"`
}

func TestRequestDecode(t *testing.T) {
	r, _ := url.ParseRequestURI("http://inke.cn/hello/world?a=b&c=2")
	h := &http.Request{URL: r, Method: "GET"}
	var m req
	err := ReqDecode(h, &m)
	assert.Equal(t, m, req{A: "b", C: 2})
	t.Log(m, err)
}
