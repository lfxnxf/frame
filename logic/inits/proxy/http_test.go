package proxy

import (
	"encoding/json"
	"io/ioutil"
	"testing"
	"time"

	"github.com/lfxnxf/frame/logic/inits"
	httpclient "github.com/lfxnxf/frame/logic/inits/http/client"
	httpserver "github.com/lfxnxf/frame/logic/inits/http/server"
	"github.com/lfxnxf/frame/tpc/inf/go-upstream/config"
	"github.com/lfxnxf/frame/tpc/inf/go-upstream/upstream"
	"github.com/stretchr/testify/assert"
	"golang.org/x/net/context"
)

type jsonTestObj struct {
	DMErr  int    `json:"dm_error"`
	ErrMsg string `json:"error_msg"`
	Data   struct {
		Action string
	} `json:"data"`
}

func TestHTTP_JSONGetPost(t *testing.T) {
	getJson := jsonTestObj{
		DMErr:  0,
		ErrMsg: "操作成功",
		Data:   struct{ Action string }{Action: "get json"},
	}

	postJson := jsonTestObj{
		DMErr:  0,
		ErrMsg: "操作成功",
		Data:   struct{ Action string }{Action: "post json"},
	}

	testInitServer(
		"JSONGetPost.service",
		map[string]httpserver.HandlerFunc{
			"/json/get": func(c *httpserver.Context) {
				c.Raw(getJson, 0)
				return
			},
			"/json/post": func(c *httpserver.Context) {
				bodyB, err := ioutil.ReadAll(c.Request.Body)
				if err != nil {
					t.Fail()
					c.Raw(nil, 500)
					return
				}
				body := map[string]string{}
				err = json.Unmarshal(bodyB, &body)
				if err != nil {
					t.Fail()
					c.Raw(nil, 500)
					return
				}
				if body["arg"] != "post action" {
					t.Fail()
					c.Raw(nil, 500)
					return
				}
				c.Raw(postJson, 0)
				return
			},
			"/text": func(c *httpserver.Context) {
				bodyB, err := ioutil.ReadAll(c.Request.Body)
				if err != nil {
					t.Fail()
					c.Raw(nil, 500)
					return
				}
				body := map[string]string{}
				err = json.Unmarshal(bodyB, &body)
				if err != nil {
					t.Fail()
					c.Raw(nil, 500)
					return
				}
				if body["arg"] != "post action" {
					t.Fail()
					_, _ = c.Response.WriteString("bad args")
					return
				}
				_, _ = c.Response.WriteString("success")
				return
			},
		},
		22266,
	)
	time.Sleep(3 * time.Second)

	ns, d := testInitInits("HTTP_JSONGetPost", "./testconfig/", "./testconfig/config.toml", "")
	ctx := inits.WithAPPKey(context.Background(), ns)

	clusterName := "json-http-client"
	cluster := config.NewCluster()
	cluster.Name = clusterName
	cluster.StaticEndpoints = "localhost:22266"
	manager := upstream.NewClusterManager()
	_ = manager.InitService(cluster)

	httpClient := httpclient.NewClient(
		httpclient.Cluster(manager.Cluster(clusterName)),
		httpclient.ServiceName("json-http-client"),
	)

	_ = d.AddHTTPClient("json-http-client", httpClient)

	http := InitHTTP("json-http-client")

	jsonGetResponse := jsonTestObj{}
	err := http.JSONGet(ctx, "/json/get", nil, &jsonGetResponse)
	assert.Equal(t, err, nil)
	assert.Equal(t, jsonGetResponse, getJson)

	jsonPostResponse := jsonTestObj{}
	err = http.JSONPost(ctx, "/json/post", nil, map[string]string{"arg": "post action"}, &jsonPostResponse)
	assert.Equal(t, err, nil)
	assert.Equal(t, jsonPostResponse, postJson)

	rsp, err := http.Call(ctx, httpclient.NewRequest(ctx).WithPath("/text").WithStruct(map[string]string{"arg": "post action"}).WithMethod("POST"))
	assert.Equal(t, err, nil)
	contentType := rsp.GetHeader("Content-Type")
	assert.Equal(t, contentType, "text/plain; charset=utf-8")
	assert.Equal(t, rsp.String(), "success")
}

func TestHTTP_WithApp(t *testing.T) {
	h := InitHTTP("imcenter.message.base").WithApp("nvwa")
	assert.Equal(t, "nvwa.imcenter.message.base", h.name)
}
