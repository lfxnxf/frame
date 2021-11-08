package school_errors

//
// 基础中台中心错误码
// base:800-899
//
type baseCenterCode struct {
	CensorWordsNoRightAppErr Code

	// 动态配置
	DynCfgDuplicateKey    Code // 已存在
	DynCfgNotExist        Code // 不存在
	DynCfgServiceNameErr  Code // 服务名称错误
	DynCfgGroupNoError    Code // 分组号错误
	DynCfgSpecialGroup    Code // 特殊配置组不可编辑
	DynCfgSpecialCond     Code // 特殊配置此条件不可编辑
	DynCfgValueTypeErr    Code // 值类型不匹配
	DynCfgDictKeyIsEmpty  Code // 字典元素，key不可以为空
	DynCfgListKeyErr      Code // 素组元素，key不可有值
	DynCfgUpdateKeyErr    Code // 更新元素值，key不可修改
	DynCfgPriorityErr     Code // 分组配置不可使用最高或最低优先级
	DynCfgOutOfLength     Code // 参数值过长
	DynCfgIsEmpty         Code // 必传参数值不可为空
	DynCfgNotMeetStandard Code // 字段值不符合规范
	DynCfgActionErr       Code // 未知行为操作
}

var _baseCenterCodes = baseCenterCode{
	CensorWordsNoRightAppErr: genError(800, "请先申请敏感词库").SetToast("请先申请敏感词库"),

	// 动态配置
	DynCfgDuplicateKey:    genError(820, "已存在"),
	DynCfgNotExist:        genError(821, "不存在"),
	DynCfgServiceNameErr:  genError(822, "服务名称错误"),
	DynCfgGroupNoError:    genError(823, "分组号错误"),
	DynCfgSpecialGroup:    genError(824, "特殊配置组不可编辑"),
	DynCfgSpecialCond:     genError(825, "特殊配置此条件不可编辑"),
	DynCfgValueTypeErr:    genError(826, "值类型不匹配"),
	DynCfgDictKeyIsEmpty:  genError(827, "字典元素，key不可以为空"),
	DynCfgListKeyErr:      genError(828, "素组元素，key不可有值"),
	DynCfgUpdateKeyErr:    genError(829, "更新元素值，key不可修改"),
	DynCfgPriorityErr:     genError(830, "分组配置不可使用最高或最低优先级"),
	DynCfgOutOfLength:     genError(831, "参数值过长"),
	DynCfgIsEmpty:         genError(832, "必传参数值不可为空"),
	DynCfgNotMeetStandard: genError(833, "字段值不符合规范"),
	DynCfgActionErr:       genError(834, "未知行为操作"),
}
