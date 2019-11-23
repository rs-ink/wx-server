package rest

import (
	"github.com/devfeel/dotweb"
)

func init() {

}

//TODO /rest
func InitRestRouter(group dotweb.Group) {
	initWxmRest(group)
	initRestOssRouter(group)
}
