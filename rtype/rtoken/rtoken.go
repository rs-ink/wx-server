package rtoken

import (
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"wx-server/rlog"
)

const tokenPwd = "dudu-8888"

func NewToken(info interface{}) string {
	data, _ := json.Marshal(info)
	var mapResult map[string]interface{}
	_ = json.Unmarshal(data, &mapResult)
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims(mapResult))
	tokenString, err := token.SignedString([]byte(tokenPwd))
	if err != nil {
		panic(err)
	}
	return tokenString
}

func ParseToken(token string) (map[string]interface{}, bool) {
	t, err := jwt.Parse(token, func(token *jwt.Token) (i interface{}, e error) {
		return []byte(tokenPwd), nil
	})
	if err == nil && t.Valid {
		return t.Claims.(jwt.MapClaims), true
	}
	return nil, false
}

func BindInTokenByByte(token []byte, value interface{}) bool {
	t, err := jwt.Parse(string(token), func(token *jwt.Token) (i interface{}, e error) {
		return []byte(tokenPwd), nil
	})
	if err == nil && t.Valid {
		data, err := json.Marshal(t.Claims.(jwt.MapClaims))
		if err == nil {
			err = json.Unmarshal(data, value)
			if err == nil {
				return true
			}
		}
	}
	rlog.Warn(err)
	return false
}

func BindInToken(token string, value interface{}) bool {
	t, err := jwt.Parse(token, func(token *jwt.Token) (i interface{}, e error) {
		return []byte(tokenPwd), nil
	})
	if err == nil && t.Valid {
		data, err := json.Marshal(t.Claims.(jwt.MapClaims))
		if err == nil {
			err = json.Unmarshal(data, value)
			if err == nil {
				return true
			}
		}
	}
	rlog.Warn(err)
	return false
}
