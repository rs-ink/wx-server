package middleware

import (
	"encoding/xml"
	"github.com/devfeel/dotweb"
	"wx-server/rlog"
	"wx-server/rtype/wx"
	"wx-server/token"
	"wx-server/util"
)

const tk = "duduchewang"
const aesKey = "KAK4o6mos89GSHSoauSHy6Yc7qTXy86bRhwjBMxctw6"

func NewWxNotifyMsgMiddleware() *WxNotifyMsgMiddleware {
	return &WxNotifyMsgMiddleware{}
}

type WxNotifyMsgMiddleware struct {
	dotweb.BaseMiddleware
}

func (wxm WxNotifyMsgMiddleware) Handle(ctx dotweb.Context) error {
	var msg wx.NotifyMsg
	var param wx.NotifyParam
	_ = util.BindUrlParams(ctx.Request().Url(), &param)
	if ctx.Request().Method == "GET" {
		_ = ctx.WriteString(param.EchoStr)
		ctx.End()
		return nil
	}
	rlog.WarnF("%+v", msg)
	rlog.WarnF("%+v", param)
	rlog.Warn(param.CheckSign(tk))
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
	rlog.WarnF("%+v", msg)
	return nil
}
