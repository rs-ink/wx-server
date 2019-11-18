package rtype

import (
	"wx-server/config"
	"wx-server/rtype/wx"
)

func GetWxConfig(appId ...string) (conf wx.Config) {
	if len(appId)>0{
		conf.AppId=config.Cfg().Wxs.AppId
		conf.AppSecret=config.Cfg().Wxs.AppSecret
	}else{
		conf.AppId=config.Cfg().Wxs.AppId
		conf.AppSecret=config.Cfg().Wxs.AppSecret
	}
	return
}