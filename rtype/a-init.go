package rtype

import (
	"time"
	"wx-server/database"
	"xorm.io/xorm"
)

var db *xorm.Engine

func init() {
	db = database.Engine()
}

type ID struct {
	Id int `json:"id" xorm:"not null pk autoincr unique INT(11)" info:"id"`
}

type BaseTime struct {
	CreateTime time.Time `json:"createTime" xorm:"DATETIME  created " info:"创建时间"`
	UpdateTime time.Time `json:"updateTime" xorm:"DATETIME  updated " info:"更新时间"`
}

func NewBaseTime() BaseTime {
	return BaseTime{
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	}
}
