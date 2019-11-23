package rtype

import (
	"encoding/gob"
	"wx-server/rlog"
	"wx-server/rtype/wx"
)

func init() {
	gob.Register(CustomerBase{})
	err := db.CreateTables(&Customer{})
	rlog.CheckShowError(err)
}

type Customer struct {
	CustomerBase    `xorm:"extends"`
	wx.BaseCustomer `xorm:"extends"`
	wx.BaseOther    `xorm:"extends"`
}

type CustomerBase struct {
	ID       `xorm:"extends"`
	BaseTime `xorm:"extends"`
}

func (CustomerBase) TableName() string {
	return "tb_customer"
}

func NewCustomer() *Customer {
	return &Customer{
		CustomerBase: CustomerBase{
			BaseTime: NewBaseTime(),
		},
	}
}
