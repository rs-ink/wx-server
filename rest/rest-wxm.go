package rest

import (
	"github.com/devfeel/dotweb"
	"wx-server/middleware"
	"wx-server/rlog"
	"wx-server/rtype"
	"wx-server/rtype/wx"
	"wx-server/token"
)

func initWxmRest(group dotweb.Group) {
	g := group.Group("/wxm")

	g.POST(code())
	g.POST(info())
}

//TODO 解密参数有问题 2019-11-20 11:23:40 frank
func info() (path string, handle dotweb.HttpHandle) {
	path = "/info"
	handle = func(ctx dotweb.Context) error {
		var param struct {
			wx.EncryptedData
		}
		rlog.Warn(string(ctx.Request().PostBody()))
		_ = ctx.BindJsonBody(&param)
		appSession := middleware.GetAppSession(ctx)
		result, err := token.Decrypt(appSession.WxSession, param.EncryptedData)
		rlog.Warn(result, err)

		return ctx.WriteJson(rtype.Success())
	}
	return
}

func code() (path string, handle dotweb.HttpHandle) {
	path = "/code"
	handle = func(ctx dotweb.Context) error {
		var param struct {
			Code string `json:"code"`
		}
		rlog.Warn(string(ctx.Request().PostBody()))

		wxm := rtype.GetWxmConfig()
		_ = ctx.BindJsonBody(&param)
		if param.Code != "" {
			mis, _ := token.GetWxSession(param.Code, wxm)
			rlog.WarnF("%+v", mis)
		}
		return ctx.WriteJson(rtype.Success())
	}
	return
}
