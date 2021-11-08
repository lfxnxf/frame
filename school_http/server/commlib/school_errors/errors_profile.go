package school_errors

//用户资料错误码
//1300-1399

type profileCode struct {
	UpdateProfileSex      Code
	InvalidUID            Code
	InvalidVersion        Code
	AuditNotExist         Code
	AuditConsistent       Code
	InvalidType           Code
	KeyNotExist           Code
	AuditArrayNotExist    Code
	FollowLimit           Code
	HaveSetGender         Code
	SensitiveWord         Code
	ProfileArrayLenLimit  Code
	FollowSelfErr         Code
	RepeatedPhotoErr      Code
	ProfileLenLimit       Code
	ProfileUpdateFrequent Code
}

var _profileCodes = profileCode{
	UpdateProfileSex:      genError(100001, "请设置性别").SetToast("请设置性别"),
	InvalidUID:            genError(1300, "invalid uid").SetToast("非法uid"),
	InvalidVersion:        genError(1301, "invalid version").SetToast("非法version"),
	AuditNotExist:         genError(1302, "audit not exist").SetToast("当前字段不支持审核"),
	AuditConsistent:       genError(1303, "audit consistent").SetToast("审核字段数据不一致"),
	InvalidType:           genError(1304, "invalid type").SetToast("无效的数据类型"),
	KeyNotExist:           genError(1305, "key not exist").SetToast("资料字段不存在"),
	AuditArrayNotExist:    genError(1306, "audit array not exist").SetToast("资料数组数据不存在"),
	FollowLimit:           genError(1307, "follow limit").SetToast("用户关注个数超限制"),
	HaveSetGender:         genError(1308, "Have set gender").SetToast("性别已设置"),
	SensitiveWord:         genError(1309, "设置信息含有敏感词汇").SetToast("设置信息含有敏感词汇"),
	ProfileArrayLenLimit:  genError(1310, "array max len ").SetToast("资料数组长度限制"),
	FollowSelfErr:         genError(1311, "can not follow self").SetToast("不能关注自己"),
	RepeatedPhotoErr:      genError(1312, "photo can not repeat").SetToast("相册里不允许有重复"),
	ProfileLenLimit:       genError(1313, "value max len limit").SetToast("资料长度限制"),
	ProfileUpdateFrequent: genError(1314, "profile update frequently").SetToast("更新资料过于频繁"),
}
