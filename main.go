package main

import (
	"github.com/devfeel/dotweb"
	"log"
	"wx-server/config"
	"wx-server/rlog"
	"wx-server/router"
	"wx-server/util"
)

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile | log.Lmicroseconds)
}

func main() {
	app := dotweb.New()
	app.SetLogger(rlog.NewRAppLog())
	app.HttpServer.SetEnabledSession(true)
	app.HttpServer.SetEnabledAutoHEAD(true)
	//TODO 跨域处理
	app.Use(util.NewSimpleCROS())
	app.HttpServer.SetEnabledAutoOPTIONS(true)
	app.HttpServer.SetEnabledIgnoreFavicon(true)

	app.HttpServer.Renderer().SetTemplatePath(config.Cfg().Web.Views)
	app.HttpServer.ServerFile("/assets/*", config.Cfg().Web.Assets)
	app.HttpServer.SetEnabledListDir(false)

	app.SetMethodNotAllowedHandle(router.MethodNotAllowedHandler)
	app.SetNotFoundHandle(router.NotFoundHandler)
	app.SetExceptionHandle(router.ExceptionHanlde)
	router.InitRoute(app.HttpServer)

	_ = app.StartServer(config.Cfg().Web.Port)
}
