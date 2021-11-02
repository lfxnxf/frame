## 添加错误码

```
错误码添加分为中台侧和业务侧
```


 - 中台侧，错误码分别在对应的 xxx_errors.go 文件中添加，错误码范围参考：codes_apply.txt，配置完成后，合入common_errors.go，合并方法参考common_errors.go
 
示例第1步：增值付费中心错误码
 ```go
// 工程文件 commlib/nvwa_errors/commerce_errors.go
package nvwa_errors

//
// 增值付费中心错误码
// commerce:1200-1299
//
type commerceCode struct {
	BagItemNumError Code // 道具余额不足

	// 送礼
	GiftOffline  Code // 礼物已下线
	GiftNotExist Code // 礼物不存在
}

var _commerceCodes = commerceCode{
	BagItemNumError: genError(1200, "道具不足").SetToast("剩余数量不足"),

	GiftOffline:  genError(1220, "gift offline").SetToast("礼物已下架，请选择其他礼物"),
	GiftNotExist: genError(1221, "git not exist").SetToast("礼物不存在，请打开礼物墙选择您中意的礼物送礼"),
}
```

示例第2步：合入common_errors.go


```go
// 工程文件 commlib/nvwa_errors/common_errors.go
package code
 
 
var Codes = struct {
   ClientError        Code // 请求参数错误（调用端参数有问题）
   
   // 增值付费中心错误码
   commerceCode
}{
   ClientError:        genError(499, "请求参数错误").SetToast("网络异常，请重试"),
   
   commerceCode: _commerceCodes,
}
```


 - 业务侧，错误码可以使用 nvwa_errors.AddError 方式添加，错误码要 >= 10002，使用示例如下
 ```go
 // 工程文件 api/code/code.go
 package code

 import (
	"github.com/lfxnxf/frame/school_http/server/commlib/school_errors"
 )

 var (
    ProfileIncomplete = nvwa_errors.AddError(10002, "资料缺失").SetToast("您的资料不完整，请完善资料后再次尝试")
 )
 ```

## 错误码使用
 - 示例 server/http/handler.go
 ```go
 func ping(c *httpserver.Context) {

     // // 服务内部调用，返回错误信息
     // 示例1
     c.JSON(nil, code.ProfileIncomplete)
     // 示例2
     c.JSONAbort(nil, code.ProfileIncomplete)

     // 客户端调用，返回用户提示信息
     // 示例1
     c.JSON(nil, code.ProfileIncomplete.Toast())
     // 示例2
     c.JSONAbort(nil, code.ProfileIncomplete.Toast())
     
     // 透传的错误码服务端调用返回
     c.JSONAbort(nil, err)
     
     // 透传的错误码客户端调用接口，使用nvwa_errors.Toast返回
     // 示例1
     c.JSON(nil, nvwa_errors.Toast(err))
     // 示例2
     c.JSONAbort(nil, nvwa_errors.Toast(err))
 }
 ```
 
 
 ## 配套套餐
  - 为了更好的使用体验，建议使用中台的nvwa_http库进行请求数据的解析和发起http请求  
  
GET 请求解析
```go
var req model.XXXRequest
// 示例1：同步得到原子信息
atom, err := nvwa_http.Requests.Query(c.Ctx, c.Request).Parse(&req).Atom()
// 示例2：仅解析
err := nvwa_http.Requests.Query(c.Ctx, c.Request).Parse(&req).Error()
if err != nil {
	log.Error("get atom error")
	c.JSONAbort(nil, err)
	return
}
```
POST 请求解析
```go
var req model.XXXRequest
if err := nvwa_http.Requests.Body(c.Ctx, c.Request).ParseJson(&req).Error(); err != nil {
	log.Error("parse error", "err", err.Error())
	c.JSONAbort(nil, err)
	return
}
```

HTTP 请求其他服务接口
```go
req := coinTransferReq{
	AppKey:  appKey,
	Source:  source,
	OrderID: orderID,
	Uid:     from.Uid,
	From:    from,
	To:      toList,
	Extra:   extra,
}
// 示例1，无返回数据使用，err值可以直接填入c.JSONAbort(nil, err)，GET调用将Post改为Get即可
if err := nvwa_http.NewReq(ctx, m.coinClient).Post(_coinTransferUri).WithBodyJson(req).Response().ParseEmpty(); err != nil {
	logging.For(ctx, "func", "CoinTransfer", "appKey", appKey, "source", source, "orderID", orderID, "from", from, "toList", toList, "extra", extra).Errorw("request error", "err", err)
	return err
}
// 示例2，有返回数据，err值可以直接填入c.JSONAbort(nil, err)，GET调用将Post改为Get即可
data := TransferData{}
if err := nvwa_http.NewReq(ctx, m.coinClient).Post(_coinTransferUri).WithBodyJson(req).Response().ParseDataJson(&data); err != nil {
	logging.For(ctx, "func", "CoinTransfer", "appKey", appKey, "source", source, "orderID", orderID, "from", from, "toList", toList, "extra", extra).Errorw("request error", "err", err)
	return err
}
```
