package middleware

import (
	"github.com/devfeel/dotweb"
	"wx-server/rlog"
	"wx-server/rtype"
	"wx-server/token"
	"wx-server/util"
)

func NewWxLoginMiddleware() *WxLoginMiddleware {
	return &WxLoginMiddleware{}
}

type WxLoginMiddleware struct {
	dotweb.BaseMiddleware
}

func (wxm WxLoginMiddleware) Handle(ctx dotweb.Context) error {
	//TODO 用于兼容多公众号授权
	conf := rtype.GetWxsConfig()
	appSession := GetAppSession(ctx)
	rlog.WarnF("%+v", appSession.WxSession)
	rlog.WarnF("%+v", appSession.WxUserInfo)
	if appSession.WxSession.OpenId == "" {
		code := ctx.QueryString("code")
		if code == "" {
			return util.RedirectToForWxCode(ctx, conf.AppId)
		} else {
			session, _ := token.GetWxSession(code, conf)
			//40029 无效的 oauth_code
			//40163	oauth_code已使用
			if session.ErrCode == 40163 || session.ErrCode == 40029 {
				return util.RedirectToForWxCode(ctx, conf.AppId)
			}
			appSession.WxSession = session
			rlog.WarnF("%+v", session)
			wxInfo, _ := token.GetWxInfo(conf, session)
			rlog.WarnF("%+v", wxInfo)
			appSession.WxUserInfo = wxInfo
			_ = appSession.Update()
		}
	}
	return wxm.Next(ctx)
}
