package http

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"

	log "github.com/lfxnxf/frame/BackendPlatform/golang/logging"
	"github.com/lfxnxf/frame/logic/rpc-go"
	"github.com/bitly/go-simplejson"
)

// PostMustSucc send http post request
func PostMustSucc(ctx context.Context, serviceName, url string, body interface{}, result interface{}) error {
	b, err := json.Marshal(body)
	if err != nil {
		return err
	}
	buf := bytes.NewBuffer(b)
	resp, err := rpc.HttpPost(ctx, serviceName, url, nil, buf)
	if err != nil {
		log.Errorf("http post request error,serviceName(%v),url(%v),body(%v),resp(%v),err(%v)", serviceName, url, body, string(resp), err)
		return err
	}
	dmErr, err := simplejson.NewJson(resp)
	if err != nil || dmErr.Get("dm_error").MustInt() != 0 {
		log.Errorf("http post dm_error != 0 serviceName(%v),url(%v),body(%v),resp(%v),err(%v)", serviceName, url, body, string(resp), err)
		return errors.New("error:dm_error != 0")
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		log.Errorf("http post unmarshal error,serviceName(%v),url(%v),body(%v),resp(%v),err(%v)", serviceName, url, body, string(resp), err)
		return err
	}
	return nil
}

// GetMustSucc send http get request
func GetMustSucc(ctx context.Context, serviceName, url string, result interface{}) error {
	resp, err := rpc.HttpGet(ctx, serviceName, url, nil)
	if err != nil {
		log.Errorf("http post request error,serviceName(%v),url(%v),resp(%v),err(%v)", serviceName, url, string(resp), err)
		return err
	}
	dmErr, err := simplejson.NewJson(resp)
	if err != nil || dmErr.Get("dm_error").MustInt() != 0 {
		log.Errorf("http post dm_error != 0 error,serviceName(%v),url(%v),resp(%v),err(%v)", serviceName, url, string(resp), err)
		return errors.New("error:dm_error != 0")
	}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		log.Errorf("http post unmarshal error,serviceName(%v),url(%v),resp(%v),err(%v)", serviceName, url, string(resp), err)
		return err
	}
	return nil
}
