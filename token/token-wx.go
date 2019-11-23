package token

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
	"wx-server/rlog"
	"wx-server/rtype"
	"wx-server/rtype/wx"
	"wx-server/util"
)

var client http.Client

func init() {
	tr := &http.Transport{
		//Proxy:           func(r *http.Request) (*url.URL, error) { return url.Parse("http://127.0.0.1:8888") },
		TLSClientConfig: &tls.Config{
			VerifyPeerCertificate: func(rawCerts [][]byte, verifiedChains [][]*x509.Certificate) error {
				return nil
			},
			InsecureSkipVerify: true,
		},
	}
	jar, err := util.NewCookieJar(nil)
	rlog.CheckShowError(err)
	client = http.Client{
		Transport: tr,
		Jar:       jar,
	}
}

const WxApiServer = "https://api.weixin.qq.com"
const WxCgiApi = WxApiServer + "/cgi-bin"
const WxSnsApi = WxApiServer + "/sns"

var tokenMap sync.Map

func GetAccessToken(appId ...string) wx.AccessToken {
	conf := rtype.GetWxsConfig(appId...)
	tk, ok := tokenMap.Load(conf)
	if !ok || tk.(wx.AccessToken).CreateTime.Unix()+tk.(wx.AccessToken).ExpiresIn <= time.Now().Unix() {
		url := fmt.Sprintf("%s/token?grant_type=client_credential&appid=%s&secret=%s", WxCgiApi, conf.AppId, conf.AppSecret)
		resp, err := client.Get(url)
		if err == nil {
			defer resp.Body.Close()
			data, err := ioutil.ReadAll(resp.Body)
			if err == nil {
				var token wx.AccessToken
				err = json.Unmarshal(data, &token)
				if err == nil {
					if token.ErrCode == 0 {
						token.CreateTime = time.Now()
						tk, ok = tokenMap.LoadOrStore(conf, token)
					} else {
						rlog.Error(token)
					}
				} else {
					panic(err)
				}
			} else {
				panic(err)
			}
		} else {
			panic(err)
		}

	}
	return tk.(wx.AccessToken)
}

var ticketMap sync.Map

func GetTicket(appId ...string) wx.Ticket {
	accessToken := GetAccessToken(appId...)
	conf := rtype.GetWxsConfig(appId...)
	tk, ok := ticketMap.Load(conf)
	tick := tk.(wx.Ticket)
	if !ok || tick.CreateTime.Unix()+tick.ExpiresIn <= time.Now().Unix() {
		url := fmt.Sprintf("%s/ticket/getticket?access_token=%s&type=jsapi", WxCgiApi, accessToken.AccessToken)
		resp, err := client.Get(url)
		if err == nil {
			defer resp.Body.Close()
			data, _ := ioutil.ReadAll(resp.Body)
			rlog.WarnF("ticket 微信返回：%v", string(data))
			err = json.Unmarshal(data, &tick)
			rlog.CheckShowError(err)
			if tick.ErrCode == 0 {
				tick.CreateTime = time.Now()
				tokenMap.LoadOrStore(conf, tick)
				return tick
			}
		} else {
			panic(err)
		}
	}
	return tk.(wx.Ticket)
}
