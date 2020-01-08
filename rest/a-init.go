package rest

import (
	"github.com/devfeel/dotweb"
	"wx-server/database"
	"xorm.io/xorm"
)

var db *xorm.Engine

func init() {
	db = database.Engine()
}

type Handle func() (path string, handle dotweb.HttpHandle)
type Method string
type Methods []Method

const (
	GET  Method = "GET"
	POST Method = "POST"
)

type route struct {
	method     Methods
	group      dotweb.Group
	handle     Handle
	middleware []dotweb.Middleware
}

type routes []route

//TODO /rest
func InitRestRouter(group dotweb.Group) {
	initWxmRest(group)
	initRestOssRouter(group)
	initRoutes(group.Group("/open"), wxOpenRoutes)
}

func initRoutes(group dotweb.Group, routes []route) {
	for _, r := range routes {
		for _, m := range r.method {
			var node dotweb.RouterNode
			switch m {
			case GET:
				node = initRouteGet(group, r)
			case POST:
				node = initRoutePost(group, r)
			}
			if r.middleware != nil && len(r.middleware) > 0 {
				node.Use(r.middleware...)
			}
		}
	}
}

func initRouteGet(group dotweb.Group, route route) dotweb.RouterNode {
	if route.group != nil {
		return route.group.GET(route.handle())
	} else {
		return group.GET(route.handle())
	}
}

func initRoutePost(group dotweb.Group, route route) dotweb.RouterNode {
	if route.group != nil {
		return route.group.POST(route.handle())
	} else {
		return group.POST(route.handle())
	}
}
