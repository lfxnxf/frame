package rpc

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/lfxnxf/frame/BackendPlatform/jaeger-client-go"
	"github.com/lfxnxf/frame/logic/inits"
	httpclient "github.com/lfxnxf/frame/logic/inits/http/client"
	"github.com/lfxnxf/frame/logic/rpc-go/tracing"
	"golang.org/x/net/context"
)

var defaultIpportFormat = string("{ipport}")

type HttpDResponse struct {
	// http的错误码
	HttpCode int
	// http的请求信息,可以做日志的打印
	SuccIp string
	FailIp []string
	// 返回的错误信息
	RspBody        []byte
	Method         string
	ServiceNameUrl string
}

type HttpDRequest struct {
	// 请求的方法名称
	Method string
	// 带serviceName的url例如:http://servicename/api/test/api?uid=1000
	ServiceNameUrl string
	// 备份的ip: 只有在ServiceNameUrl 没有获取到IP时候,使用该备份的ip进行请求
	BackUpIp string
	// 超时时间
	TimeoutMs int
	// 重试次数
	RetryTimes int
	HeaderMaps map[string]string
	// 请求的body
	Body []byte
}

type RequestOptionIntercace interface {
	SetTimeOut(timeout int)
	SetRetryTimes(retryTimes int)

	SetSlowTime(time int)
	GetSlowTime() int

	GetTimeOut() int
	GetRetryTimes() int

	SetHeader(key string, value string)
	GetHeaderMap() map[string]string

	SetMetricTag(key string, value string)
	GetMetricTags() map[string]string
}

type RequestOption struct {
	Timeout    int
	RetryTimes int
	SlowTime   int
	Proto      string
	HeaderMap  map[string]string
	ro         *httpclient.RequestOption
	metricTags map[string]string
}

func NewRequestOptional() *RequestOption {
	return &RequestOption{
		Timeout:    0,
		RetryTimes: 0,
		HeaderMap:  make(map[string]string),
		ro:         &httpclient.RequestOption{},
		metricTags: make(map[string]string),
	}
}

func (c *RequestOption) SetHeader(key string, value string) {
	if c.HeaderMap == nil {
		c.HeaderMap = make(map[string]string)
	}
	c.HeaderMap[key] = value
}

func (c *RequestOption) GetHeaderMap() map[string]string {
	if c.HeaderMap == nil {
		c.HeaderMap = make(map[string]string)
	}
	return c.HeaderMap
}

func (c *RequestOption) SetSlowTime(timeout int) {
	c.SlowTime = timeout
	c.ro.SlowTimeoutMS(timeout)
}

func (c *RequestOption) GetSlowTime() int {
	return c.SlowTime
}

func (c *RequestOption) SetTimeOut(timeout int) {
	c.Timeout = timeout
	c.ro.RequestTimeoutMS(timeout)
}

func (c *RequestOption) SetRetryTimes(retryTimes int) {
	c.RetryTimes = retryTimes
	c.ro.RetryTimes(retryTimes)
}

func (c *RequestOption) GetTimeOut() int {
	return c.Timeout
}

func (c *RequestOption) GetRetryTimes() int {
	return c.RetryTimes
}

func (c *RequestOption) SetMetricTag(key, value string) {
	if c.metricTags == nil {
		c.metricTags = make(map[string]string)
	}
	c.metricTags[key] = value
}

func (c *RequestOption) GetMetricTags() map[string]string {
	if c.metricTags == nil {
		c.metricTags = make(map[string]string)
	}
	return c.metricTags
}

func DoHttpDRequest(ctx context.Context, httpDrequest HttpDRequest) (HttpDResponse, error) {
	snUrl := httpDrequest.ServiceNameUrl
	method := httpDrequest.Method
	u, err := url.Parse(snUrl)
	if err != nil {
		return HttpDResponse{}, fmt.Errorf("parse servicename_url error:%v", snUrl)
	}
	sName := u.Host
	uri := u.RequestURI()
	ro := NewRequestOptional()
	ro.RetryTimes = httpDrequest.RetryTimes
	ro.Timeout = httpDrequest.TimeoutMs
	if httpDrequest.HeaderMaps == nil {
		httpDrequest.HeaderMaps = make(map[string]string)
	}
	for k, v := range httpDrequest.HeaderMaps {
		ro.SetHeader(k, v)
	}
	body := bytes.NewReader(httpDrequest.Body)
	return callRespObj(ctx, sName, method, uri, ro, body)
}

// CallHTTP call request with business log
func CallHTTP(ctx context.Context, request *http.Request) ([]byte, error) {
	ctx = tracing.CaptureTraceContext(ctx)
	client := inits.HTTPClient(ctx, "")
	if client == nil {
		return nil, fmt.Errorf("get http client failed")
	}
	req := request.WithContext(ctx)
	r := httpclient.NewRequest().WithRequest(req)
	result, err := callbyCustomRequest(ctx, client, r)
	if err != nil {
		return nil, err
	}
	return result.RspBody, nil
}

// CallHTTPBackend call request with business log, and ingnore parent context cancel
func CallHTTPBackend(ctx context.Context, request *http.Request) ([]byte, error) {
	ctx = tracing.CaptureTraceContext(ctx)
	client := inits.HTTPClient(ctx, "")
	if client == nil {
		return nil, fmt.Errorf("get http client failed")
	}
	now := time.Now()
	// parent timeout on request
	if deadline, ok := ctx.Deadline(); ok {
		var parentCancel context.CancelFunc
		ctx, parentCancel = context.WithTimeout(ctx, deadline.Sub(now))
		defer parentCancel()
	}
	req := request.WithContext(ctx)
	r := httpclient.NewRequest().WithRequest(req)
	result, err := callbyCustomRequest(ctx, client, r)
	if err != nil {
		return nil, err
	}
	return result.RspBody, nil
}

// 对于以下方法中的service参数说明:
// 如果对应的server_client配置了app_name选项,则需要调用方保证service参数带上app_name前缀
// 如果没有配置,则保持原有逻辑,	service参数不用改动
func NewRequest(service, method, urlStr string, body io.Reader) (*http.Request, error) {
	if strings.Contains(urlStr, defaultIpportFormat) {
		sc, err := serviceClient(service, nil)
		if err != nil {
			return nil, err
		}
		host := inits.Default.Clusters.ChooseHost(context.TODO(), sc.Cluster.Name)
		if host != nil {
			urlStr = strings.Replace(urlStr, defaultIpportFormat, host.Address(), -1)
		} else {
			return nil, ErrClientLB
		}
	}
	return http.NewRequest(method, urlStr, body)
}

func HttpPut(ctx context.Context, service string, uri string, config RequestOptionIntercace) ([]byte, error) {
	return callRespByte(ctx, service, "PUT", uri, config, nil)
}

func HttpDelete(ctx context.Context, service string, uri string, config RequestOptionIntercace) ([]byte, error) {
	return callRespByte(ctx, service, "DELETE", uri, config, nil)
}

func HttpPostD(ctx context.Context, service string, uri string, config RequestOptionIntercace, body io.Reader) (HttpDResponse, error) {
	return callRespObj(ctx, service, "POST", uri, config, body)
}

func HttpGetD(ctx context.Context, service string, uri string, config RequestOptionIntercace) (HttpDResponse, error) {
	return callRespObj(ctx, service, "GET", uri, config, nil)
}

func HttpPost(ctx context.Context, service string, uri string, config RequestOptionIntercace, body io.Reader) ([]byte, error) {
	return callRespByte(ctx, service, "POST", uri, config, body)
}

func HttpGet(ctx context.Context, service string, uri string, config RequestOptionIntercace) ([]byte, error) {
	return callRespByte(ctx, service, "GET", uri, config, nil)
}

func callRespByte(ctx context.Context, service, method, uri string, config RequestOptionIntercace, body io.Reader) ([]byte, error) {
	rsp, err := callRespObj(ctx, service, method, uri, config, body)
	if err != nil {
		return nil, err
	}
	return rsp.RspBody, nil
}

func callRespObj(ctx context.Context, service, method, uri string, config RequestOptionIntercace, body io.Reader) (HttpDResponse, error) {
	client := inits.HTTPClient(ctx, service)
	if client == nil {
		return HttpDResponse{}, fmt.Errorf("httpclient with %s not found", service)
	}
	// only use http or https
	ro := &httpclient.RequestOption{}
	if len(service) > 0 {
		sc, err := serviceClient(service, config)
		if err != nil {
			return HttpDResponse{}, err
		}
		// request option
		ro.RequestTimeoutMS(sc.ReadTimeout).
			SlowTimeoutMS(sc.SlowTime)
	}

	// request body
	var bodyBytes []byte
	if body != nil {
		bodyBytes, _ = ioutil.ReadAll(body)
	}
	nbody := bytes.NewReader(bodyBytes)

	// new request
	req := httpclient.NewRequest(ctx).
		WithMethod(method).
		WithURL(uri).
		WithBody(nbody)

	// request header
	var attach map[string]string
	if config != nil {
		attach = config.GetMetricTags()
		for key, value := range config.GetHeaderMap() {
			keyLower := strings.ToLower(key)
			// trace相关的header不放入request header中
			if keyLower == jaeger.TraceContextHeaderName ||
				keyLower == jaeger.JaegerDebugHeader ||
				keyLower == jaeger.JaegerBaggageHeader ||
				strings.HasPrefix(keyLower, jaeger.TraceBaggageHeaderPrefix) {
				continue
			}
			req.AddHeader(key, value)
		}
	}

	ro.MetricTags(attach)
	req.WithOption(ro)

	result, err := callbyCustomRequest(ctx, client, req)
	result.ServiceNameUrl = service + uri
	return result, err
}

func callbyCustomRequest(ctx context.Context, client httpclient.Client, request *httpclient.Request) (HttpDResponse, error) {
	var address string
	result := HttpDResponse{
		SuccIp:   "",
		RspBody:  nil,
		Method:   "",
		FailIp:   nil,
		HttpCode: -1,
	}
	rsp, err := client.Call(request)
	select {
	case <-ctx.Done():
		err = ctx.Err()
	default:
		if err != nil {
			failIPs := make([]string, 0)
			failIPs = append(failIPs, address)
			result.FailIp = failIPs
		}
	}

	if rsp != nil {
		result.HttpCode = rsp.Code()
		result.RspBody = rsp.Bytes()
		result.Method = rsp.RawRequest().Method
		result.SuccIp = rsp.RawRequest().URL.Host
	}
	return result, err
}
