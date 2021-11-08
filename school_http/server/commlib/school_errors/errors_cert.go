package school_errors

type certCenterCode struct {
	//CensorWordsNoRightAppErr Code

	// 实名认证logic
	CertCancelErr Code // 取消实名认证
	UserWashWhiteErr Code //洗白失败

}

var _certCenterCode = certCenterCode{

	// 实名认证logic
	CertCancelErr: genError(1500, "取消认证失败"),
	UserWashWhiteErr : genError(1501, "洗白失败"),
}
