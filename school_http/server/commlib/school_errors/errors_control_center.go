package school_errors

// 映客娱乐云错误码 1600～1699
type controlCenterCode struct {
	InkeCloudTicketNotValidError Code
}

var controlCenterCodes = controlCenterCode{
	InkeCloudTicketNotValidError: genError(1601, "ticket不合法"),
}
