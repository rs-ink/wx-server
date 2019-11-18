package rest

import (
	"github.com/devfeel/dotweb"
	"wx-server/middleware"
)

func InitRestRouter(group dotweb.Group) {
	initWxRest(group.Group("/wx", middleware.NewWxLoginMiddleware()))
}
