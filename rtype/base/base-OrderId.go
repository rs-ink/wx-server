package base

import (
	"encoding/json"
	"strconv"
	"wx-server/rtype/snowflakeId"
)

type IDuInt64 uint64

func (id IDuInt64) String() string {
	return strconv.FormatUint(uint64(id), 10)
}

func (id *IDuInt64) UnmarshalJSON(data []byte) error {
	var idStr string
	err := json.Unmarshal(data, &idStr)
	if err != nil {
		return err
	}
	n, err := strconv.ParseUint(idStr, 10, 64)
	if err == nil {
		*id = IDuInt64(n)
		return nil
	} else {
		return err
	}
}

func (id IDuInt64) MarshalJSON() ([]byte, error) {
	str := strconv.FormatUint(uint64(id), 10)
	return json.Marshal(str)
}

type OrderID struct {
	Id IDuInt64 `json:"id" xorm:"not null pk unique"` //手动指定unsigned
}

func NewOrderId() OrderID {
	return OrderID{Id: IDuInt64(snowflakeId.NextID())}
}
