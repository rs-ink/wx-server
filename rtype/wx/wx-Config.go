package wx

import "wx-server/config"

type Type int

const (
	_ Type = iota
	OfficialAccount
	MiniProgram
)

func (t Type) String() string {
	switch t {
	case OfficialAccount:
		return "公众号"
	case MiniProgram:
		return "小程序"
	default:
		return "未知"
	}
}

type Config struct {
	AppId     string `json:"appId"`
	AppSecret string `json:"appSecret"`
	Type      Type   `json:"type"`
}

func GetWxsConfig(appId ...string) (conf Config) {
	if len(appId) > 0 {
		conf.AppId = config.Cfg().Wxs.AppId
		conf.AppSecret = config.Cfg().Wxs.AppSecret
	} else {
		conf.AppId = config.Cfg().Wxs.AppId
		conf.AppSecret = config.Cfg().Wxs.AppSecret
	}
	conf.Type = OfficialAccount
	return
}

func GetWxmConfig(appId ...string) (conf Config) {
	if len(appId) > 0 {
		conf.AppId = config.Cfg().Wxm.AppId
		conf.AppSecret = config.Cfg().Wxm.AppSecret
	} else {
		conf.AppId = config.Cfg().Wxm.AppId
		conf.AppSecret = config.Cfg().Wxm.AppSecret
	}
	conf.Type = MiniProgram
	return
}
