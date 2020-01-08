package base

//经纬度
type LonAndLat struct {
	Longitude float64 `json:"longitude" info:"经度"`
	Latitude  float64 `json:"latitude"  info:"纬度"`
}

//运动经纬度
type LonAndLatMotion struct {
	LonAndLat
	Speed    int `json:"speed" info:"速度"`
	Accuracy int `json:"accuracy" info:"精确度"`
}
