package token

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"sync"
	"time"
	"wx-server/rlog"
	"wx-server/rtype"
	"wx-server/rtype/wx"
)

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
