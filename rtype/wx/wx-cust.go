package wx

import (
	"fmt"
	"net/http"
	"time"
)

type Type int
const (
	_ Type = iota
	OfficialAccount
	MiniProgram
)

func (t Type)String() string {
	switch t {
	case OfficialAccount:return "公众号"
	case MiniProgram:return "小程序"
	default:
		return "未知"
	}
}

type Config struct {
	AppId string `json:"appId"`
	AppSecret string `json:"appSecret"`
	Type Type `json:"type"`
}

//WxBaseCustomer 微信基础用户标示
type BaseCustomer struct {
	AppId string `json:"appId"`
	OpenId string `json:"openId"`
	UnionId string `json:"unitId"`
}

//WxBaseOther 授权获取的信息
type BaseOther struct {
	Gender           int       `json:"gender"`
	City             string    `json:"city"`
	SubscribeScene   string    `json:"subscribeScene"`
	Subscribe        int       `json:"subscribe"`
	SubscribeTime    time.Time `json:"subscribeTime"`
	GroupId          int       `json:"groupId"`
	Province         string    `json:"province"`
	Country          string    `json:"country"`
}

type UserInfo struct {
	Subscribe      int    `json:"subscribe"`
	OpenId         string `json:"openid"`
	NickName       string `json:"nickname"`
	Sex            int    `json:"sex"`
	Language       string `json:"language"`
	City           string `json:"city"`
	Province       string `json:"province"`
	Country        string `json:"country"`
	HeadImgUrl     string `json:"headimgurl"`
	SubscribeTime  int64  `json:"subscribe_time"`
	UnionId        string `json:"unionid"`
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
	SessionKey   string `json:"session_key"`
	AccessToken  string `json:"access_token"`
	Expires      int    `json:"expires_in"`
	Scope        string `json:"scope"`
	RefreshToken string `json:"refreshToken"`
	OpenId       string `json:"openid"`
	UnionId      string `json:"unionId"`
	CreateTime time.Time
}

