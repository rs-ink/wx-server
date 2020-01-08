package open

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"sync"
	"time"
	"wx-server/config"
	"wx-server/rlog"
	"wx-server/rtype/base"
	"wx-server/rtype/wx"
)

const WxApiServer = "https://api.weixin.qq.com"
const WxCgiApi = WxApiServer + "/cgi-bin"

type InfoType string

func (info InfoType) String() string {
	switch info {
	case TypeComponentVerifyTicket:
		return "验证票据"
	case TypeUnAuthorized:
		return "取消授权"
	case TypeUpdateAuthorized:
		return "更新授权"
	case TypeAuthorized:
		return "授权成功"
	default:
		return string(info)
	}
}

const (
	TypeComponentVerifyTicket InfoType = "component_verify_ticket"
	TypeUnAuthorized                   = "unauthorized"
	TypeUpdateAuthorized               = "updateauthorized"
	TypeAuthorized                     = "authorized"
)

type NotifyVerifyTicket struct {
	NotifyEventBase
	ComponentVerifyTicket string `json:"componentVerifyTicket" xml:"ComponentVerifyTicket"`
}
type NotifyAuthorization struct {
	NotifyEventBase
	AuthorizerAppId              string               `xml:"AuthorizerAppid"`
	AuthorizationCode            string               `xml:"AuthorizationCode"`
	AuthorizationCodeExpiredTime base.RDateTimeSecond `xml:"AuthorizationCodeExpiredTime"`
	PreAuthCode                  string               `xml:"PreAuthCode"`
}
type NotifyEventBase struct {
	AppId      string               `json:"appId" xml:"AppId"`
	CreateTime base.RDateTimeSecond `json:"createTime" xml:"CreateTime"`
	InfoType   InfoType             `json:"infoType" xml:"InfoType"`
}

var NotifyVerifyTicketMutex sync.Mutex

var currentOpenNotifyVerifyTicket NotifyVerifyTicket

func init() {
	data, err := ioutil.ReadFile("./currentOpenNotifyVerifyTicket")
	if err == nil {
		var tick NotifyVerifyTicket
		_ = json.Unmarshal(data, &tick)
		if tick.ComponentVerifyTicket != "" {
			SetCurrentNotifyVerifyTicket(tick)
		}
	}
}

func SetCurrentNotifyVerifyTicket(ticket NotifyVerifyTicket) {
	NotifyVerifyTicketMutex.Lock()
	defer NotifyVerifyTicketMutex.Unlock()
	currentOpenNotifyVerifyTicket = ticket
	defer func() {
		data, _ := json.Marshal(currentOpenNotifyVerifyTicket)
		_ = ioutil.WriteFile("./currentOpenNotifyVerifyTicket", data, os.ModePerm)
	}()
}
func GetCurrentNotifyVerifyTicket() NotifyVerifyTicket {
	NotifyVerifyTicketMutex.Lock()
	defer NotifyVerifyTicketMutex.Unlock()
	rlog.WarnF("%+v", currentOpenNotifyVerifyTicket)
	return currentOpenNotifyVerifyTicket
}

type AccessToken struct {
	wx.Error
	ComponentAccessToken string    `json:"component_access_token"`
	ExpiresIn            int64     `json:"expires_in"`
	CreateTime           time.Time `json:"createTime"`
}

var accessTokenMutex sync.Mutex

var currentAccessToken AccessToken

//获取第三方平台AccessToken
func GetComponentAccessToken() AccessToken {
	accessTokenMutex.Lock()
	defer accessTokenMutex.Unlock()
	if currentAccessToken.ComponentAccessToken == "" || time.Now().After(currentAccessToken.CreateTime.Add(time.Duration(currentAccessToken.ExpiresIn)*time.Second)) {
		data := []byte(fmt.Sprintf(`{
 				 "component_appid":  "%v" ,
     			 "component_appsecret":  "%v",
				 "component_verify_ticket": "%v"
				}`, config.Cfg().WxOpen.AppId, config.Cfg().WxOpen.AppSecret, GetCurrentNotifyVerifyTicket().ComponentVerifyTicket))
		var token AccessToken
		rlog.WarnF("%s", data)
		data, err := PostWx(fmt.Sprintf("%s/component/api_component_token", WxCgiApi), data, &token)
		if err == nil {
			token.CreateTime = time.Now()
			currentAccessToken = token
		} else {
			rlog.ErrorF("%s", data)
		}
		if currentAccessToken.IsError() {
			rlog.WarnF("currentAccessToken:%+v", currentAccessToken)
			rlog.ErrorF("%s", data)
		}
	}
	return currentAccessToken
}

type PreAuthCode struct {
	PreAuthCode string    `json:"pre_auth_code"`
	ExpiresIn   int64     `json:"expires_in"`
	CreateTime  time.Time `json:"createTime"`
}

func GetPreAuthCode() (pre PreAuthCode) {
	body := []byte(`{
  			"component_appid": "` + config.Cfg().WxOpen.AppId + `",
		}`)
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/component/api_create_preauthcode?component_access_token=%s", WxCgiApi, GetComponentAccessToken().ComponentAccessToken), bytes.NewReader(body))
	if err == nil {
		resp, err := client.Do(req)
		if err == nil {
			defer resp.Body.Close()
			data, _ := ioutil.ReadAll(resp.Body)
			_ = json.Unmarshal(data, &pre)
			pre.CreateTime = time.Now()
		}
	}
	return
}
