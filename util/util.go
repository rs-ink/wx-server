package util

import (
	"bytes"
	"crypto/rand"
	"encoding/json"
	"math/big"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func IsNumber(str string) bool {
	if strings.HasPrefix(str, "0") {
		return false
	}

	ok, _ := regexp.MatchString("^[0-9]*$", str)
	return ok
}

func ToJsonData(v interface{}) string {
	data, _ := json.Marshal(v)
	return string(data)
}

func ToDefaultInt(str string, re int) int {
	if str != "" {
		result, err := strconv.Atoi(str)
		if err == nil {
			return result
		}
	}
	return re
}

func ExitsOrPrefixExits(slice []string, key string) bool {
	for _, v := range slice {
		if v == key || strings.HasPrefix(key, v) {
			return true
		}
	}
	return false
}

var loc *time.Location

const DefaultTimeFormate = "2006-01-02 15:04:05"
const DefaultDateFormate = "2006-01-02"

func init() {
	loc, _ = time.LoadLocation("Local")
}

func GetCurrentDateStr() string {
	return time.Now().Format(DefaultDateFormate)
}

func ParseTime(timeStr string) (t time.Time, err error) {
	return time.Parse(DefaultTimeFormate, timeStr)
}

func CreateRandomString(len int) string {
	var container string
	var str = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	b := bytes.NewBufferString(str)
	length := b.Len()
	bigInt := big.NewInt(int64(length))
	for i := 0; i < len; i++ {
		randomInt, _ := rand.Int(rand.Reader, bigInt)
		container += string(str[randomInt.Int64()])
	}
	return container
}
