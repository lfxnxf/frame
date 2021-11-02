package proxy

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
	"time"

	httpserver "github.com/lfxnxf/frame/logic/inits/http/server"
	"github.com/stretchr/testify/assert"
)

var testRecoverPanicData = `
[server]
service_name="inf.recover.panic"
port=30201
recover_panic=true
[log]
level="debug"
logpath="logs"
rotate="hour"
`

func TestRecoverPanic(t *testing.T) {
	_, dae := testInitInits("TestRecoverPanic", "./testconfig3/", "./testconfig3/config.toml", testRecoverPanicData)

	srv := dae.HTTPServer()
	srv.ANY("/server/panic", func(c *httpserver.Context) {
		if c.Request.Method == "GET" {
			var m map[string]interface{}
			m["panic"] = map[string]interface{}{"recover_panic": true}
			c.JSON(m["panic"], nil)
		} else {
			c.JSON(nil, nil)
		}
	})

	go func() {
		if err := srv.Run(); err != nil {
			fmt.Println(err)
		}
	}()

	time.Sleep(time.Second * 1)

	fmt.Printf("start http get")

	resp1, err := http.Post("http://127.0.0.1:30201/server/panic", "text/plain; charset=utf-8", nil)
	assert.Equal(t, nil, err)
	assert.NotEqual(t, nil, resp1)

	fmt.Printf("resp1:%+v\n", resp1)

	assert.Equal(t, http.StatusOK, resp1.StatusCode)

	bodyBytes1, err := ioutil.ReadAll(resp1.Body)
	_ = resp1.Body.Close()
	fmt.Printf("bodyBytes1:%s\n", string(bodyBytes1))

	resp2, err := http.Get("http://127.0.0.1:30201/server/panic")
	assert.Equal(t, nil, err)
	assert.NotEqual(t, nil, resp2)

	fmt.Printf("resp2:%+v\n", resp2)

	assert.Equal(t, http.StatusInternalServerError, resp2.StatusCode)

	bodyBytes2, err := ioutil.ReadAll(resp2.Body)
	_ = resp2.Body.Close()
	fmt.Printf("bodyBytes2:%s\n", string(bodyBytes2))
}
