package school_errors

//
// 金融中台中心错误码
// financial:1000-1099
//
type financialCenterCode struct {
	// 金融公共错误码 1000 ～ 1090
	OrderSuccessRepeatError Code
	ReqeustFrequentlyError  Code
	ThirdSysReturnError     Code
	OrderInfoNotMatchError  Code
	AuthError               Code
	OrderInProcessingError  Code

	// 虚拟币服务错误码 1020~1029
	CoinNotEnoughtError             Code
	SourceError                     Code
	CoinKindNotExistError           Code
	CoinKindNotSupportExchangeError Code
	CoinTransferWorthError          Code
	OrderNotAllowRefundError        Code
	CoinTransferToInfoErr           Code
	CoinTransFromUserNotBeToUser    Code

	// 支付服务 1030~1039

	// 提现服务 1040 ~ 1049
	WithDrawalExceedLimitError Code
	WithdrawalFailedError      Code

	//电商支付服务错误码
	GenOrderIDError                   Code
	IdInfoError                       Code
	OrganizationInfoError             Code
	BusinessInfoError                 Code
	OrderNotExistsError               Code
	UnsupportedMannerError            Code
	PaidTimeOutError                  Code
	ProfitSharingInfoError            Code
	PriceMissmatchError               Code
	WxRetryError                      Code
	ItemNotExistsError                Code
	ProfitSharingNotAllowed           Code
	RefundMissMatchErr                Code
	RefundNotAllowedErr               Code
	RefundPostageErr                  Code
	MchNotExistsError                 Code
	TaskInProcessingError             Code
	OrderAlreadyFullyRefundError      Code
	ProfitSharingAlreadyFinishedError Code
	SubMchSettlementInfoError         Code
	SettlementInfoNotExistsError      Code
	BalanceNotEnoughError             Code
	OrderFinishedError                Code
}

var _financialCenterCodes = financialCenterCode{

	OrderSuccessRepeatError:           genError(1001, "订单已被处理，重复的交易订单").SetToast("重复的交易订单"),
	ThirdSysReturnError:               genError(1002, "第三方系统返回非零错误码").SetToast("第三方系统返回非零错误码,请确保请求参数正确"),
	ReqeustFrequentlyError:            genError(1003, "请求太频繁了，请稍等片刻").SetToast("请求太频繁了，请稍等片刻"),
	AuthError:                         genError(1004, "请求权限错误").SetToast("请求权限错误"),
	OrderInfoNotMatchError:            genError(1005, "订单信息不匹配").SetToast("订单信息不匹配"),
	OrderInProcessingError:            genError(1006, "订单正在处理中").SetToast("订单正在处理中，请稍后再试"),
	WithDrawalExceedLimitError:        genError(1040, "提现金额超出限额").SetToast("提现金额超出限额"),
	WithdrawalFailedError:             genError(1041, "提现失败").SetToast("提现失败"),
	CoinNotEnoughtError:               genError(1020, "货币余额不足").SetToast("您的余额不足,请保证余额充足后重试"),
	OrderNotAllowRefundError:          genError(1021, "订单类型不允许退还货币").SetToast("订单类型不允许退还货币"),
	SourceError:                       genError(1022, "source校验未通过").SetToast("source校验未通过"),
	CoinKindNotExistError:             genError(1023, "货币种类不支持").SetToast("货币种类不支持"),
	CoinKindNotSupportExchangeError:   genError(1024, "货币兑换种类不支持").SetToast("货币兑换种类不支持"),
	CoinTransferWorthError:            genError(1025, "流入货币总价值超出流出货币").SetToast("流入货币总价值超出流出货币"),
	CoinTransferToInfoErr:             genError(1026, "流入货币信息不允许有重复").SetToast("流入货币信息不允许有重复"),
	CoinTransFromUserNotBeToUser:      genError(1027, "流出用户不允许成为流入用户").SetToast("流出用户不允许成为流入用户"),
	GenOrderIDError:                   genError(1050, "订单id生成错误").SetToast("订单id生成错误"),
	IdInfoError:                       genError(1051, "身份信息数据错误").SetToast("身份信息数据错误,请检查数据是否正确"),
	OrganizationInfoError:             genError(1052, "组织信息数据错误").SetToast("组织信息数据错误,请检查数据是否正确"),
	BusinessInfoError:                 genError(1053, "营业执照信息数据错误").SetToast("营业执照信息数据错误,请检查数据是否正确"),
	OrderNotExistsError:               genError(1054, "订单未找到或不存在").SetToast("订单未找到或不存在"),
	UnsupportedMannerError:            genError(1055, "不支持的manner类型").SetToast("不支持的manner类型,请验证"),
	PaidTimeOutError:                  genError(1056, "订单超时失效").SetToast("订单已经超时失效,请重新下单"),
	ProfitSharingInfoError:            genError(1057, "分账信息错误").SetToast("分账信息错误"),
	PriceMissmatchError:               genError(1058, "商品总价和订单金额不匹配").SetToast("价格信息错误,请验证订单金额和商品价格是否填写正确"),
	WxRetryError:                      genError(1059, "接口异常,请重试").SetToast("接口异常,请重试"),
	ItemNotExistsError:                genError(1060, "商品信息未找到或不存在").SetToast("商品信息未找到或不存在"),
	ProfitSharingNotAllowed:           genError(1061, "商品当前无法分账").SetToast("商品当前无法分账, 请确认是否有未结束的退款流程"),
	RefundMissMatchErr:                genError(1062, "商品退款总计和订单退款不匹配").SetToast("商品退款总计和订单退款不匹配"),
	RefundNotAllowedErr:               genError(1063, "商品无法退款").SetToast("商品已经进行分账，不支持退款"),
	MchNotExistsError:                 genError(1064, "商户不存在").SetToast("商户不存在，请检查商户是否通过审核"),
	TaskInProcessingError:             genError(1065, "任务处理中").SetToast("当前任务正在处理中,请勿重复提交"),
	OrderAlreadyFullyRefundError:      genError(1066, "订单退款已达上限").SetToast("订单退款已达上限,请勿发起无效的退款"),
	ProfitSharingAlreadyFinishedError: genError(1067, "商品分账已达上限").SetToast("商品分账已达上限,请勿发起无效的分账"),
	SubMchSettlementInfoError:         genError(1068, "提现银行卡信息错误").SetToast("提现银行卡信息错误, 请提交正确的银行卡信息"),
	SettlementInfoNotExistsError:      genError(1069, "商户银行卡未绑定或未找到").SetToast("商户银行卡未绑定或未找到, 请先绑定银行卡"),
	BalanceNotEnoughError:             genError(1070, "商户余额不足").SetToast("商户余额不足, 请确保余额充足的情况下重试"),
	RefundPostageErr:                  genError(1071, "退款邮费与订单原始邮费不符").SetToast("退款邮费与订单原始邮费不符, 请下检查后重试"),
	OrderFinishedError:                genError(1072, "订单已完成").SetToast("订单已完成, 无法进行更多操作"),
}
