package open

import (
	"encoding/json"
	"fmt"
	"wx-server/rlog"
	"wx-server/rtype/base"
)

func GetAuthorizerList(request base.PageRequest) {
	var param struct {
		RequestComponentBase
		OffsetRequest
	}
	param.RequestComponentBase = GetRequestComponentBase()
	param.OffsetRequest = ParseOffsetRequest(request)
	data, _ := json.Marshal(param)
	data, err := PostWx(fmt.Sprintf("%v/component/api_get_authorizer_list?component_access_token=%v", WxCgiApi, GetComponentAccessToken().ComponentAccessToken), data, nil)
	if err != nil {
		rlog.CheckShowError(err)
	} else {
		rlog.WarnF("%s", data)
	}
}
