package school_http

import (
	"github.com/lfxnxf/frame/BackendPlatform/golang/ecode"
	"github.com/lfxnxf/frame/logic/inits/utils"
	"github.com/lfxnxf/frame/school_http/server/commlib/school_errors"
)

type jumpType int

var JumpTypes = struct {
	Dialog  jumpType // 弹窗确认
	Jump    jumpType // 直接跳转
	H5Popup jumpType // H5弹窗
}{
	Dialog:  0,
	Jump:    1,
	H5Popup: 2,
}

type JumpInfo struct {
	JumpType    jumpType `json:"jt"`   // 触发形式。0-弹窗；1-直接跳转
	Url         string   `json:"url"`  // 跳转地址
	Title       string   `json:"tt"`   // 弹窗标题
	Content     string   `json:"ctt"`  // 弹窗展示内容
	Confirm     string   `json:"cfm"`  // 确认按钮文案
	Cancel      string   `json:"ccl"`  // 取消按钮文案
	UrlNeedAtom bool     `json:"una"`  // 是否需要在Url后附上原子参数
}

// 女娲API层，通用结构封装
type nvwaWrapResp struct {
	utils.WrapResp
	Jump JumpInfo `json:"jump"`
}

type response struct{}

// Deprecated: 废弃，请使用Responses
var Response response

var Responses response

func (h response) Jump(c *nwcontext, j JumpInfo, data interface{}, err error) {
	c.Response.WriteHeader(c.Response.Status())
	e := ecode.Cause(err)
	c.SetBusiCode(int32(e.Code()))

	resp := nvwaWrapResp{
		WrapResp: utils.WrapResp{
			Code: e.Code(),
			Msg:  e.Message(),
			Data: data,
		},
		Jump: j,
	}
	if resp.Code != school_errors.Codes.Success.Code() {
		resp.Code = school_errors.Codes.Redirect.Code()
		resp.Msg = e.Message()
	}
	c.Response.WriteJSON(resp)
}
func (h response) JumpAbort(c *nwcontext, j JumpInfo, data interface{}, err error) {
	h.Jump(c, j, data, err)
	c.Abort()
}
func (h response) Json(c *nwcontext, data interface{}, err error) {
	if HasSession(c.Ctx) {
		v, ok := Session(c.Ctx).Get(_jumpInfoKey)
		if ok && v != nil {
			if jmpInfo, ok := v.(JumpInfo); ok {
				h.Jump(c, jmpInfo, data, err)
				return
			}
		}
	}
	c.JSON(data, err)
}
func (h response) JsonAbort(c *nwcontext, data interface{}, err error) {
	h.Json(c, data, err)
	c.Abort()
}
