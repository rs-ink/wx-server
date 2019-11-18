package rtype

import (
	"wx-server/rlog"
	"wx-server/rtype/wx"
)

func init()  {
	err := db.CreateTables(&Customer{})
	rlog.CheckShowError(err)
}

type Customer struct {
	ID `xorm:"extends"`
	BaseTime `xorm:"extends"`
	wx.BaseCustomer `xorm:"extends"`
	wx.BaseOther `xorm:"extends"`
}

func (c Customer) TableName() string {
	return "tb_customer"
}
