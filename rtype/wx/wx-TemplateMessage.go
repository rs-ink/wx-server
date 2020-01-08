package wx

type TempMsg struct {
	AccessToken AccessToken `json:"access_token"`
	TouSer      string      `json:"touser"`
	TemplateId  string      `json:"template_id"`
	FormId      string      `json:"form_id"`
	Data        interface{} `json:"data"`
	TempMsgOther
}
type TempMiniProgram struct {
	AppId    string `json:"appid,omitempty"`
	PagePath string `json:"pagepath,omitempty"`
}

type TempMsgOther struct {
	Url         string          `json:"url,omitempty"`
	MiniProgram TempMiniProgram `json:"miniprogram,omitempty"`
}
