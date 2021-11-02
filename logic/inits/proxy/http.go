package proxy

import (
	"github.com/lfxnxf/frame/logic/inits"
	"github.com/lfxnxf/frame/logic/inits/http/client"
	"golang.org/x/net/context"
)

const (
	GETMethod             = "GET"
	POSTMethod            = "POST"
	_inkeApp              = "inke"
	_serviceNameSeperator = "."
	_jsonContentType      = "application/json; charset=utf-8"
)

type HTTP struct {
	name string
}

func InitHTTP(name string) *HTTP {
	return &HTTP{name}
}

func (h *HTTP) WithApp(name string) *HTTP {
	if name != _inkeApp {
		h.name = name + _serviceNameSeperator + h.name
	}
	return h
}

func (h *HTTP) Call(ctx context.Context, request *client.Request) (*client.Response, error) {
	return inits.HTTPClient(ctx, h.name).Call(request)
}

func (h *HTTP) JSONGet(ctx context.Context, uri string, ro *client.RequestOption, respObj interface{}) error {
	req := client.NewRequest(ctx).WithURL(uri).WithMethod(GETMethod).AddHeader("Content-Type", _jsonContentType)
	if ro != nil {
		req = req.WithOption(ro)
	}
	return h.jsonCall(ctx, req, respObj)
}

func (h *HTTP) JSONPost(ctx context.Context, uri string, ro *client.RequestOption, body, respObj interface{}) error {
	req := client.NewRequest(ctx).WithURL(uri).WithMethod(POSTMethod).WithStruct(body).AddHeader("Content-Type", _jsonContentType)
	if ro != nil {
		req.WithOption(ro)
	}
	return h.jsonCall(ctx, req, respObj)
}

func (h *HTTP) jsonCall(ctx context.Context, request *client.Request, respObj interface{}) error {
	resp, err := inits.HTTPClient(ctx, h.name).Call(request)
	if err != nil {
		return err
	}
	return resp.JSON(respObj)
}
