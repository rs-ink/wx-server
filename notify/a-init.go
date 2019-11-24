package notify

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/binary"
	"encoding/xml"
	"errors"
	"fmt"
	"github.com/devfeel/dotweb"
	"io"
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
	Encrypt      string
}
type WxNotifyParam struct {
	Signature    string                `json:"signature"`
	Timestamp    rtype.RDateTimeSecond `json:"timestamp"`
	Nonce        string                `jsn:"nonce"`
	EncrtypeType string                `json:"encrypt_type"`
	MsgSignature string                `json:"msg_signature"`
}

const tk = "sdfasdf"
const aesKey = "lZ5ekIXWlSp9xa2qj8hzTNwUeJB1ZgvsKHttGU1RyOp"

func wxNotify() (path string, handle dotweb.HttpHandle) {
	path = "/wx"
	handle = func(ctx dotweb.Context) error {
		var msg WxNotifyMsg
		var param WxNotifyParam

		_ = ctx.Bind(&param)
		data := ctx.Request().PostBody()

		_ = xml.Unmarshal(data, &msg)
		rlog.WarnF("%+v", msg)
		rlog.WarnF("%+v", param)

		en, _ := base64.StdEncoding.DecodeString(msg.Encrypt)
		key, _ := base64.StdEncoding.DecodeString(aesKey + "=")
		dd, err := aesDecrypt(en, key)
		if err == nil {
			msg.Encrypt = ""
			_ = parseEncryptTextRequestBody(msg.ToUserName, dd, &msg)
			rlog.WarnF("%+v", msg)
		} else {
			rlog.Error(err)
		}

		return ctx.WriteBlob("", data)
	}
	return
}
func DecryptAES(appId, ev, key string) (result string) {
	data, _ := base64.StdEncoding.DecodeString(ev)
	aesKey, _ := base64.StdEncoding.DecodeString(key + "=")
	dd, err := aesDecrypt(data, aesKey)
	if err == nil {
		parseEncryptTextRequestBody(appId, dd, nil)
	} else {
		rlog.Error(err)
	}
	return
}
func aesDecrypt(cipherData []byte, aesKey []byte) ([]byte, error) {
	k := len(aesKey) //PKCS#7
	if len(cipherData)%k != 0 {
		return nil, errors.New("crypto/cipher: ciphertext size is not multiple of aes key length")
	}
	block, err := aes.NewCipher(aesKey)
	if err != nil {
		return nil, err
	}
	iv := make([]byte, aes.BlockSize)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}
	blockMode := cipher.NewCBCDecrypter(block, iv)
	plainData := make([]byte, len(cipherData))
	blockMode.CryptBlocks(plainData, cipherData)
	return plainData, nil
}
func parseEncryptTextRequestBody(appId string, plainText []byte, bind interface{}) error {
	fmt.Println(string(plainText))

	// Read length
	buf := bytes.NewBuffer(plainText[16:20])
	var length int32
	_ = binary.Read(buf, binary.BigEndian, &length)
	rlog.Warn(string(plainText[20 : 20+length]))

	// appID validation
	appIDstart := 20 + length
	id := plainText[appIDstart : int(appIDstart)+len(appId)]
	if !validateAppId(id) {
		rlog.Warn("Wechat Service: appid is invalid!")
		return errors.New("AppId is invalid")
	}

	// xml Decoding
	_ = xml.Unmarshal(plainText[20:20+length], bind)
	return nil
}

func validateAppId(appId []byte) bool {
	return true
}
