package wx

import "time"

type AccessToken struct {
	Error
	ExpiresIn   int64     `json:"expires_in"`
	AccessToken string    `json:"access_token"`
	CreateTime  time.Time `json:"createTime"`
}
type Ticket struct {
	Error
	ExpiresIn  int64     `json:"expires_in"`
	Ticket     string    `json:"ticket"`
	CreateTime time.Time `json:"createTime"`
}

type EncryptedData struct {
	EncryptedData string `json:"encryptedData"`
	Iv            string `json:"iv"`
}

type JsSignature struct {
	AppId     string `json:"appId"`
	NonceStr  string `json:"nonceStr"`
	TimeStamp string `json:"timeStamp"`
	Signature string `json:"signature"`
	SignType  string `json:"signType"`
}
