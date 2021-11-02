package main

import (
	"github.com/lfxnxf/frame/school_http/server/commlib/school_errors"
	"github.com/lfxnxf/frame/school_http/server/commlib/school_http"
)

func main() {
	//if err := nvwa_http.NewReq(nil, nil).Response().ParseJson(nil); err != nil {

	//}
	//nvwa_http.NewReq(nil, nil).
	//nvwa_http.NewReq(nil, nil).Get("").
	err := school_http.NewReq(nil, nil).Post("").WithTimeout(1000, 2, 200).Response().ParseDataJson(nil)
	if err != nil {
		school_errors.Toast(err)
	}
}
