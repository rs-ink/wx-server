package router

import (
	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
	"github.com/devfeel/dotweb"
	"image/png"
	"wx-server/notify"
	"wx-server/rest"
)

func InitRoute(server *dotweb.HttpServer) {
	server.RegisterModule(wxFilter())
	server.GET("/code.png", func(ctx dotweb.Context) error {
		base64 := ctx.Request().QueryString("url")
		code, _ := qr.Encode(base64, qr.L, qr.Unicode)
		code, _ = barcode.Scale(code, 300, 300)
		ctx.Response().SetContentType("image/png")
		return png.Encode(ctx.Response().Writer(), code)
	})
	server.GET("/", func(ctx dotweb.Context) error {
		return ctx.View("index.html")
	})
	notify.InitNotifyRouter(server.Group("/notify"))
	rest.InitRestRouter(server.Group("/rest"))

	server.GET("/auth", func(ctx dotweb.Context) error {
		return ctx.View("auth.html")
	})
}
