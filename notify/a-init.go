package notify

import (
	"encoding/xml"
	"github.com/devfeel/dotweb"
	"wx-server/rlog"
	"wx-server/rtype"
)

func InitNotifyRouter(g dotweb.Group) {
	g.POST(wxNotify())
}

/**
<xml>
  <ToUserName><![CDATA[toUser]]></ToUserName>
  <FromUserName><![CDATA[fromUser]]></FromUserName>
  <CreateTime>1348831860</CreateTime>
  <MsgType><![CDATA[text]]></MsgType>
  <Content><![CDATA[this is a test]]></Content>
  <MsgId>1234567890123456</MsgId>
</xml>
*/
type WxNotifyMsg struct {
	ToUserName   string
	FromUserName string
	CreateTime   rtype.RDateTimeSecond
	MsgType      string
	Content      string
	MsgId        int64
}

func wxNotify() (path string, handle dotweb.HttpHandle) {
	path = "/wx"
	handle = func(ctx dotweb.Context) error {
		var msg WxNotifyMsg
		data := ctx.Request().PostBody()
		rlog.WarnF("%s", data)
		_ = xml.Unmarshal(data, &msg)
		rlog.Warn(msg)

		return ctx.WriteBlob("", data)
	}
	return
}
