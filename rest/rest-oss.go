package rest

import (
	"github.com/devfeel/dotweb"
	"wx-server/rtype"
	"wx-server/rtype/oss"
)

func initRestOssRouter(group dotweb.Group) {
	g := group.Group("/oss")

	g.GET(temp())

	g.POST(tempWx())
	g.GET(tempWx())
}

func temp() (path string, handle dotweb.HttpHandle) {
	path = "/temp"
	handle = func(ctx dotweb.Context) error {
		token, filePath := oss.GetOssTempToken()
		return ctx.WriteJson(rtype.Success().SetData(token).SetExt(filePath))
	}
	return
}
func tempWx() (path string, handle dotweb.HttpHandle) {
	path = "/temp-wx"
	handle = func(ctx dotweb.Context) error {
		return ctx.WriteJson(rtype.Success().SetData(oss.CreateWxOssSign()))
	}
	return
}
