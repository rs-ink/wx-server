package domain

type Operator struct {
	OperatorId int    `json:"operatorId"`
	NickName   string `json:"nickName"`
	Avatar     string `json:"avatar"`
	Phone      string `json:"phone"`
}
