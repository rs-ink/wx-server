package domain

import "wx-server/rtype/base"

type ShopAddress struct {
	ShopAddressId int    `json:"shopAddressId"`
	ShopId        int    `json:"shopId" info:"店铺编码"`
	Name          string `json:"name" info:"店铺名称"`
	Address       string `json:"address" info:"地址"`
	base.LonAndLat
	WxLongitude float64 `json:"wxLongitude" info:"微信经度"`
	WxLatitude  float64 `json:"wxLatitude" info:"微信纬度"`
}
