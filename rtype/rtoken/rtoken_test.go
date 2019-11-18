package rtoken

import (
	"testing"
)

func TestNewToken(t *testing.T) {
	info := make(map[string]interface{}, 0)
	log.Info(NewToken(info))
}

func TestParseToken(t *testing.T) {
	log.Info(ParseToken(""))
}
