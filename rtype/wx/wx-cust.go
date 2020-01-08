package wx

import (
	"encoding/gob"
	"time"
)

func init() {
	gob.Register(Session{})
	gob.Register(UserInfoOfficialAccount{})
}

type UserInfoOfficialAccount struct {
	Subscribe      int    `json:"subscribe"`
	OpenId         string `json:"openid"`
	UnionId        string `json:"unionid"`
	NickName       string `json:"nickname"`
	Sex            int    `json:"sex"`
	Language       string `json:"language"`
	City           string `json:"city"`
	Province       string `json:"province"`
	Country        string `json:"country"`
	HeadImgUrl     string `json:"headimgurl"`
	SubscribeTime  int64  `json:"subscribe_time"`
	Remark         string `json:"remark"`
	GroupId        int    `json:"groupid"`
	TagIdList      []int  `json:"tagid_list"`
	SubscribeScene string `json:"subscribe_scene"`
	QrScene        string `json:"qr_scene"`
	QrSceneStr     string `json:"qr_scene_str"`
}

type Session struct {
	Error
	AppId        string `json:"appId"`
	WxType       Type   `json:"type"`
	SessionKey   string `json:"session_key"`
	AccessToken  string `json:"access_token"`
	Expires      int    `json:"expires_in"`
	Scope        string `json:"scope"`
	RefreshToken string `json:"refreshToken"`
	OpenId       string `json:"openid"`
	UnionId      string `json:"unionId"`
	CreateTime   time.Time
}
