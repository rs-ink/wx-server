package domain

type Permission []string

type Role struct {
	Admin     bool `json:"admin"`
	Agent     bool `json:"agent"`
	ShopAdmin bool `json:"shopAdmin"`
}

type Authority struct {
	OperatorId int  `json:"operatorId"`
	Role       Role `json:"role"`
	Permission Permission
	Token      string `json:"service"`
	Agent      Agent  `json:"agent,omitempty"`
	Shop       Shop   `json:"shop,omitempty"`
}
