package rtype

import (
	"wx-server/config"
	"wx-server/rtype/wx"
)

func GetWxsConfig(appId ...string) (conf wx.Config) {
	if len(appId) > 0 {
		conf.AppId = config.Cfg().Wxs.AppId
		conf.AppSecret = config.Cfg().Wxs.AppSecret
	} else {
		conf.AppId = config.Cfg().Wxs.AppId
		conf.AppSecret = config.Cfg().Wxs.AppSecret
	}
	conf.Type = wx.OfficialAccount
	return
}

func GetWxmConfig(appId ...string) (conf wx.Config) {
	if len(appId) > 0 {
		conf.AppId = config.Cfg().Wxm.AppId
		conf.AppSecret = config.Cfg().Wxm.AppSecret
	} else {
		conf.AppId = config.Cfg().Wxm.AppId
		conf.AppSecret = config.Cfg().Wxm.AppSecret
	}
	conf.Type = wx.MiniProgram
	return
}
