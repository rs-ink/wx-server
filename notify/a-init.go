package notify

import (
	"encoding/xml"
	"fmt"
	"github.com/devfeel/dotweb"
	"html/template"
	"wx-server/config"
	"wx-server/rlog"
	"wx-server/rtype/open"
	"wx-server/rtype/wx"
	"wx-server/token"
	"wx-server/util"
)

func InitNotifyRouter(g dotweb.Group) {
	g.POST(wxNotifyMsg())
	g.GET(wxNotifyMsg())

	g.GET(wxOpenEvent())
	g.POST(wxOpenEvent())

	g.GET(preCode())
	g.GET(preLogin())
}

const tk = "duduchewang"
const aesKey = "KAK4o6mos89GSHSoauSHy6Yc7qTXy86bRhwjBMxctw6"

func preLogin() (path string, handle dotweb.HttpHandle) {
	path = "/pre-login"
	handle = func(ctx dotweb.Context) error {
		callBack := util.GetRequestHost(ctx)
		pre := open.GetPreAuthCode()
		uri := fmt.Sprintf("https://mp.weixin.qq.com/cgi-bin/componentloginpage?"+
			"action=bindcomponent"+
			"&component_appid=%s"+
			"&pre_auth_code=%s"+
			"&redirect_uri=%s"+
			//"&auth_type=1"+
			"&biz_appid=wxbff164af4a8b6ef6"+
			//"#wechat_redirect" +
			"", config.Cfg().WxOpen.AppId, template.URLQueryEscaper(pre.PreAuthCode), template.URLQueryEscaper(callBack))
		//code, _ := qr.Encode(uri, qr.L, qr.Unicode)
		//code, _ = barcode.Scale(code, 300, 300)
		//ctx.Response().SetContentType("image/png")
		rlog.Warn(uri)
		return ctx.Redirect(301, uri)
	}
	return
}

func preCode() (path string, handle dotweb.HttpHandle) {
	path = "/pre-code"
	handle = func(ctx dotweb.Context) error {
		callBack := util.GetRequestHost(ctx) + "/"
		pre := open.GetPreAuthCode()
		uri := fmt.Sprintf("https://mp.weixin.qq.com/safe/bindcomponent?"+
			"action=bindcomponent"+
			"&no_scan=0"+
			"&component_appid=%s"+
			"&pre_auth_code=%s"+
			"&redirect_uri=%s"+
			//"&auth_type=1"+
			"&biz_appid=wxbff164af4a8b6ef6"+
			"#wechat_redirect"+
			"", config.Cfg().WxOpen.AppId, template.URLQueryEscaper(pre.PreAuthCode), template.URLQueryEscaper(callBack))
		//code, _ := qr.Encode(uri, qr.L, qr.Unicode)
		//code, _ = barcode.Scale(code, 300, 300)
		//ctx.Response().SetContentType("image/png")
		//rlog.Warn(uri)
		//return png.Encode(ctx.Response().Writer(), code)
		return ctx.Redirect(301, uri)
	}
	return
}

func wxOpenEvent() (path string, handle dotweb.HttpHandle) {
	path = "/wx/event"
	handle = func(ctx dotweb.Context) error {
		var param wx.NotifyParam
		_ = util.BindUrlParams(ctx.Request().Url(), &param)

		if param.CheckSign(tk) {
			var event open.NotifyEventBase
			var data []byte
			data = ctx.Request().PostBody()
			if param.IsAESEncryptType() {
				var encrypt struct {
					AppId   string
					Encrypt string
				}
				_ = xml.Unmarshal(ctx.Request().PostBody(), &encrypt)
				data, _ = token.BindDecrypt(encrypt.AppId, aesKey, encrypt.Encrypt, &event)
			} else {
				_ = xml.Unmarshal(ctx.Request().PostBody(), &event)
			}
			//rlog.WarnF("%s", data)
			//rlog.WarnF("%+v", event)

			if event.InfoType == open.TypeComponentVerifyTicket {
				var ticket open.NotifyVerifyTicket
				_ = xml.Unmarshal(data, &ticket)
				rlog.WarnF("%+v", ticket)
				if ticket.ComponentVerifyTicket == "" {
					rlog.WarnF("%s", data)
				}
				open.SetCurrentNotifyVerifyTicket(ticket)
			} else if event.InfoType == open.TypeAuthorized {
				var auth open.NotifyAuthorization
				_ = xml.Unmarshal(data, &auth)
				rlog.WarnF("%+v", auth)
				authCode, _ := open.QueryAuth(auth.AuthorizationCode)
				rlog.WarnF("%+v", authCode)
			} else {
				var auth open.NotifyAuthorization
				_ = xml.Unmarshal(data, &auth)
				rlog.WarnF("%+v", auth)
			}
		}
		return ctx.WriteString("success")
	}
	return
}

func wxNotifyMsg() (path string, handle dotweb.HttpHandle) {
	path = "/wx/msg/:appId/callback"
	handle = func(ctx dotweb.Context) error {
		appId := ctx.GetRouterName("appId")
		var msg wx.NotifyMsg
		var param wx.NotifyParam
		_ = util.BindUrlParams(ctx.Request().Url(), &param)
		if ctx.Request().Method == "GET" {
			rlog.Warn(ctx.Request().Url())
			rlog.WarnF("%+v", param)
			if param.CheckSign(tk) {
				return ctx.WriteString(ctx.QueryString("echostr"))
			} else {
				return ctx.WriteString("")
			}
		}
		rlog.Warn(ctx.Request().Url())
		rlog.WarnF("%+v", param)
		if param.CheckSign(tk) {
			if param.IsAESEncryptType() {
				var encrypt struct {
					ToUserName string
					Encrypt    string
				}
				_ = xml.Unmarshal(ctx.Request().PostBody(), &encrypt)
				_, _ = token.BindDecrypt(encrypt.ToUserName, aesKey, encrypt.Encrypt, &msg)
			} else {
				_ = xml.Unmarshal(ctx.Request().PostBody(), &msg)
			}
			rlog.Warn(appId)
			rlog.WarnF("%+v", msg)
		}
		return nil
	}
	return
}
