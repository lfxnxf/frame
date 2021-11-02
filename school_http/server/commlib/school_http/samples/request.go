package main

import "github.com/lfxnxf/frame/school_http/server/commlib/school_http"

func main() {
	// 解析自定义数据并获取Atom
	atom, err := school_http.Request.Query(nil, nil).Parse(nil).Atom()
	if atom == nil || err != nil {
	}

	// 仅解析自定义数据
	err = school_http.Request.Query(nil, nil).Parse(nil).Error()
	if err != nil {
	}

	// 仅解析原子信息
	atom, err := school_http.Request.Query(nil, nil).Atom()
	if atom == nil || err != nil {
	}

	// 解析body数据
	err = school_http.Request.Body(nil, nil).ParseJson(nil).Error()
	if err != nil {

	}
}
