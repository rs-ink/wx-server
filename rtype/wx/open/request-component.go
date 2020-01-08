package open

import (
	"wx-server/config"
	"wx-server/rtype/base"
)

type OffsetRequest struct {
	Offset int `json:"offset"`
	Count  int `json:"count"`
}

func ParseOffsetRequest(request base.PageRequest) OffsetRequest {
	request.CheckMaxPageRequest(500)
	return OffsetRequest{
		Offset: request.Start,
		Count:  request.Limit,
	}
}

type RequestComponentBase struct {
	ComponentAccessToken string `json:"component_access_token"`
	ComponentAppId       string `json:"component_appid"`
}

func GetRequestComponentBase() RequestComponentBase {
	return RequestComponentBase{
		ComponentAccessToken: GetComponentAccessToken().ComponentAccessToken,
		ComponentAppId:       config.Cfg().WxOpen.AppId,
	}
}
