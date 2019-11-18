package middleware

import (
	"github.com/devfeel/dotweb"
	"wx-server/rlog"
)

func NewWxLoginMiddleware() *WxLoginMiddleware {
	return &WxLoginMiddleware{}
}

type WxLoginMiddleware struct {
	dotweb.BaseMiddleware
}

func (wxm WxLoginMiddleware) Handle(ctx dotweb.Context) error {
	rlog.Info("implement me")

	return wxm.Next(ctx)
}
