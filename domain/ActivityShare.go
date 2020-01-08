package domain

import (
	"wx-server/rtype"
	"wx-server/rtype/base"
)

//活动分享
type ActivityShare struct {
	ActivityShareId int               `json:"activityShareId" info:"分享Id"`
	ActivityId      int               `json:"activityId" info:"活动Id"`
	CustomerId      int               `json:"customerId" info:"customerId"`
	ShopId          int               `json:"shopId" info:"店铺Id"`
	OwnerId         int               `json:"ownerId" info:"店铺管理员Id"`
	AgentId         int               `json:"agentId" info:"代理商Id"`
	FromCusId       int               `json:"fromCusId" info:"来源客户Id"`
	FromShareId     int               `json:"fromShareId" info:"来源分享Id"`
	WxChannel       rtype.FromChannel `json:"wxChannel" info:"分享通道"`
	base.LonAndLatMotion
}
