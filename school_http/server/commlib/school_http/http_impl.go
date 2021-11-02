package school_http

import (
	"github.com/lfxnxf/frame/logic/rpc-go"
	"go.uber.org/multierr"
	"net/url"
	"reflect"

	goquery "github.com/google/go-querystring/query"
)

func (r *requestImpl) queryUrlValues(query url.Values) *requestImpl {
	for k, vs := range query {
		r.query[k] = append(r.query[k], vs...)
	}
	return r
}

func (r *requestImpl) queryString(query string) *requestImpl {
	q, err := url.ParseQuery(query)
	if err != nil {
		r.err = multierr.Append(r.err, err)
		return r
	}
	r.queryUrlValues(q)
	return r
}

func (r *requestImpl) queryReflect(query interface{}) *requestImpl {
	val := reflect.ValueOf(query)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	switch val.Kind() {
	case reflect.Struct:
		q, err := goquery.Values(query)
		if err != nil {
			r.err = multierr.Append(r.err, err)
			return r
		}

		for k, vs := range q {
			r.query[k] = append(r.query[k], vs...)
		}
	}
	return r
}

func (r *requestImpl) initOption() {
	if r.option == nil {
		r.option = rpc.NewRequestOptional()
	}
}

