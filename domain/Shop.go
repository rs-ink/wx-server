package domain

type Shop struct {
	ShopId        int           `json:"shopId"`
	AgentId       int           `json:"agentId"`
	ShopName      string        `json:"shopName"`
	ShortShopName string        `json:"shortShopName"`
	Address       []ShopAddress `json:"address"`
	OwnerBase
}
