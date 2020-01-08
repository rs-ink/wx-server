package domain

import "wx-server/rtype/base"

type CardCategory struct {
	CardCategoryId int            `json:"cardCategoryId"`
	State          base.StateBase `json:"state"`
}

type CardCategorySku struct {
}
