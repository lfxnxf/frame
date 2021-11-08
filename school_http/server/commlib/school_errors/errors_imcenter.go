package school_errors

//
// im中心错误码
// 900-999
//
type imCenterCode struct {
	MsgContentInvalidErr Code

	RoomIDillegal        Code
	RoomTitleTooLong     Code
	RoomDESCtooLong      Code
	RoomUserKickoutError Code

	GroupPermissionDinedError       Code
	GroupGroupAlreadyDismissedError Code
	GroupOwnerCanNotQuitError       Code
	GroupTransCanNotBeAdminError    Code
	GroupTransNotOwnerError         Code
	NotInGroupErr                   Code
	RepeatToAddGroup                Code
	GroupMemberOutLimitErr          Code

	RoomManagePermissionDinedError Code
	RoomAlreadyCloseError          Code
	RoomKickOutError               Code
	RoomUserNotInRoomError         Code
	RoomNotExistError              Code
	RoomOperateCreatorError        Code
	RoomKickOperateSelfError       Code
	RoomForbidOperateSelfError     Code
	RoomAdminOperateSelfError      Code
	RoomNotCreatorError            Code

	RoomChatPerSecondsLimitError           Code
	RoomChatDisabledError                  Code
	RoomChatAudienceChatDurationLimitError Code

	RoomUserIntoSecondDeviceError Code

	RoomUserIsForbidError Code
}

var _imCenterCodes = imCenterCode{
	MsgContentInvalidErr: genError(900, "检测到您的发言内含有敏感词汇，请修改后重试").SetToast("检测到您的发言内含有敏感词汇，请修改后重试"),

	RoomIDillegal:        genError(910, "房间id非法"),
	RoomTitleTooLong:     genError(911, "房间标题过长，请限制在32个字符以内"),
	RoomDESCtooLong:      genError(912, "房间描述过长，请限制在100个字符以内"),
	RoomUserKickoutError: genError(913, "你已被移出直播间"),

	GroupPermissionDinedError:       genError(914, "用户权限不足"),
	GroupGroupAlreadyDismissedError: genError(915, "群聊不存在或已解散"),
	GroupOwnerCanNotQuitError:       genError(916, "群主不能退出群聊"),
	NotInGroupErr:                   genError(917, "用户未在群组中"),
	RepeatToAddGroup:                genError(918, "请勿重复加群"),
	GroupMemberOutLimitErr:          genError(919, "群聊人数超过上限"),
	RoomManagePermissionDinedError:  genError(920, "您的权限不足！"),
	RoomAlreadyCloseError:           genError(921, "房间已经关闭！"),
	RoomKickOutError:                genError(922, "您已被踢出该房间，暂不可进入！"),
	RoomUserNotInRoomError:          genError(923, "用户不在该房间！"),
	RoomNotExistError:               genError(924, "房间不存在！"),
	RoomOperateCreatorError:         genError(925, "非法操作，您的操作对象是房主！"),
	RoomKickOperateSelfError:        genError(926, "非法操作，您不能将自己踢出房间！"),
	RoomForbidOperateSelfError:      genError(927, "非法操作，您不能禁言自己！"),
	RoomAdminOperateSelfError:       genError(928, "非法操作，您不能操作自己"),
	RoomNotCreatorError:             genError(929, "您当前没有创建房间！"),

	GroupTransCanNotBeAdminError: genError(930, "非法操作,操作对象不能是管理员或群主"),
	GroupTransNotOwnerError:      genError(931, "非法操作,用户不是群主,无法执行群主转让"),

	RoomChatAudienceChatDurationLimitError: genError(932, "發言頻率太快了～"),
	RoomChatPerSecondsLimitError:           genError(933, "當前網絡有延遲，請稍後重試"),
	RoomChatDisabledError:                  genError(934, "當前網絡有延遲，請稍後重試"),

	RoomUserIntoSecondDeviceError: genError(935, "您已在其他设备进入该房间，请勿重复进入！"),

	RoomUserIsForbidError: genError(936, "您已被禁言！"),
}
