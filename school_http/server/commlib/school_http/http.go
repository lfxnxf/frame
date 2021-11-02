package school_http

import (
	"bytes"
	"context"
	"fmt"
	"github.com/lfxnxf/frame/BackendPlatform/golang/logging"
	"github.com/lfxnxf/frame/logic/inits"
	"github.com/lfxnxf/frame/logic/inits/config/encoder/json"
	"github.com/lfxnxf/frame/logic/inits/http/client"
	"github.com/lfxnxf/frame/logic/inits/proxy"
	"github.com/lfxnxf/frame/logic/inits/utils"
	"github.com/lfxnxf/frame/logic/rpc-go"
	"github.com/lfxnxf/frame/school_http/server/commlib/school_errors"
	"go.uber.org/multierr"
	"go.uber.org/zap"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	_jumpInfoKey  = "_NVWA_HTTP_JUMP_INFO_"
	_appKeyHeader = "uberctx-_namespace_appkey_"
)

type requestImpl struct {
	ctx  context.Context
	http *proxy.HTTP

	method string
	uri    string

	option *rpc.RequestOption

	query      url.Values
	reqBodyRaw interface{}
	body       io.Reader

	req *client.Request
	rsp *client.Response

	err error
}
type requestDetailImpl struct {
	impl   *requestImpl
	header bool
}

// New return *Req, don't reuse Req
func NewReq(ctx context.Context, http *proxy.HTTP) *requestImpl {
	return &requestImpl{
		ctx:   ctx,
		http:  http,
		query: url.Values{},
		req:   client.NewRequest(ctx),
	}
}

// Get execute a HTTP GET request
func (r *requestImpl) Get(uri string) *requestDetailImpl {
	if r.err != nil {
		return &requestDetailImpl{
			impl: r,
		}
	}
	r.method = http.MethodGet
	r.uri = uri
	return &requestDetailImpl{
		impl: r,
	}
}

// Post execute a HTTP POST request
func (r *requestImpl) Post(uri string) *requestDetailImpl {
	if r.err != nil {
		return &requestDetailImpl{
			impl: r,
		}
	}
	r.method = http.MethodPost
	r.uri = uri
	return &requestDetailImpl{
		impl: r,
	}
}

// Deprecated: 废弃，请先使用Get或Post
func (r *requestImpl) AddHeader(k string, v interface{}) *requestDetailImpl {
	impl := &requestDetailImpl{
		impl: r,
	}
	return impl.WithHeader(k, v)
}

// Deprecated: 废弃，请使用 WithHeader
func (r *requestDetailImpl) AddHeader(k string, v interface{}) *requestDetailImpl {
	return r.WithHeader(k, v)
}
func (r *requestDetailImpl) WithHeader(k string, v interface{}) *requestDetailImpl {
	//r.initOption()
	//r.option.SetHeader(k, fmt.Sprint(v))
	if k == _appKeyHeader {
		r.header = true
	}
	r.impl.req = r.impl.req.AddHeader(k, fmt.Sprint(v))
	return r
}

// Deprecated: 废弃，请先使用Get或Post
func (r *requestImpl) SetHeader(header map[string]interface{}) *requestDetailImpl {
	impl := &requestDetailImpl{
		impl: r,
	}
	return impl.WithHeaderMap(header)
}

// Deprecated: 废弃，请使用 WithHeaderMap
func (r *requestDetailImpl) SetHeader(header map[string]interface{}) *requestDetailImpl {
	return r.WithHeaderMap(header)
}
func (r *requestDetailImpl) WithHeaderMap(header map[string]interface{}) *requestDetailImpl {
	//r.initOption()
	for k, v := range header {
		//r.option.SetHeader(k, fmt.Sprint(v))
		if k == _appKeyHeader {
			r.header = true
		}
		r.impl.req = r.impl.req.AddHeader(k, fmt.Sprint(v))
	}
	return r
}
func (r *requestDetailImpl) WithHeaders(keyAndValues ...interface{}) *requestDetailImpl {
	//r.initOption()
	l := len(keyAndValues) - 1
	//for k, v := range header {
	for i := 0; i < l; i += 2 {
		//r.option.SetHeader(k, fmt.Sprint(v))
		k := fmt.Sprint(keyAndValues[i])
		if k == _appKeyHeader {
			r.header = true
		}
		r.impl.req = r.impl.req.AddHeader(k, fmt.Sprint(keyAndValues[i+1]))
	}
	if (l+1)%2 == 1 {
		logging.For(r.impl.ctx, zap.String("func", "nvwa_http.NewReq().XXX().WithHeaders")).Warnw("the keys are not aligned")
		k := fmt.Sprint(keyAndValues[l])
		if k == _appKeyHeader {
			r.header = true
		}
		r.impl.req = r.impl.req.AddHeader(k, "")
	}
	return r
}

// Param add query param
// Deprecated: 废弃，请先使用Get或Post
func (r *requestImpl) Param(k string, v interface{}) *requestDetailImpl {
	impl := &requestDetailImpl{
		impl: r,
	}
	return impl.WithQueryParam(k, v)
}

// Deprecated: 废弃，请使用 WithQueryParam
func (r *requestDetailImpl) Param(k string, v interface{}) *requestDetailImpl {
	return r.WithQueryParam(k, v)
}
func (r *requestDetailImpl) WithQueryParam(k string, v interface{}) *requestDetailImpl {
	if r.impl.err != nil {
		return r
	}
	r.impl.query.Add(k, fmt.Sprint(v))
	return r
}
func (r *requestDetailImpl) WithQueryParams(keyAndValues ...interface{}) *requestDetailImpl {
	if r.impl.err != nil {
		return r
	}
	l := len(keyAndValues) - 1
	for i := 0; i < l; i += 2 {
		r.impl.query.Add(fmt.Sprint(keyAndValues[i]), fmt.Sprint(keyAndValues[i+1]))
	}
	if (l+1)%2 == 1 {
		logging.For(r.impl.ctx, zap.String("func", "nvwa_http.NewReq().XXX().WithQueryParams")).Warnw("the keys are not aligned")
		r.impl.query.Add(fmt.Sprint(keyAndValues[l]), "")
	}
	return r
}

// Query add query param, support:
// Query("k1=v1&k2=v2")
// Query(url.Values{})
// Query(map[string] string{})
// Query(map[string] interface{})
// Query(struct{}) using url tag, reference: https://github.com/google/go-querystring
// Deprecated: 废弃，请先使用Get或Post
func (r *requestImpl) Query(query interface{}) *requestDetailImpl {
	impl := &requestDetailImpl{
		impl: r,
	}
	return impl.WithQuery(query)
}

// Deprecated: 废弃，请使用 WithQuery
func (r *requestDetailImpl) Query(query interface{}) *requestDetailImpl {
	return r.WithQuery(query)
}

// WithQuery add query param, support:
// WithQuery("k1=v1&k2=v2")
// WithQuery(url.Values{})
// WithQuery(map[string] string{})
// WithQuery(map[string] interface{})
// WithQuery(struct{}) using url tag, reference: https://github.com/google/go-querystring
func (r *requestDetailImpl) WithQuery(query interface{}) *requestDetailImpl {
	switch query := query.(type) {
	case string:
		r.impl.queryString(query)
	case []byte:
		r.impl.queryString(string(query))
	case url.Values:
		r.impl.queryUrlValues(query)
	case map[string]string:
		for k, v := range query {
			r.impl.query.Add(k, v)
		}
	case map[string]interface{}:
		for k, v := range query {
			r.impl.query.Add(k, fmt.Sprint(v))
		}
		//slow path
	default:
		r.impl.queryReflect(query)
	}
	return r
}
func (r *requestDetailImpl) initOptions() *rpc.RequestOption {
	if r.impl.option == nil {
		r.impl.option = rpc.NewRequestOptional()
	}
	return r.impl.option
}

// 设置超时时间。
// timeout：超时时间，单位毫秒，如不设置，则默认使用配置文件配置，如配置文件亦没有，则无限等待，直至网络异常断开。
// retryTimes：重试次数，如不设置，则默认使用配置文件配置，如配置文件亦没有，则默认0次重试
// slowTime：慢调用时间，单位毫秒，如不设置，则默认使用配置文件配置，如配置文件亦没有，则默认为0，不启用慢时间。
func (r *requestDetailImpl) WithTimeout(timeout int, retryTimes int, slowTime int) *requestDetailImpl {
	opt := r.initOptions()
	opt.SetTimeOut(timeout)
	opt.SetRetryTimes(retryTimes)
	opt.SetSlowTime(slowTime)
	return r
}

// BodyJson set request body, support ioReader, string, []byte, otherwise use json.Marshal
// Deprecated: 废弃，请先使用Get或Post
func (r *requestImpl) BodyJson(model interface{}) *requestDetailImpl {
	impl := &requestDetailImpl{
		impl: r,
	}
	return impl.WithBodyJson(model)
}

// Deprecated: 废弃，请使用 WithBodyJson
func (r *requestDetailImpl) BodyJson(model interface{}) *requestDetailImpl {
	return r.WithBodyJson(model)
}

// Deprecated: 废弃，请使用 WithBodyJson
func (r *requestDetailImpl) WithBodyJson(model interface{}) *requestDetailImpl {
	return r.WithBody(model)
}

// body：如为struct，则将自动解析成json，否则直接流式放入body
func (r *requestDetailImpl) WithBody(body interface{}) *requestDetailImpl {
	r.impl.reqBodyRaw = body
	switch v := body.(type) {
	case io.Reader:
		r.impl.body = v
	case []byte:
		r.impl.body = bytes.NewReader(v)
	case string:
		r.impl.body = strings.NewReader(v)
	default:
		buf, err := json.NewEncoder().Encode(body)
		if err != nil {
			r.impl.err = multierr.Append(r.impl.err, err)
			return r
		}
		r.impl.body = bytes.NewReader(buf)
	}
	return r
}

func (r *requestImpl) toBytes() (*client.Response, error) {
	if len(r.query) != 0 {
		if strings.IndexByte(r.uri, '?') == -1 {
			r.uri += "?" + r.query.Encode()
		} else {
			r.uri += "&" + r.query.Encode()
		}
	}

	//v, ok := inits.FromContext(r.ctx)
	//if !ok || v.Namespace == "" {
	//	//return &model.RpcError{LocalErr: fmt.Errorf("namespace must not be empty")}
	//	return nil, fmt.Errorf("can't find namespace from context")
	//}

	//logging.Debugw("request", "uri", r.uri, "method", r.method)
	req := r.req.WithPath(r.uri).WithMethod(r.method).WithBody(r.body)
	rsp, err := r.http.Call(r.ctx, req)
	if err != nil {
		return nil, err
	}
	return rsp, nil
}

//
func (r *requestDetailImpl) Response() *responseImpl {
	if r.impl.err != nil {
		return &responseImpl{
			req: r.impl,
		}
	}
	if !r.header {
		d, ok := inits.FromContext(r.impl.ctx)
		if ok {
			r.WithHeader("uberctx-_namespace_appkey_", d.Namespace)
		}
	}
	rsp, err := r.impl.toBytes()
	if err != nil {
		r.impl.err = multierr.Append(r.impl.err, err)
		return &responseImpl{
			req: r.impl,
		}
	}
	r.impl.rsp = rsp
	return &responseImpl{
		req: r.impl,
	}
}

type responseImpl struct {
	req *requestImpl
}

// Deprecated: 废弃，请使用ParseDataJson
func (r *responseImpl) ParseJson(dataModel interface{}) error {
	return r.ParseDataJson(dataModel)
}

func (r *responseImpl) ParseEmpty() error {
	return r.ParseDataJson(nil)
}

// Deprecated: 废弃，请使用ParseDataJson
func (r *responseImpl) ParseData(dataModel interface{}) error {
	return r.ParseDataJson(dataModel)
}

// 解析数据
// json结构，解析整个response
func (r *responseImpl) Json(resp interface{}) error {
	if r.req.err != nil {
		return r.req.err
	}
	err := r.req.rsp.JSON(&resp)
	if err != nil {
		r.req.err = multierr.Append(r.req.err, err)
		return r.req.err
	}
	return nil
}

// 解析response的data区数据，并根据错误自动返回对应error。反解error，可以使用nvwa_errors.DMError函数
func (r *responseImpl) ParseDataJson(dataModel interface{}) error {
	if r.req.err != nil {
		return r.req.err
	}

	resp := nvwaWrapResp{
		WrapResp: utils.WrapResp{
			Data: dataModel,
		},
	}

	err := r.req.rsp.JSON(&resp)
	if err != nil {
		r.req.err = multierr.Append(r.req.err, err)
		return r.req.err
	}

	if resp.Code != 0 {
		// 重定向错误码，则存储，以便透传到客户端
		if resp.Code == school_errors.Codes.Redirect.Code() {
			//context.WithValue(r.req.ctx, _jumpInfoKey, resp.Jump)
			timeout := time.Second * 3
			if r.req.option != nil && r.req.option.GetTimeOut() > 0 {
				timeout = time.Millisecond * time.Duration(r.req.option.GetTimeOut())
			}
			if timeout <= 0 { // 没有设置http timeout，则默认10秒
				timeout = time.Second * 3
			} else if timeout > time.Second*10 { // http timeout 超过10秒，则默认保留10秒
				timeout = time.Second * 10
			}
			Session(r.req.ctx).WithTimeout(timeout).Add(_jumpInfoKey, resp.Jump)
		}
		//return errors.Errorf("server response business error code")
		r.req.err = multierr.Append(r.req.err, school_errors.Get(resp.Code))
		// 注释，通过NewTmpError实现，避免下游服务的错误码不在错误码仓库内
		//return nvwa_errors.Get(resp.Code)
		return school_errors.NewTmpError(resp.Code, resp.Msg)
	}
	return nil
}
