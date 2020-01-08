package domain

type OwnerBase struct {
	OwnerId   int    `json:"ownerId"`
	OwnerName string `json:"ownerName"`
	Phone     string `json:"phone"`
}
