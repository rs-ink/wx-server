package notify

import (
	"github.com/devfeel/dotweb"
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
	return
}
