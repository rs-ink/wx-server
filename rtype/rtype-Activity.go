package rtype

func init() {
	_ = db.CreateTables(ActivityEntity{})
}

type ActivityState int

type ActivityEntity struct {
	ID
	BaseTime
	ShopId       int    `json:"shopId"`
	AgentId      int    `json:"agentId"`
	OwnerId      int    `json:"ownerId"`
	ActivityName string `json:"activityName"`
	ActivityData string `json:"activityData"`
	Price        int    `json:"price"`
	ActivityShareConfig
}

func (a ActivityEntity) TableName() string {
	return "tb_activity"
}
