package world

// MapInfo 描述地图的基础信息与世界边界
type MapInfo struct {
	ID     string
	PicURI string
	BGM    string
	Bound  Bound
}

// Bound 描述地图在世界坐标中的边界
type Bound struct {
	Left   float64 // 左边界
	Right  float64 // 右边界
	Bottom float64 // 底边界
}
