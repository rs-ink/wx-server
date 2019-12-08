package wx

import (
	"crypto/sha1"
	"fmt"
	"sort"
	"strings"
	"wx-server/rtype/base"
)

type NotifyMsg struct {
	ToUserName   string
	FromUserName string
	CreateTime   base.RDateTimeSecond
	MsgType      string
	Content      string
	MsgId        int64
}

//WxNotifyParam
type NotifyParam struct {
	Signature    string               `json:"signature"`
	Timestamp    base.RDateTimeSecond `json:"timestamp"`
	Nonce        int64                `jsn:"nonce"`
	EncryptType  string               `json:"encrypt_type"`
	MsgSignature string
	EchoStr      string `json:"echostr"`
}

func (param NotifyParam) IsAESEncryptType() bool {
	return param.EncryptType == "aes"
}
func (param NotifyParam) CheckSign(tk string) bool {
	str := make([]string, 3)
	str = append(str, fmt.Sprint(param.Nonce))
	str = append(str, fmt.Sprint(param.Timestamp.Time().Unix()))
	str = append(str, tk)
	sort.Strings(str)
	sha := sha1.New()
	sha.Write([]byte(strings.Join(str, "")))
	//rlog.Warn(fmt.Sprintf("%+x", sha.Sum(nil)))
	//rlog.Warn(param.Signature)
	return fmt.Sprintf("%+x", sha.Sum(nil)) == param.Signature
}
