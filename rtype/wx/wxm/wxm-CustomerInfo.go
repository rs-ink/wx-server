package wxm

type UserInfo struct {
	OpenId     string `json:"openid"`
	UnionId    string `json:"unionid"`
	NickName   string `json:"nickname"`
	Sex        int    `json:"gender"`
	Language   string `json:"language"`
	City       string `json:"city"`
	Province   string `json:"province"`
	Country    string `json:"country"`
	HeadImgUrl string `json:"avatarUrl"`
}
