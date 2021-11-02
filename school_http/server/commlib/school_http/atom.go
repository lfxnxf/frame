package school_http

/**
 * 请求原子信息。客户端每个到达服务器的请求，均在url后拼装原子信息参数
 * 示例：http://service.imilive.cn/api/base/serviceinfo/get_all?uid=123456&sid=546asdf4af5asd56f14a6ds5f46as5df4a&cv=GA1.0.60_Android&vc=12345&deviceId=ioasdfasdf84asdf56afsd&devi=fdas415asdf&imei=4asd5f5as&imsi=asdfa4
 * 参考：http://wiki.inkept.cn/pages/viewpage.action?pageId=19528498
 */
type Atom struct {
	Uid       int64  `json:"uid" schema:"uid" url:"uid"`                   // 用户uid
	Sid       string `json:"sid" schema:"sid" url:"sid"`                   // 用户登录 session id
	CV        string `json:"cv" schema:"cv" url:"cv"`                      // 客户端版本。示例：GA1.0.50_Android
	Cpu       string `json:"cpu" schema:"cpu" url:"cpu"`                   // cpu型号。示例：%5BAdreno_%28TM%29_308%5D%5BAArch64_638_Qualcomm_Technologies%2C_Inc_MSM8917%5D
	Gpu       string `json:"gpu" schema:"gpu" url:"gpu"`                   // gpu型号
	UA        string `json:"ua" schema:"ua" url:"ua"`                      // 客户端设备型号。示例：vivovivoY66iA、XiaomiRedmi5A
	Devi      string `json:"devi" schema:"evid" url:"evid"`                // 设备号。示例：5675e48f7d74
	Imei      string `json:"imei" schema:"meid" url:"meid"`                // SIM卡槽号。示例：865990031815521
	Imsi      string `json:"imsi" schema:"msid" url:"msid"`                // SIM卡号。示例：460000835503231
	SMID      string `json:"smid" schema:"smid" url:"smid"`                // 设备指纹-数盟（设备指纹-数盟、设备指纹-数美 有其一即可）
	NDID      string `json:"ndid" schema:"ndid" url:"ndid"`                // 设备指纹-数美（设备指纹-数盟、设备指纹-数美 有其一即可）
	UDID      string `json:"udid" schema:"udid" url:"udid"`                // （移动智能终端补充设备标识）设备唯⼀标识符最长64位，返 回空字符串表⽰不⽀持，异常状态包括⽹络异常、appid异常、应⽤异常等
	OAID      string `json:"oaid" schema:"oaid" url:"oaid"`                // （移动智能终端补充设备标识）匿名设备标识符最长64位，返 回空字符串表⽰不⽀持，异常状态包括⽹络异常、appid异常、应⽤异常等
	VAID      string `json:"vaid" schema:"vaid" url:"vaid"`                // （移动智能终端补充设备标识）开发者匿名设备标识符最长64位，返回空字符串表⽰不⽀持，异常状态包括⽹络异常、appid异常、应⽤异常等
	AAID      string `json:"aaid" schema:"aaid" url:"aaid"`                // （移动智能终端补充设备标识）应⽤匿名设备标识符最长64位，返回空字符串表⽰不⽀持，异常状态包括⽹络异常、appid异常、应⽤异常等
	OSVersion string `json:"osversion" schema:"osversion" url:"osversion"` // 操作系统版本。示例：android_25
	Conn      string `json:"conn" schema:"conn" url:"conn"`                // 网络连接类型。示例：wifi、5G、4G、3G、2G、GPRS、Unknown
	CC        string `json:"cc" schema:"cc" url:"cc"`                      // 客户端渠道号。示例：TG45620
	Mtxid     string `json:"mtxid" schema:"mtxid" url:"mtxid"`             // 对应的网卡的MAC地址。示例：06:69:6c:eb:6f:0d
	Ram       string `json:"ram" schema:"ram" url:"ram"`                   // 手机运行内存大小
	Icc       string `json:"icc" schema:"icc" url:"icc"`                   // 手机sim卡的序号。示例：898600f10116f0025658
	Ast       string `json:"ast" schema:"ast" url:"ast"`                   //
	Aid       string `json:"aid" schema:"aid" url:"aid"`                   // ANDROID_ID。ANDROID_ID是Android系统第一次启动时产生的一个64bit（16BYTES）数，如果设备被wipe还原后，该ID将被重置（变化）
	Mtid      string `json:"mtid" schema:"mtid" url:"mtid"`                // 路由器发射的无线信号的名。示例： "Inke-Office"
	Idfa      string `json:"idfa" schema:"idfa" url:"idfa"`                // iOS idfa信息
	LogID     string `json:"logid" schema:"logid" url:"logid"`
	XRealIP   string `json:"xrealip" schema:"xrealip" url:"xrealip"` // 客户端入网IP（外网IP）
}

/**
 * 请求原子信息。Web平台每个到达服务器的请求，均在url后拼装原子信息参数
 * 示例：http://api.meeshow.com/api_web/v1/controlcenter/fps/serviceinfo/list/get?user_id=123456&ticket=q3sar2rar&system_id=21&username=&email=&department=
 * 参考：https://wiki.inkept.cn/pages/viewpage.action?pageId=104675262
 */
type AtomWeb struct {
	UserID     string `json:"user_id" schema:"user_id" url:"user_id"`
	Ticket     string `json:"ticket" schema:"ticket" url:"ticket"`
	SystemID   string `json:"system_id" schema:"system_id" url:"system_id"`
	Username   string `json:"username" schema:"username" url:"username"`
	Email      string `json:"email" schema:"email" url:"email"`
	Department string `json:"department" schema:"department" url:"department"`
}
