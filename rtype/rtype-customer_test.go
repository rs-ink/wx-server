package rtype

import (
	"testing"
	"wx-server/rlog"
	"wx-server/rtype/wx"
)

func TestShowCustomer(t *testing.T) {
	_, _ = db.Insert(Customer{
		ID:           ID{},
		BaseTime:     NewBaseTime(),
		BaseCustomer: wx.BaseCustomer{},
		BaseOther:    wx.BaseOther{},
	})
	var cus Customer

	_, _ = db.Get(&cus)
	rlog.InfoF("%+v", cus)
}
