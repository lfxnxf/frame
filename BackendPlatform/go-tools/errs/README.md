# 通用的错误定义
逻辑返回可以返回DMError类型，实现了error接口


### 示例1
```go
//handler处理请求
type ActionListHandler struct {
	sv *livemng.MngSv
}

func (h *ActionListHandler) Serve(ctx context.Context, request *http.Request) (interface{}, int) {
	var r define.ActionListReq
	var atom lib.Atom
	//解析请求并验证
	if err := req.ReqDecode(request, &r, &atom); err != nil {
		return errs.Resp(errs.ErrParam)
	}
	//逻辑处理
	data, err := h.sv.GetActionList(ctx, r, atom)
	//返回错误或者数据
	if err != nil {
		return errs.Resp(err)
	}
	return errs.Resp(err, data)
}

//逻辑部分
var ErrGetListFail=errs.New(1000, "获取权限列表错误")

type MngSv struct {
	dao    *cache.MngCache
}

func (s *MngSv) GetActionList(c context.Context, req define.ActionListReq, atom lib.Atom) (interface{}, error) {
	if err := s.dao.GetList(); err != nil {
		return nil, ErrGetListFail
	}
	return nil, nil
}

```

### 示例2
```go
type ActionListReq struct {
	LiveID   string `json:"live_id"`
	LiveUID  int64  `schema:"live_uid" json:"live_uid" validate:"gt=0"`
	ID       int64  `schema:"id" json:"id" validate:"gt=0"`
	LiveType string `json:"live_type"`
}
type Atom struct {
	Uid       int64  `schema:"uid" json:"uid" validate:"required,gt=0"`
	Gender    int    `schema:"gender" json:"gender"`
	Lc        string `schema:"lc"  json:"lc"`
	Cc        string `schema:"cc" json:"cc"`
	Cv        string `schema:"cv" json:"cv"`
	Ua        string `schema:"ua" json:"ua"`
	Conn      string `schema:"conn" json:"conn"`
	Devi      string `schema:"devi" json:"devi"`
	Idfv      string `schema:"idfv" json:"idfv"`
	Idfa      string `schema:"idfa" json:"idfa"`
	Proto     string `schema:"proto" json:"proto"`
	Osversion string `schema:"osversion" json:"osversion"`
	Logid     string `schema:"logid" json:"logid"`
	Smid      string `schema:"smid" json:"smid"`
	Xrealip   string `schema:"xrealip" json:"xrealip"`
	SrcType   string `schema:"src_type" json:"src_type"`
}

func (h *ActionListHandler) Serve(ctx context.Context, request *http.Request) (interface{}, int) {
	var r ActionListReq
	var atom Atom
	if err := req.ReqDecode(request, &r, &atom); err != nil {
		return errs.Resp(errs.ErrParam)
	}
	// logic
	data, err := h.sv.GetActionList(ctx, r, atom)
	if err != nil {
		return errs.Resp(err)
	}
	return errs.Resp(err, data)
}
```