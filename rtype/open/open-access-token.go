package open

import (
	"fmt"
	"wx-server/config"
	"wx-server/rlog"
)

type (
	WxAccessToken struct {
		AuthorizationInfo struct {
			AuthorizationInfo  string `json:"authorization_info"`
			AuthorAccessToken  string `json:"authorizer_access_token"`
			ExpiresIn          int64  `json:"expires_in"`
			AuthorRefreshToken string `json:"authorizer_refresh_token"`
			FuncInfo           []struct {
				FuncScopeCategory struct {
					Id int `json:"id"`
				} `json:"funcscope_category"`
			} `json:"func_info"`
		} `json:"authorization_info"`
	}
)

func QueryAuth(authCodeValue string) (token WxAccessToken, err error) {
	data := []byte(fmt.Sprintf(`
			"component_appid":"%v" ,
			"authorization_code": "%v"
		`, config.Cfg().WxOpen.AppId, authCodeValue))
	result, err := PostWx(fmt.Sprintf("%v/component/api_query_auth?"+
		"component_access_token=%v", WxCgiApi, GetComponentAccessToken().ComponentAccessToken), data, &token)
	rlog.WarnF("%s", result)
	rlog.WarnF("%+v", token)
	return
}
