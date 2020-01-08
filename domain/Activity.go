package domain

import "time"

//活动
type Activity struct {
	ActivityId   int           `json:"activityId"`
	ActivityName string        `json:"activityName"`
	ActivityData string        `json:"activityData"`
	ShopId       int           `json:"shopId"`
	AgentId      int           `json:"agentId"`
	ShopAddress  []ShopAddress `json:"shopAddress"`
	ActivityConfig
	ActivityRedPackConfig
	ActivityTechnicalSupport
}

//活动页面技术支持
type ActivityTechnicalSupport struct {
	Name      string `json:"name"`
	QrCodeUrl string `json:"qrCodeUrl"`
}

type ActivityViewConfig struct {
}

//活动基本配置
type ActivityConfig struct {
	Price     int       `json:"price"`
	StartTime time.Time `json:"startTime"`
	EndTime   time.Time `json:"endTime"`
}

//活动红包配置
type ActivityRedPackConfig struct {
	RedPackAmount int `json:"redPackAmount"`
}
