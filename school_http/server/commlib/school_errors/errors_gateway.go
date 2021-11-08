package school_errors

//
// 限流中心错误码
// commerce:1400-1499
//

type gateWayErr struct {
	FalconTeamQueryLengthTooLongErr Code
	AssignTeamTooMuchErr            Code
	AssignNoTeamErr    				Code
}

var _gateWayCodes = gateWayErr{
	FalconTeamQueryLengthTooLongErr: genError(1400, "搜索条件请勿超过64个字符"),
	AssignTeamTooMuchErr: genError(1401, "报警组请勿超过10个"),
	AssignNoTeamErr:      genError(1402, "危险操作！至少指定一个报警组").SetToast("危险操作！至少指定一个报警组"),
}