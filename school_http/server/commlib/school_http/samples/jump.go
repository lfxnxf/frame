package main

import "github.com/lfxnxf/frame/school_http/server/commlib/school_http"

func main() {
	school_http.Response.Jump(nil, school_http.JumpInfo{
		JumpType: school_http.JumpTypes.Dialog,
		Url:      "https://www.baidu.com",
		Title:    "开播安全",
		Content:  "尊敬的用户您好，根据相关政策要求实名认证后才可开播，点击“去认证”按钮进行认证",
		Confirm:  "去认证",
		Cancel:   "放弃",
	}, nil, nil)

}
