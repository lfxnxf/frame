/**
错误码封装
*/
package school_errors

//
// 通用错误码
// 0~599
//
var Codes = struct {
	Success Code

	Redirect Code // 跳转

	ClientError        Code // 请求参数错误（调用端参数有问题）
	CSignExpired       Code // 签名过期
	CSignError         Code // 签名错误
	CSignRepeat        Code // 签名重复（防止重放）
	CAppKeyIsNone      Code // 未传app key
	CAppKeyUnsupported Code // 不支持的app key
	CNumOverLimit      Code // 请求个数超过限制

	//
	// Server System Error Code
	//
	ServerError Code // 服务器遇到了一个未曾预料的状况，导致了它无法完成对请求的处理。一般来说，这个问题都会在服务器端的源代码出现错误时出现。

	//
	// Server net error
	//
	SNotImplemented          Code // 服务器不支持当前请求所需要的某个功能。当服务器无法识别请求的方法，并且无法支持其对任何资源的请求。
	SBadGateway              Code // 作为网关或者代理工作的服务器尝试执行请求时，从上游服务器接收到无效的响应。
	SServiceUnavailable      Code // 由于临时的服务器维护或者过载，服务器当前无法处理请求。这个状况是临时的，并且将在一段时间以后恢复。如果能够预计延迟时间，那么响应中可以包含一个 Retry-After 头用以标明这个延迟时间。如果没有给出这个 Retry-After 信息，那么客户端应当以处理500响应的方式处理它。
	SGatewayTimeout          Code // 作为网关或者代理工作的服务器尝试执行请求时，未能及时从上游服务器（URI标识出的服务器，例如HTTP、FTP、LDAP）或者辅助服务器（例如DNS）收到响应（某些代理服务器在DNS查询超时时会返回400或者500错误）。
	SHttpVersionNotSupported Code // 服务器不支持，或者拒绝支持在请求中使用的 HTTP 版本。这暗示着服务器不能或不愿使用与客户端相同的版本。响应中应当包含一个描述了为何版本不被支持以及服务器支持哪些协议的实体。
	SVariantAlsoNegotiates   Code // 由《透明内容协商协议》（RFC 2295）扩展，代表服务器存在内部配置错误：被请求的协商变元资源被配置为在透明内容协商中使用自己，因此在一个协商处理中不是一个合适的重点。
	SInsufficientStorage     Code // 服务器无法存储完成请求所必须的内容。这个状况被认为是临时的。WebDAV (RFC 4918)
	SBandwidthLimitExceeded  Code // 服务器达到带宽限制。这不是一个官方的状态码，但是仍被广泛使用。
	SNotExtended             Code // 获取资源所需要的策略并没有被满足。（RFC 2774）

	//
	// Server code or logic exception
	//
	SException Code // 服务内部错误

	SAlarm Code // 监控系统收到此错误码，将立即报警

	//
	// Account
	//
	accountCode

	//
	// im center
	//
	imCenterCode

	//
	// media center
	//
	mediaCenterCode

	//
	//基础中心 base
	//
	baseCenterCode

	//
	//金融中心
	//
	financialCenterCode

	//
	//资源中心
	//
	sourceCenterCode

	//
	//增值付费中心
	//
	commerceCode

	//
	// Profile
	//
	profileCode

	//
	// Gateway
	//
	gateWayErr

	//
	//
	//
	certCenterCode

	//
	// 映客娱乐云
	//
	controlCenterCode
}{
	Success:            genError(0, "success").SetToast("操作成功"),
	Redirect:           genError(302, ""),
	ClientError:        genError(499, "请求参数错误").SetToast("网络异常，请重试"),
	CSignExpired:       genError(451, "请求异常").SetToast("网络异常，请重试"),
	CSignError:         genError(452, "请求异常").SetToast("网络异常，请重试"),
	CSignRepeat:        genError(453, "请求异常").SetToast("网络异常，请重试"),
	CNumOverLimit:      genError(454, "请求数量超过限制").SetToast("网络异常，请重试"),
	CAppKeyIsNone:      genError(461, "not found appkey").SetToast("网络异常，请重试"),
	CAppKeyUnsupported: genError(462, "unsupported appkey").SetToast("网络异常，请重试"),

	ServerError: genError(500, "内部系统错误").SetToast("系统开小差了，请稍后重试"),

	SNotImplemented:          genError(501, "不支持此请求").SetToast("网络开小差了，请稍后重试"),
	SBadGateway:              genError(502, "依赖接口数据错误").SetToast("网络开小差了，请稍后重试"),
	SServiceUnavailable:      genError(503, "服务暂时不可用").SetToast("网络开小差了，请稍后重试"),
	SGatewayTimeout:          genError(504, "访问依赖接口超时").SetToast("网络开小差了，请稍后重试"),
	SHttpVersionNotSupported: genError(505, "不支持的 HTTP 版本").SetToast("网络开小差了，请稍后重试"),
	SVariantAlsoNegotiates:   genError(506, "服务内部配置错误").SetToast("网络开小差了，请稍后重试"),
	SInsufficientStorage:     genError(507, "关键存储路径出错").SetToast("网络开小差了，请稍后重试"),
	SBandwidthLimitExceeded:  genError(509, "服务器带宽超限").SetToast("网络开小差了，请稍后重试"),
	SNotExtended:             genError(510, "资源依赖访问策略不足").SetToast("网络开小差了，请稍后重试"),

	SException: genError(540, "服务内部错误").SetToast("系统开小差了，工程师们正在紧急修复，请稍后重试"),

	SAlarm: genError(599, "服务出现严重故障").SetToast("数据异常，请联系客服"),

	//
	// AccountCodes
	//
	accountCode: _accountCodes,

	//
	// imCenterCodes
	//
	imCenterCode: _imCenterCodes,

	//
	// medeaCenterCodes
	//
	mediaCenterCode: _mediaCenterCodes,

	//
	// baseCenterCodes
	//
	baseCenterCode: _baseCenterCodes,

	//
	// financialCenterCodes
	//
	financialCenterCode: _financialCenterCodes,

	//
	// sourceCenterCodes
	//
	sourceCenterCode: _sourceCenterCodes,

	//
	commerceCode: _commerceCodes,

	//
	// ProfileCodes
	//
	profileCode: _profileCodes,

	//
	// Gateway
	//
	gateWayErr: _gateWayCodes,

	//
	//Cert
	//
	certCenterCode: _certCenterCode,
}
