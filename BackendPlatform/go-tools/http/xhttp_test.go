package http

import (
	"io/ioutil"
	"testing"
)

func TestGet(t *testing.T) {
	resp, err := Get("http://httpbin.org/get")
	buf, err := ioutil.ReadAll(resp.Body)
	t.Log(string(buf))
	t.Log(err)
}

func TestPost(t *testing.T) {
	resp, err := Post("http://httpbin.org/post", "application/json", nil)
	buf, err := ioutil.ReadAll(resp.Body)
	t.Log(string(buf))
	t.Log(err)
}

func TestGetAndUnmarshal(t *testing.T) {
	var m map[string]interface{}
	err := GetAndUnmarshal("http://httpbin.org/get", &m)
	t.Log(m)
	t.Log(err)
}
func TestPostAndUnmarshal(t *testing.T) {
	var result struct {
		Args struct {
		} `json:"args"`
		Data string `json:"data"`
		Files struct {
		} `json:"files"`
		Form struct {
		} `json:"form"`
		Headers struct {
			Accept        string `json:"Accept"`
			ContentLength string `json:"Content-Length"`
			ContentType   string `json:"Content-Type"`
			Host          string `json:"Host"`
			UserAgent     string `json:"User-Agent"`
		} `json:"headers"`
		JSON   interface{} `json:"json"`
		Origin string      `json:"origin"`
		URL    string      `json:"url"`
	}
	err := PostAndUnmarshal("http://httpbin.org/post", []byte(`hello`), &result)
	t.Log(result)
	t.Log(err)
	m := map[string]interface{}{
		"hello": "world",
	}
	err = PostAndUnmarshal("http://httpbin.org/post", m, &result)
	t.Log(result)
	t.Log(err)
}
