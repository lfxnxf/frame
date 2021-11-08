package school_errors

//
// media中心错误码
// media:700-799
//
type mediaCenterCode struct {
	MediaHeartbeatExpired          Code
	UserAlreadyOnLive              Code
	UserAlreadyOnMic               Code
	MediaHeartbeatWrongUID         Code
	LiveNotFound                   Code
	LiveStatusError                Code
	LiveSlotAlreadyUsed            Code
	LivePermissionDenied           Code
	ApplyMicUserNotInRoom          Code
	AnchorStreamNotFound           Code
	MicApplyingNotFound            Code
	MicInviteNotFound              Code
	NoSelfLiveSlotInLive           Code
	LiveRoomInfoNotFound           Code
	MicCardNotEnough               Code
	MediaOpenIdNotFound            Code
	MediaUserCertifyFailed         Code
	MediaUserCheckBindMobileFailed Code
	LiveUserNotFound               Code
	LiveIdNotFound                 Code
	LiveRoomIdNotFound             Code
	StartLivePermissionDenied      Code
	LiveUserGenderMissed           Code
	LiveUserInBlacklist            Code
	LiveAlreadyStopped             Code
	LiveAlreadyStarted             Code
	UserAlreadyOnNotAvailable      Code
	CallAlreadyEnd                 Code
	UserAlreadyCalling             Code
	DontRepeatStart                Code
	LiveTitleIllegal               Code
	TimeIntervalOverlap            Code // 时间区间重叠
	LiveOperPermissionDenied       Code // 没有直播运营权限
	LiveUserGenderWrong            Code
	LiveNoValidSlot                Code
	LiveOwnerBeOccupied            Code

	AlreadyStartCallErr Code
	CallSelfErr         Code
	CallRecordNotExist  Code
	InValidCall         Code
	SelfBusyErr         Code
	PeerBusyErr         Code
	IllegalErr  		Code
}

var _mediaCenterCodes = mediaCenterCode{
	MediaHeartbeatExpired:          genError(700, "心跳超时").SetToast("心跳超时"),
	UserAlreadyOnLive:              genError(701, "检测到异常退出，请1分钟后重试").SetToast("检测到异常退出，请1分钟后重试"),
	UserAlreadyOnMic:               genError(702, "正在连麦中").SetToast("正在连麦中"),
	MediaHeartbeatWrongUID:         genError(703, "User ID is wrong").SetToast("错误UID"),
	LiveNotFound:                   genError(704, "直播不存在").SetToast("直播不存在"),
	LiveStatusError:                genError(710, "直播状态错误").SetToast("直播状态错误"),
	LiveSlotAlreadyUsed:            genError(711, "麦位已占用").SetToast("麦位已占用"),
	LivePermissionDenied:           genError(712, "用户无权限").SetToast("用户权限错误"),
	ApplyMicUserNotInRoom:          genError(713, "用户不在当前房间").SetToast("用户已不在当前房间"),
	AnchorStreamNotFound:           genError(714, "Anchor stream not found").SetToast("未找到直播信息"),
	MicApplyingNotFound:            genError(715, "用户不在该房间").SetToast("用户已不在该房间"),
	MicInviteNotFound:              genError(716, "未找到邀请上麦记录").SetToast("未找到邀请上麦记录"),
	NoSelfLiveSlotInLive:           genError(717, "未找到麦位信息").SetToast("未找到麦位信息"),
	LiveRoomInfoNotFound:           genError(718, "Room info not found").SetToast("未找到房间信息"),
	MicCardNotEnough:               genError(719, "Mic card not enough").SetToast("上麦卡不足"),
	MediaOpenIdNotFound:            genError(720, "User openid not found").SetToast("找不到OpenId"),
	MediaUserCertifyFailed:         genError(721, "User certify failed").SetToast("用户未实名认证"),
	MediaUserCheckBindMobileFailed: genError(722, "User check bind mobile failed").SetToast("用户未绑定手机"),
	LiveUserNotFound:               genError(723, "User not found").SetToast("用户不存在"),
	LiveIdNotFound:                 genError(724, "直播间不存在").SetToast("直播ID不存在"),
	LiveRoomIdNotFound:             genError(725, "Room ID not found").SetToast("房间ID不存在"),
	StartLivePermissionDenied:      genError(726, "Start live permission denied").SetToast("您没有开播权限，如需开播，请在「我的资料」联系客服"),
	LiveUserGenderMissed:           genError(727, "Gender of the user is not set").SetToast("用户没有设置性别，不能连麦"),
	LiveUserInBlacklist:            genError(728, "User is in blacklist").SetToast("用户账号已被限制"),
	LiveAlreadyStopped:             genError(729, "直播已结束，请重新开播").SetToast("直播已停播，请重新开播"),
	LiveAlreadyStarted:             genError(730, "直播已经开始，不可重复开播").SetToast("直播已经开始，不可重复开播"),

	// 1v1
	UserAlreadyOnNotAvailable: genError(731, "peer is busy now, please try later").SetToast("对方通话中，请稍候再试"),
	CallAlreadyEnd:            genError(732, "call has been ended.").SetToast("通话已结束"),
	UserAlreadyCalling:        genError(733, "please prepare after you cancel.").SetToast("请勿重复呼叫"),
	DontRepeatStart:           genError(734, "call has been connected, dont repeat start."),
	LiveTitleIllegal:          genError(740, "live title illegal.").SetToast("直播标题非法"),
	AlreadyStartCallErr:       genError(741, "already start call").SetToast("通话已开始"),
	CallSelfErr:               genError(742, "can not call yourself").SetToast("不能和自己通话"),
	CallRecordNotExist:        genError(743, "call record not exist").SetToast("通话不存在"),
	InValidCall:			   genError(744, "invalid call").SetToast("无效的通话"),
	SelfBusyErr:    		   genError(745, "self is busy now, please call later").SetToast("当前用户正忙"),
	PeerBusyErr:    		   genError(746, "peer is busy now, please call later").SetToast("对方正忙，请稍后重试"),
	IllegalErr:     		   genError(747, "illegal operation").SetToast("非法操作"),
	// 1v1 end
	// live robot
	TimeIntervalOverlap:      genError(750, "time interval overlap").SetToast("时间区间有重叠"),
	LiveOperPermissionDenied: genError(751, "permission denied").SetToast("没有运营权限"),
	// live robot end
	// live buz
	LiveUserGenderWrong: genError(752, "live user gender wrong").SetToast("用户性别设置错误，不能连麦"),
	LiveNoValidSlot:     genError(753, "麦位已满").SetToast("麦位已满"),
	// live buz end
	LiveOwnerBeOccupied: genError(760, "live owner be occupied").SetToast("慢了一步，主麦已经有人了"),
}
