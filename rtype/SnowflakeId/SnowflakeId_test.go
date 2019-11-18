package SnowflakeId

import (
	"github.com/sony/sonyflake"
	"log"
	"testing"
)

func TestGetSnowflakeId(t *testing.T) {
	id, _ := sonyflake.NewSonyflake(sonyflake.Settings{}).NextID()
	log.Println(id)
	log.Println(int64(id))
	log.Println(NextID())
}
func BenchmarkNextID(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NextID()
	}
}
