package rtype

import (
	"testing"
	"time"
	"wx-server/rlog"
)

func TestShowTime(t *testing.T) {
	rlog.Warn(time.Now().Unix())
}
