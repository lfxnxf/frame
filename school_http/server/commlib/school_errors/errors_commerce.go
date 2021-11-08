package school_errors

//
// 增值付费中心错误码
// commerce:1200-1299
//
type commerceCode struct {
	BagItemNumError Code // 道具余额不足

	// 送礼
	GiftOffline  Code // 礼物已下线
	GiftNotExist Code // 礼物不存在

	GiftWebItemNotExist Code // 物品不存在
	GiftWebItemInUse    Code // 物品使用中
	GiftWebAuthError    Code // 无权限或权限校验失败

	GiftReturnValueNullError Code // 业务接口返回数据为空错误

	UIDANDOrderIDNotMatchError Code //订单号和用户不匹配
	InvalidOrderError          Code //无效订单
	OrderStatusUsedError       Code //订单状态已被使用
	InsufficientQuantityError  Code //数量不足
	GoodsExpiredError          Code //物品已过期
}

var _commerceCodes = commerceCode{
	BagItemNumError: genError(1200, "道具不足").SetToast("剩余数量不足"),

	GiftOffline:  genError(1220, "gift offline").SetToast("礼物已下架，请选择其他礼物"),
	GiftNotExist: genError(1221, "git not exist").SetToast("礼物不存在，请打开礼物墙选择您中意的礼物送礼"),

	GiftWebItemNotExist: genError(1222, "item not exist").SetToast("物品不存在，请检查"),
	GiftWebItemInUse:    genError(1223, "item in use").SetToast("物品正在使用中, 无法进行操作"),

	GiftWebAuthError: genError(1224, "auth error").SetToast("获取权限失败"),

	GiftReturnValueNullError: genError(1225, "return data is null").SetToast("接口返回数据为空，请检查"),

	UIDANDOrderIDNotMatchError: genError(1230, "uid and order_id do not match").SetToast("UID和OrderID不匹配"),
	InvalidOrderError:          genError(1231, "invalid order").SetToast("无效订单"),
	OrderStatusUsedError:       genError(1232, "order status has been used").SetToast("订单状态已被使用"),
	InsufficientQuantityError:  genError(1233, "num insufficient quantity").SetToast("物品数量不足"),
	GoodsExpiredError:          genError(1234, "goods expired").SetToast("物品已过期"),
}
