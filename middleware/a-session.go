package middleware

import (
	"encoding/gob"
	"github.com/devfeel/dotweb"
	"wx-server/rtype/wx"
)

func init() {
	gob.Register(AppSession{})
}

type AppSession struct {
	WxSession  wx.Session
	WxUserInfo wx.UserInfo
	ctx        dotweb.Context
}

func NewAppSession() *AppSession {
	return &AppSession{}
}

func (appSession *AppSession) Update(ctx ...dotweb.Context) (err error) {
	if len(ctx) > 0 {
		appSession.ctx = ctx[0]
	}
	if appSession.ctx != nil {
		err = appSession.ctx.Session().Set("app-session", *appSession)
	}
	return
}

func GetAppSession(ctx dotweb.Context) (appSession *AppSession) {
	a := ctx.Session().Get("app-session")
	appSession = NewAppSession()
	if a != nil {
		*appSession, _ = a.(AppSession)
		appSession.ctx = ctx
	}
	return
}
