package token

import (
	"encoding/json"
	"fmt"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"net/http"
	"time"
	"wx-server/config"
	"wx-server/rlog"
	"wx-server/rtype"
	"wx-server/rtype/wx"
)

func TransferWxLocation(lat, lng float64) (gpsLat, gpsLng float64) {
	reqUrl := fmt.Sprintf("https://apis.map.qq.com/ws/coord/v1/translate?locations=%v,%v&type=1&key=%v", lat, lng, config.Cfg().WxMap.Key)
	req, _ := http.NewRequest("GET", reqUrl, nil)
	req.Header.Add("Referer", "http://s.rs.ink")
	resp, err := client.Do(req)
	if err == nil {
		defer resp.Body.Close()
		data, err := ioutil.ReadAll(resp.Body)
		if err == nil {
			dd := gjson.ParseBytes(data)
			rlog.Warn(string(data))
			locResult := dd.Get("locations")
			if locResult.Exists() && locResult.IsArray() && len(locResult.Array()) > 0 {
				loc := locResult.Array()[0]
				rlog.WarnF("%+v", loc)
				return loc.Get("lat").Float(), loc.Get("lng").Float()
			}
		}
	}
	return 0, 0
}

func CheckRefreshToken(session *wx.Session,appId ...string)  {
	conf := rtype.GetWxConfig(appId...)
	if session.CreateTime.Unix()+int64(session.Expires) <= time.Now().Unix(){
		resp,err := http.DefaultClient.Get(fmt.Sprintf("https://api.weixin.qq.com/sns/oauth2/refresh_token?appid=%v&grant_type=refresh_token&refresh_token=%v",conf.AppId,session.RefreshToken))
		if err !=nil{
			rlog.Error(err)
		}else{
			defer resp.Body.Close()
			data,_ := ioutil.ReadAll(resp.Body)
			_ = json.Unmarshal(data, session)
		}
	}
}

func GetWxSession(code string,appId ...string) (mis wx.Session, err error) {
	wxc := rtype.GetWxConfig(appId...)
	var resp *http.Response
	if wxc.Type == wx.MiniProgram{
		resp, err = client.Get(fmt.Sprintf("%s/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code", WxSnsApi, wxc.AppId, wxc.AppSecret, code))
	}else{
		resp, err = client.Get(fmt.Sprintf("%s/oauth2/access_token?appid=%s&secret=%s&code=%s&grant_type=authorization_code", WxSnsApi, wxc.AppId, wxc.AppSecret, code))
	}
	if err == nil {
		defer resp.Body.Close()
		assets12, err := ioutil.ReadAll(resp.Body)
		if err == nil {
			rlog.Warn("微信解析返回：：", string(assets12))
			err = json.Unmarshal(assets12, &mis)
			mis.CreateTime=time.Now()
		}
	}
	return
}

//网页体验授权
func GetWxInfoBySnsOpenId(wxs wx.Session) (mis wx.UserInfo, err error) {
	var resp *http.Response
	var data []byte
	resp, err = client.Get(fmt.Sprintf("%v/userinfo?access_token=%s&openid=%s&lang=zh_CN", WxSnsApi, wxs.AccessToken, wxs.OpenId))
	if err == nil {
		defer resp.Body.Close()
		data, err = ioutil.ReadAll(resp.Body)
		rlog.Warn("微信详细信息 SNS：", string(data))
		err = json.Unmarshal(data, &mis)
	}
	return
}

//获取微信用户详细信息
func GetWxInfoByOpenId(openId string,appId ...string) (mis wx.UserInfo, err error) {
	at := GetAccessToken(appId...)
	var resp *http.Response
	var data []byte
	resp, err = client.Get(fmt.Sprintf("%s/user/info?access_token=%v&openid=%v&lang=zh_CN", WxCgiApi, at.AccessToken, openId))
	if err == nil {
		defer resp.Body.Close()
		data, err = ioutil.ReadAll(resp.Body)
		rlog.Warn("微信详细信息：", string(data))
		_ = json.Unmarshal(data, &mis)
	}
	return
}
