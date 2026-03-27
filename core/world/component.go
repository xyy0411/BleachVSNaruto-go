package world

// MapInfo 地图基础信息
type MapInfo struct {
	ID      string
	PicURI  string
	BGM     string
	Bound   Bound
}

type Bound struct {
	Left   float64 // 左边界
	Right  float64
	Bottom float64
}
