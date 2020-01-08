package service

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha1"
	"encoding/base64"
	"encoding/binary"
	"encoding/xml"
	"errors"
	"fmt"
	"github.com/axgle/mahonia"
	"io"
	"strconv"
	"time"
	"wx-server/rlog"
	"wx-server/rtype/wx"
	"wx-server/util"
)

func JsSignatureUrl(url string, appId ...string) (sign wx.JsSignature) {
	wxs := wx.GetWxsConfig(appId...)
	sign.AppId = wxs.AppId
	sign.SignType = "SHA1"
	sign.NonceStr = util.CreateRandomString(32)
	sign.TimeStamp = strconv.FormatInt(time.Now().Unix(), 10)
	h := sha1.New()
	h.Write([]byte(fmt.Sprintf("jsapi_ticket=%s&noncestr=%s&timestamp=%s&url=%s", GetTicket(appId...).Ticket, sign.NonceStr, sign.TimeStamp, url)))
	sign.Signature = fmt.Sprintf("%x", h.Sum(nil))
	return
}

func BindDecrypt(appId, key, cipherData string, bind interface{}) (data []byte, err error) {
	en, _ := base64.StdEncoding.DecodeString(cipherData)
	aesKey, _ := base64.StdEncoding.DecodeString(key + "=")
	dd, err := aesDecrypt(en, aesKey)
	if err == nil {
		data, err = parseEncryptTextRequestBody(appId, dd, bind)
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
func parseEncryptTextRequestBody(appId string, plainText []byte, bind interface{}) (data []byte, err error) {
	buf := bytes.NewBuffer(plainText[16:20])
	var length int32
	_ = binary.Read(buf, binary.BigEndian, &length)
	appIdStart := 20 + length
	id := plainText[appIdStart : int(appIdStart)+len(appId)]
	if !validateAppId(id) {
		rlog.Warn("WeChat Service: appId is invalid!")
		err = errors.New("AppId is invalid")
		return
	}
	// xml Decoding
	data = plainText[20 : 20+length]
	_ = xml.Unmarshal(plainText[20:20+length], bind)
	return
}

func validateAppId(appId []byte) bool {
	return true
}

func Decrypt(mis wx.Session, data wx.EncryptedData, charsets ...string) (result string, err error) {
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
				var cryptByte []byte
				cryptByte, err = base64.StdEncoding.DecodeString(data.EncryptedData)
				rlog.CheckShowError(err)
				if err == nil {
					decrypt := cipher.NewCBCDecrypter(aesBlock, iv)
					orig := make([]byte, len(cryptByte))
					decrypt.CryptBlocks(orig, cryptByte)
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
