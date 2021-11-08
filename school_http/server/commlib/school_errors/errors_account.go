package school_errors

//
// 账号相关错误码
// 600~699
//

type accountCode struct {
	SessionEmpty   Code // session id 为空，接口需要用户登录
	SessionExpired Code // session id 过期

	QQ_Oauth_Error       Code
	Apple_Oauth_Error    Code
	Weibo_Oauth_Error    Code
	Weixin_Oauth_Error   Code
	Baidu_Oauth_Error    Code
	Google_Oauth_Error   Code
	Facebook_Oauth_Error Code

	PhoneError            Code //输入手机号有误
	PhoneCodeError        Code //验证码错误
	PhoneFrequentError    Code //今日验证码次数超出限制
	PhoneChangeLimit      Code //一个月只能换绑一次手机号
	PhoneNotRightError    Code //该手机号与绑定手机号不一致
	PhoneRisk             Code //暂不支持该手机号,请联系客服
	PhoneCodeInvalid      Code //验证码已失效
	PhoneAlreadyBindError Code //该手机号已绑定
	PhoneCodeAlreadyUse   Code
	VisitorAuthError      Code
	WeiXinHaveBind        Code //该微信已绑定其他账号，暂时无法绑定
	PhoneHaveRegin        Code
	RsaPhoneError         Code
	PhoneLoginCodeError   Code // 手机输入验证码超过次数错误
	RiskCheckError        Code
	NeedWashWhite         Code

	UidAlreadyBindError Code
	Change_Phone_Step1  Code
	Change_Phone_Step2  Code
	Change_Phone_Step3  Code

	InvalidPhone Code

	LoginDevLimit  Code
	LoginDevLogOut Code

	JIGUANGAuthErr Code

	UserNotBindPhone Code

	EmailError              Code //邮箱格式错误
	EmailCodeError          Code //邮箱验证码错误
	EmailCodeCntError       Code //邮箱验证码超限制
	EmailFailErr            Code //邮箱登录失败 检查邮箱服务配置
	SendCodeNeedSliderError Code //code发送需要滑块
	SliderCheckError        Code //滑块验证失败
}

var _accountCodes = accountCode{
	SessionEmpty:            genError(603, "User is not logged in").SetToast("请登录后，访问该功能"),
	SessionExpired:          genError(604, "账号过期，请重新登录").SetToast("账号过期，请重新登录"),
	QQ_Oauth_Error:          genError(605, "QQ认证失败").SetToast("QQ 认证失败"),
	Weibo_Oauth_Error:       genError(606, "微博认证失败").SetToast("微博认证失败"),
	Weixin_Oauth_Error:      genError(607, "微信认证失败").SetToast("微信认证失败"),
	Baidu_Oauth_Error:       genError(608, "百度认证失败").SetToast("百度认证失败"),
	PhoneError:              genError(609, "输入手机号有误").SetToast("输入手机号有误"),
	PhoneCodeError:          genError(610, "验证码错误").SetToast("验证码错误"),
	PhoneFrequentError:      genError(611, "今日验证码次数超出限制").SetToast("今日验证码次数超出限制"),
	PhoneChangeLimit:        genError(612, "30天内只能换绑一次").SetToast("30天内只能换绑一次"),
	PhoneRisk:               genError(613, "手机号与绑定手机号不一致").SetToast("手机号与绑定手机号不一致"),
	PhoneCodeInvalid:        genError(614, "验证码已过期").SetToast("验证码已过期"),
	PhoneAlreadyBindError:   genError(615, "该手机号已绑定").SetToast("该手机号已绑定"),
	PhoneCodeAlreadyUse:     genError(616, "验证码已使用").SetToast("验证码已被使用"),
	UidAlreadyBindError:     genError(617, "已经绑定手机号").SetToast("已经绑定手机号"),
	Change_Phone_Step1:      genError(618, "换绑前，请使用原手机号发送验证码").SetToast("换绑前，请先用原手机号发验证码"),
	Change_Phone_Step2:      genError(619, "换绑前，请使用原手机验证验证码").SetToast("换绑前，请验证原手机号验证码"),
	Change_Phone_Step3:      genError(620, "发送验证码").SetToast("换绑前，请先用用手机号发验证码"),
	VisitorAuthError:        genError(621, "无效签名").SetToast("无效签名"),
	WeiXinHaveBind:          genError(622, "该微信已绑定其他账号，暂时无法绑定").SetToast("该微信已绑定其他账号，暂时无法绑定"),
	InvalidPhone:            genError(623, "请输入正确的手机号").SetToast("请输入正确的手机号"),
	PhoneNotRightError:      genError(624, "该手机号不是已绑定手机号").SetToast("该手机号不是已绑定手机号"),
	PhoneHaveRegin:          genError(625, "手机号已注册或者已绑定到其他账户").SetToast("手机号已注册或者已绑定到其他账户"),
	RsaPhoneError:           genError(626, "手机号解密失败").SetToast("手机号解密失败"),
	Apple_Oauth_Error:       genError(627, "苹果认证失败").SetToast("苹果认证失败"),
	LoginDevLimit:           genError(628, "登录设备数受限").SetToast("登录设备数受限"),
	LoginDevLogOut:          genError(629, "登录设备被踢出").SetToast("登录设备被踢出"),
	PhoneLoginCodeError:     genError(630, "输入验证码超过次数，请重新发送验证码").SetToast("输入验证码超过次数，请重新发送验证码"),
	JIGUANGAuthErr:          genError(631, "极光认证失败").SetToast("极光认证失败"),
	RiskCheckError:          genError(632, "风控检测未通过").SetToast("风控检测未通过"),
	NeedWashWhite:           genError(633, "账号有风险需要二次验证").SetToast("账号有风险需要二次验证"),
	Google_Oauth_Error:      genError(634, "google 认证失败").SetToast("google 认证失败"),
	UserNotBindPhone:        genError(635, "用户无绑定记录，无法换绑").SetToast("用户无绑定记录，无法换绑"),
	Facebook_Oauth_Error:    genError(636, "facebook 认证失败").SetToast("facebook 认证失败"),
	EmailError:              genError(637, "输入邮箱有误").SetToast("输入邮箱有误"),
	EmailCodeError:          genError(638, "邮箱验证码错误").SetToast("邮箱验证码错误"),
	EmailCodeCntError:       genError(639, "今日邮箱验证码次数超出限制").SetToast("今日验证码次数超出限制"),
	SendCodeNeedSliderError: genError(640, "验证码需要滑块验证").SetToast("验证码需要滑块验证"),
	SliderCheckError:        genError(641, "滑块验证失败").SetToast("滑块验证失败"),
	EmailFailErr:            genError(642, "邮箱登录失败,请检查相关配置项").SetToast("邮箱登录失败"),
}
