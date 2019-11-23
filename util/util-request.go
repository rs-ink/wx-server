package util

import (
	"bytes"
	"fmt"
	"github.com/devfeel/dotweb"
	"html/template"
	"net/url"
	"strings"
	"time"
	"wx-server/rlog"
	"wx-server/rtype"
)

func GetRequestHost(ctx dotweb.Context) string {
	return fmt.Sprintf("%v://%v", strings.Split(GetProto(ctx), "/")[0], ctx.Request().Host)
}

func GetRealClientIP(ctx dotweb.Context) string {
	forward := ctx.Request().Header.Get("X-Forwarded-For")
	if forward != "" {
		return strings.Split(forward, ",")[0]
	} else {
		return ctx.RemoteIP()
	}
}

func GetProto(ctx dotweb.Context) (proto string) {
	proto = ctx.Request().QueryHeader("X-Forwarded-Proto")
	if proto == "" {
		proto = ctx.Request().Proto
	}
	proto = strings.ToLower(proto)
	return
}

func GetAppId(ctx dotweb.Context) string {
	return rtype.GetWxsConfig().AppId
}

func RedirectToForWxCode(ctx dotweb.Context, appId string, values ...url.Values) error {
	var redirectUrl string
	scope := "snsapi_userinfo"
	baseHost := fmt.Sprintf("%v://%v", strings.Split(GetProto(ctx), "/")[0], ctx.Request().Host)
	redirectUrl = ctx.Request().Url()
	var vs0 url.Values
	if strings.Contains(redirectUrl, "?") {
		vs0, _ = url.ParseQuery(strings.Split(redirectUrl, "?")[1])
		vs0.Del("code")
		vs0.Del("state")
		vs0.Del("_")
		rlog.Warn(vs0)
	} else {
		vs0 = url.Values{}
	}
	vs0.Add("_", fmt.Sprint(time.Now().Unix()))
	redirectUrl = strings.Split(redirectUrl, "?")[0] + "?" + vs0.Encode()
	if len(values) > 0 {
		redirectUrl = redirectUrl + "&" + values[0].Encode()
	}
	rlog.Warn(baseHost + redirectUrl)

	var buf bytes.Buffer
	buf.WriteString("https://open.weixin.qq.com/connect/oauth2/authorize?")
	buf.WriteString("appid=")
	buf.WriteString(appId)
	buf.WriteString("&redirect_uri=")
	buf.WriteString(template.URLQueryEscaper(baseHost + redirectUrl))
	buf.WriteString("&response_type=code")
	buf.WriteString("&scope=")
	buf.WriteString(scope)
	buf.WriteString(fmt.Sprintf("&state=%v", time.Now().Nanosecond()))
	buf.WriteString("#wechat_redirect")
	return ctx.Redirect(301, buf.String())
}
