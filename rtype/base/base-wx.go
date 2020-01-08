package base

//WxAccount 微信基础用户标示
type WxBaseAccount struct {
	AppId   string `json:"appId"`
	OpenId  string `json:"openId"`
	UnionId string `json:"unitId"`
}
