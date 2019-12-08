package router

import (
	"github.com/devfeel/dotweb"
	"github.com/rs-ink/rslog"
	"reflect"
	"strings"
	"wx-server/rlog"
	"wx-server/rtype"
	"wx-server/util"
)

var wxFilterMap map[string]string

func init() {
	wxFilterMap = make(map[string]string)
	wxFilterMap["MP_verify_Mg3NtQxQFCNXqIAZ.txt"] = "Mg3NtQxQFCNXqIAZ"
	wxFilterMap["6878899836.txt"] = "1d803a5404c283bb0e73d9d5884e19a9"
}

func wxFilter() *dotweb.HttpModule {
	return &dotweb.HttpModule{
		Name: "wx-filter",
		OnBeginRequest: func(ctx dotweb.Context) {
			uri := ctx.Request().Url()
			if v, ok := wxFilterMap[uri[1:]]; ok {
				_ = ctx.WriteString(v)
				ctx.End()
			}
		},
	}
}

//func wxNotifyFilter() *dotweb.HttpModule {
//	mid := middleware.NewWxNotifyMsgMiddleware()
//	return &dotweb.HttpModule{
//		Name:"wx-notify1",
//		OnBeginRequest: func(ctx dotweb.Context) {
//			_ = mid.Handle(ctx)
//		},
//	}
//}

func NotFoundHandler(ctx dotweb.Context) {
	//if !ctx.IsEnd() {
	rlog.Info("===================================")
	rlog.Info("来源IP：", util.GetRealClientIP(ctx))
	rlog.Info(ctx.Request().QueryHeader("X-Forwarded-For"))
	rlog.Info(ctx.Request().Method, ctx.Request().Url(), " ContentType:", ctx.Request().ContentType(), "  Referer:", ctx.Request().Referer())
	rlog.Info(string(ctx.Request().PostBody()))
	if strings.HasPrefix(ctx.Request().Url(), "/rest") {
		_ = ctx.WriteJson(rtype.ErrCodeSystemError.Result().SetMsg("接口不存在"))
	} else {
		_ = ctx.Redirect(301, "/")
	}
	rlog.Info("===================================")
	ctx.End()
	//}
}

func MethodNotAllowedHandler(ctx dotweb.Context) {
	if !ctx.IsEnd() {
		_ = ctx.WriteJson(rtype.ErrCodeSystemError.Result().SetMsg("Method Not Allowed"))
		ctx.End()
	}
}

func ExceptionHanlde(ctx dotweb.Context, err error) {
	if !ctx.IsEnd() {
		if err != nil {
			if reflect.TypeOf(err) == reflect.TypeOf(&rtype.Result{}) {
				result := err.(*rtype.Result)
				rslog.OutPc(result.GetPcInfo(), rslog.LevelERROR, result)
				_ = ctx.WriteJson(*result)
			} else {
				_ = ctx.WriteJson(rtype.ErrCodeSystemError.Result().SetMsg(err.Error()))
			}
		} else {
			rlog.Error("系统未知异常")
			_ = ctx.WriteJson(rtype.ErrCodeSystemUndefineError.Result().SetMsg("系统未知异常"))
		}
		ctx.End()
	}
}
