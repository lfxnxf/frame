package http

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

var (
	// DefaultClient has time-out control
	DefaultClient *http.Client
)

func init() {
	DefaultClient = &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			DialContext: (&net.Dialer{
				Timeout:   3 * time.Second,   // 连接超时时间
				KeepAlive: 300 * time.Second, // 连接保持超时时间,300s就自动断开
			}).DialContext,
			MaxIdleConnsPerHost:   1000,              // 连接池host
			MaxIdleConns:          10000,             // client对与所有host最大空闲连接数总和
			IdleConnTimeout:       200 * time.Second, // 空闲连接在连接池中的超时时间
			TLSHandshakeTimeout:   10 * time.Second,  // TLS安全连接握手超时时间
			ExpectContinueTimeout: 1 * time.Second,   // 发送完请求到接收到响应头的超时时间
		},
	}
}
func Get(url string) (*http.Response, error) {
	return DefaultClient.Get(url)
}
func Post(url, contentType string, body io.Reader) (resp *http.Response, err error) {
	return DefaultClient.Post(url, contentType, body)
}

func GetAndUnmarshal(url string, result interface{}) error {
	resp, err := Get(url)
	if err != nil {
		return err
	}
	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(buf, &result)
	return err
}

func PostAndUnmarshal(url string, data interface{}, result interface{}) error {
	var body *bytes.Buffer
	b, ok := data.([]byte)
	if !ok {
		b, _ = json.Marshal(data)
	}
	body = bytes.NewBuffer(b)
	resp, err := Post(url, "", body)
	if err != nil {
		return err
	}
	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(buf, &result)
	return err
}
