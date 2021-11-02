package mhttp

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
	"time"
)

func JsonDecode(data []byte, response interface{}) error {
	dec := json.NewDecoder(bytes.NewReader(data))
	dec.UseNumber()
	return dec.Decode(response)
}

func TestMHTTP(t *testing.T) {
	r1, _ := http.NewRequest("GET", "http://service.inke.cn/api/live/simpleall", nil)
	r2, _ := http.NewRequest("GET", "http://service.inke.cn/serviceinfo/info", nil)
	opts := []RequestOpts{
		RequestOpts{
			Req: r1,
			Dec: func(data []byte, e error) error {
				var resp map[string]interface{}
				err := JsonDecode(data, &resp)
				fmt.Printf("parse response %v", resp)
				return err
			},
			Timeout: 200 * time.Millisecond,
		},
		RequestOpts{
			Req: r2,
			Dec: func(data []byte, e error) error {
				var resp map[string]interface{}
				err := JsonDecode(data, &resp)
				fmt.Printf("parse response %v", resp)
				return err
			},
			Timeout: 200 * time.Millisecond,
		},
	}
	err := Execute(context.TODO(), opts)
	fmt.Printf("execute %v", err)
}
