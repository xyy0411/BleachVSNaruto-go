package world

import "github.com/xyy0411/bleachVSnaruto/game_map"

type World struct {
	GroundY       float64
	GroundPainter game_map.MapInter
}

// ResolveGround 用于确定给定的y坐标是否接触或低于地面
// 如果不在地面则修改并返回新的y坐标(一般在地下会进行更改)
func (w *World) ResolveGround(y float64) (newY float64, onGround bool) {
	if y >= w.GroundY {
		return w.GroundY, true
	}
	return y, false
}
