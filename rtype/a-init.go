package rtype

import (
	"encoding/json"
	"time"
	"wx-server/database"
	"xorm.io/xorm"
)

var db *xorm.Engine

func init() {
	db = database.Engine()
}

type ID struct {
	Id  int         `json:"id" xorm:"not null pk autoincr unique INT(11)" info:"id"`
	Ext interface{} `xorm:"-" json:"ext,omitempty" info:"扩展字段"`
}

type BaseTime struct {
	CreateTime time.Time `json:"createTime" xorm:"DATETIME  default current_timestamp " info:"创建时间"`
	UpdateTime time.Time `json:"updateTime" xorm:"DATETIME  updated " info:"更新时间"`
}

func NewBaseTime() BaseTime {
	return BaseTime{
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	}
}

const rsDateFormat string = "2006-01-02"

const rsDateTimeFormat string = "2006-01-02 15:04:05"

type RDate time.Time

func (l RDate) MarshalJSON() ([]byte, error) {
	format := time.Time(l).Format(rsDateFormat)
	return json.Marshal(format)
}

func (l *RDate) UnmarshalJSON(b []byte) error {
	var aa string
	err := json.Unmarshal(b, &aa)
	if err == nil {
		dd, err := time.Parse(rsDateFormat, aa)
		if err == nil {
			*l = RDate(dd)
		}
	}
	return err
}

type RDateTime time.Time

func (t *RDateTime) UnmarshalJSON(b []byte) error {
	var aa string
	err := json.Unmarshal(b, &aa)
	if err == nil {
		dd, err := time.Parse(rsDateTimeFormat, aa)
		if err == nil {
			*t = RDateTime(dd)
		}
	}
	return err
}

func (t RDateTime) MarshalJSON() ([]byte, error) {
	format := time.Time(t).Format(rsDateTimeFormat)
	return json.Marshal(format)
}
