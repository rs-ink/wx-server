package domain

import (
	"wx-server/rtype"
	"wx-server/rtype/base"
)

//活动阅读
type ActivityRead struct {
	ActivityReadId int               `json:"activityReadId" info:"阅读Id"`
	ActivityId     int               `json:"activityId" info:"活动Id"`
	CustomerId     int               `json:"customerId" info:"客户Id"`
	ShopId         int               `json:"shopId" info:"店铺Id"`
	OwnerId        int               `json:"ownerId" info:"店铺管理员Id"`
	AgentId        int               `json:"agentId" info:"代理商Id"`
	Duration       int               `json:"duration" info:"阅读时长"`
	FromCusId      int               `json:"fromCusId" info:"来源客户Id"`
	FromShareId    int               `json:"fromShareId" info:"来源分享Id"`
	WxChannel      rtype.FromChannel `json:"wxChannel" info:"来源通道"`
	base.LonAndLatMotion
}
