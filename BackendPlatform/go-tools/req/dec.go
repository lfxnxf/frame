package req

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	log "github.com/lfxnxf/frame/BackendPlatform/golang/logging"
	"github.com/gorilla/schema"
	"github.com/pkg/errors"
	"go.uber.org/multierr"
	"gopkg.in/go-playground/validator.v9"
)

var (
	URIDecoder = schema.NewDecoder()
	V          = validator.New()
)

func init() {
	URIDecoder.IgnoreUnknownKeys(true)
}
func ReqDecode(request *http.Request, req interface{}, atom ...interface{}) error {
	method := request.Method
	reqURL := request.URL.String()
	switch method {
	case "GET":
		dErr := URIDecoder.Decode(req, request.URL.Query())
		vErr := V.Struct(req)
		if dErr != nil || vErr != nil {
			log.Warnf("http parse reqUrl failed:|reqUrl:%v|err:%v,%v|", reqURL, dErr, vErr)
			return multierr.Combine(vErr, dErr)
		}
	case "POST":
		body, err := ioutil.ReadAll(request.Body)
		uErr := json.Unmarshal(body, &req)
		vErr := V.Struct(req)
		if err != nil || uErr != nil || vErr != nil {
			log.Warnf("http parse body failed:|reqUrl:%v|body:%s|err:%v,%v,%v|", reqURL, body, err, uErr, vErr)
			return multierr.Combine(vErr, uErr)
		}
	default:
		log.Warnf("http method not support:|%v|%v|", method, reqURL)
		return errors.New("param error")
	}
	if len(atom) > 0 {
		dErr := URIDecoder.Decode(atom[0], request.URL.Query())
		vErr := V.Struct(atom[0])
		if dErr != nil || vErr != nil {
			log.Warnf("http parse atom failed:|reqUrl:%v|err:%v,%v|", reqURL, dErr, vErr)
			return multierr.Combine(vErr, dErr)
		}
	}
	return nil
}
