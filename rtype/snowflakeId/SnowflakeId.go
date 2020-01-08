package snowflakeId

import (
	"github.com/sony/sonyflake"
	"time"
)

var sf *sonyflake.Sonyflake

//TODO 核心配置
const baseTime = "2019-01-02 15:04:05"

func init() {
	start, _ := time.Parse("2006-01-02 15:04:05", baseTime)
	st := sonyflake.Settings{
		MachineID: func() (u uint16, e error) {
			return uint16(1), nil
		},
		StartTime: start,
	}
	sf = sonyflake.NewSonyflake(st)
}

func NextID() uint64 {
	id, err := sf.NextID()
	if err == nil {
		return id
	} else {
		return uint64(0)
	}
}
