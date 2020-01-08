package rest

import (
	"github.com/devfeel/dotweb"
	"wx-server/rtype"
	"wx-server/rtype/base"
	"wx-server/rtype/wx/open"
)

var wxOpenRoutes = routes{
	route{method: Methods{GET}, handle: list},
}

func list() (path string, handle dotweb.HttpHandle) {
	path = "/list"
	handle = func(ctx dotweb.Context) error {
		var param struct {
			base.PageRequest
		}
		_ = ctx.Bind(&param)
		param.Limit = 20
		open.GetAuthorizerList(param.PageRequest)
		return ctx.WriteJson(rtype.Success())
	}
	return
}
