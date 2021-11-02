package lib

type BaseResp struct {
	DmError  int    `json:"dm_error"`
	ErrorMsg string `json:"error_msg"`
}
type UserInfoResp struct {
	BaseResp
	Infos []RUserInfo `json:"infos"`
}
type RUserInfo struct {
	User     UserInfo `json:"user"`
	Relation string   `json:"relation"`
}
type UserInfo struct {
	Birth          string `json:"birth"`
	CurrentValue   string `json:"current_value"`
	Description    string `json:"description"`
	Emotion        string `json:"emotion"`
	Gender         int    `json:"gender"`
	Gmutex         int    `json:"gmutex"`
	Hometown       string `json:"hometown"`
	ID             int64  `json:"id"`
	InkeVerify     int    `json:"inke_verify"`
	Level          int    `json:"level"`
	Location       string `json:"location"`
	NextDiff       string `json:"next_diff"`
	Nick           string `json:"nick"`
	Portrait       string `json:"portrait"`
	Profession     string `json:"profession"`
	RankVeri       int    `json:"rank_veri"`
	RegisterAt     int    `json:"register_at"`
	Sex            int    `json:"sex"`
	ThirdPlatform  string `json:"third_platform"`
	VeriInfo       string `json:"veri_info"`
	Verified       int    `json:"verified"`
	VerifiedPrefix string `json:"verified_prefix"`
	VerifiedReason string `json:"verified_reason"`
}
type VersatileAtom struct {
	Uid       interface{} `schema:"uid" json:"uid"`
	LiveOwner interface{} `schema:"live_owner" json:"live_owner"`
	Lc        string      `schema:"lc"  json:"lc"`
	Cc        string      `schema:"cc" json:"cc"`
	Cv        string      `schema:"cv" json:"cv"`
	Ua        string      `schema:"ua" json:"ua"`
	Conn      string      `schema:"conn" json:"conn"`
	Devi      string      `schema:"devi" json:"devi"`
	Idfv      string      `schema:"idfv" json:"idfv"`
	Idfa      string      `schema:"idfa" json:"idfa"`
	Proto     string      `schema:"proto" json:"proto"`
	Osversion string      `schema:"osversion" json:"osversion"`
	Logid     string      `schema:"logid" json:"logid"`
	Smid      string      `schema:"smid" json:"smid"`
	Xrealip   string      `schema:"xrealip" json:"xrealip"`
	Location  string      `schema:"location" json:"location"`
	Mjid      interface{} `schema:"mjid" json:"mjid"`
}

type Atom struct {
	Uid       int64  `schema:"uid" json:"uid"`
	LiveOwner int64  `schema:"live_owner" json:"live_owner"`
	Lc        string `schema:"lc"  json:"lc"`
	Cc        string `schema:"cc" json:"cc"`
	Cv        string `schema:"cv" json:"cv"`
	Ua        string `schema:"ua" json:"ua"`
	Conn      string `schema:"conn" json:"conn"`
	Devi      string `schema:"devi" json:"devi"`
	Idfv      string `schema:"idfv" json:"idfv"`
	Idfa      string `schema:"idfa" json:"idfa"`
	Proto     string `schema:"proto" json:"proto"`
	Osversion string `schema:"osversion" json:"osversion"`
	Logid     string `schema:"logid" json:"logid"`
	Smid      string `schema:"smid" json:"smid"`
	Xrealip   string `schema:"xrealip" json:"xrealip"`
	Location  string `schema:"location" json:"location"`
	Mjid      int32  `schema:"mjid" json:"mjid"`
}

type RequestKick struct {
	Ids []int64 `json:"ids"`
}

type RequestForbid struct {
	To int64 `json:"to"`
}
type RequestLiveChange struct {
	Action string `json:"action"`
}
