package util

import (
	"fmt"
	"github.com/devfeel/dotweb"
	"strings"
)

func GetRequestHost(ctx dotweb.Context) string {
	return fmt.Sprintf("%v://%v", strings.Split(strings.ToLower(ctx.Request().Proto), "/")[0], ctx.Request().Host)
}

func GetRealClientIP(ctx dotweb.Context) string {
	forward := ctx.Request().Header.Get("X-Forwarded-For")
	if forward != "" {
		return strings.Split(forward, ",")[0]
	} else {
		return ctx.RemoteIP()
	}
}
