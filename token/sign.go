package token

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha1"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/axgle/mahonia"
	"strconv"
	"time"
	"wx-server/rlog"
	"wx-server/rtype"
	"wx-server/rtype/wx"
	"wx-server/util"
)

func JsSignature(url string, appId ...string) (sign wx.JsSignature) {
	wxs := rtype.GetWxsConfig(appId...)
	sign.AppId = wxs.AppId
	sign.SignType = "SHA1"
	sign.NonceStr = util.CreateRandomString(32)
	sign.TimeStamp = strconv.FormatInt(time.Now().Unix(), 10)
	h := sha1.New()
	h.Write([]byte(fmt.Sprintf("jsapi_ticket=%s&noncestr=%s&timestamp=%s&url=%s", GetTicket(appId...).Ticket, sign.NonceStr, sign.TimeStamp, url)))
	sign.Signature = fmt.Sprintf("%x", h.Sum(nil))
	return
}

func DeCrypt(mis wx.Session, data wx.EncryptedData, charsets ...string) (result string, err error) {
	charset := "utf-8"
	if len(charsets) > 0 {
		charset = charsets[0]
	}
	defer func() {
		if e := recover(); e != nil {
			rlog.Error(e)
			err = errors.New("解析失败")
		}
	}()
	key, err := base64.StdEncoding.DecodeString(mis.SessionKey)
	rlog.CheckShowError(err)
	if err == nil {
		var aesBlock cipher.Block
		aesBlock, err = aes.NewCipher(key)
		rlog.CheckShowError(err)
		if err == nil {
			var iv []byte
			iv, err = base64.StdEncoding.DecodeString(data.Iv)
			rlog.CheckShowError(err)
			if err == nil {
				var crytedByte []byte
				crytedByte, err = base64.StdEncoding.DecodeString(data.EncryptedData)
				rlog.CheckShowError(err)
				if err == nil {
					decrypter := cipher.NewCBCDecrypter(aesBlock, iv)
					orig := make([]byte, len(crytedByte))
					decrypter.CryptBlocks(orig, crytedByte)
					orig = pKCS7UnPadding(orig)
					result = mahonia.NewDecoder(charset).ConvertString(string(orig))
				}
			}
		}
	}
	return
}

func pKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	unPadding := int(origData[length-1])
	return origData[:(length - unPadding)]
}
