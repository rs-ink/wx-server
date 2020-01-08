package rtype

import (
	"time"
	"wx-server/rtype/base"
)

func init() {
}

type CustomerInfo struct {
	CustomerBaseInfo `xorm:"extends"`
	Gender           int       `json:"gender"`
	City             string    `json:"city"`
	SubscribeScene   string    `json:"subscribeScene"`
	Subscribe        int       `json:"subscribe"`
	SubscribeTime    time.Time `json:"subscribeTime"`
	GroupId          int       `json:"groupId"`
	Province         string    `json:"province"`
	Country          string    `json:"country"`
}
type CustomerBaseInfo struct {
	ID                 `xorm:"extends"`
	base.WxBaseAccount `xorm:"extends"`
	NickName           string `json:"nickName"`
	Phone              string `json:"phone"`
	CarPlate           string `json:"carPlate"`
	Remark             string `json:"remark"`
	AvatarUrl          string `json:"avatarUrl"`
	BaseTime           `xorm:"extends"`
}

func (info *CustomerBaseInfo) TableName() string {
	return "tb_customer_info"
}

func GetCustomerBaseInfo(id interface{}) (info CustomerBaseInfo) {
	session := db.NewSession()
	defer session.Close()
	_, _ = session.ID(id).Get(&info)
	return
}
