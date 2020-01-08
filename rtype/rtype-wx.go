package rtype

import "wx-server/rtype/wx"

//noinspection GoUnusedConst
const (
	_ ErrCode = WxBaseErrCode + iota
	ErrCodeNoPhone
	ErrCodeNoDetailInfo
	ErrCodeUpdatePhoneError
	ErrCodeWxIno
)

type FromChannel int

const (
	FromInit FromChannel = iota
	FromSingleMessage
	FromGroupMessage
	FromTimeline
)

func (fromType FromChannel) String() string {
	switch fromType {
	case FromSingleMessage:
		return "私聊"
	case FromGroupMessage:
		return "群组"
	case FromTimeline:
		return "朋友圈"
	default:
		return "无法识别"
	}
}

//判断微信来源
//noinspection GoUnusedExportedFunction
func FromParse(from string) FromChannel {
	switch from {
	case "singlemessage":
		return FromSingleMessage
	case "groupmessage":
		return FromGroupMessage
	case "timeline":
		return FromTimeline
	default:
		return FromInit
	}
}

type MiniSession struct {
	wx.Error
	AppId        string `json:"appId"`
	SessionKey   string `json:"session_key"`
	AccessToken  string `json:"access_token"`
	Expires      int    `json:"expires_in"`
	Scope        string `json:"scope"`
	RefreshToken string `json:"refreshToken"`
	OpenId       string `json:"openid"`
	UnionId      string `json:"unionId"`
}

func (wxm *MiniSession) IsNull() bool {
	return wxm.OpenId == ""
}
