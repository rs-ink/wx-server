package domain

import "wx-server/rtype/base"

type ShopOperatorBase struct {
	ShopId int    `json:"shopId"`
	Phone  string `json:"phone"`
	Remark string `json:"remark"`
	State  base.StateBase
}
