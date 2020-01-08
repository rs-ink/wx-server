package domain

import (
	"time"
	"wx-server/rtype/base"
)

//活动订单
type ActivityOrder struct {
	base.OrderID
	ActivityId  int `json:"activityId"`
	CustomerId  int `json:"customerId"`
	ShopId      int `json:"shopId" info:"店铺Id"`
	OwnerId     int `json:"ownerId" info:"店铺管理员Id"`
	AgentId     int `json:"agentId" info:"代理商Id"`
	FromCusId   int `json:"fromCusId"`
	FromShareId int `json:"fromShareId"`

	CarPlate     string `json:"carPlate"`
	AccountName  string `json:"accountName"`
	AccountPhone string `json:"accountPhone"`

	base.LonAndLatMotion
	ConsumeState      int       `json:"consumeState"`
	ConsumeUpdateTime time.Time `json:"consumeUpdateTime"`
}

type ActivityPayInfo struct {
	PayCode       string    `json:"payCode"`
	PayState      int       `json:"payState"`
	PayUpdateTime time.Time `json:"payUpdateTime"`
}

//活动订单历史
type ActivityOrderHistory struct {
	ActivityOrder
	ActivityName string `json:"activityName"`
	ActivityData string `json:"activityData"`
	Price        int    `json:"price"`
}
