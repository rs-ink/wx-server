package rest

import (
	"github.com/devfeel/dotweb"
	"wx-server/rlog"
	"wx-server/rtype"
)

func initWxRest(group dotweb.Group) {
	group.GET(code())
}

func code() (path string, handle dotweb.HttpHandle) {
	path = "/code"
	handle = func(ctx dotweb.Context) error {
		rlog.Warn("123123")
		return ctx.WriteJson(rtype.Success())
	}
	return
}
