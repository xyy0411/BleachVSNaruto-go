package world

import (
	"github.com/xyy0411/bleachVSnaruto/assets"
	"github.com/xyy0411/bleachVSnaruto/game_map"
)

type World struct {
	GroundY       float64
	GroundPainter game_map.MapInter
	MapInfo       *MapInfo
}

// ResolveGround 用于确定给定的y坐标是否接触或低于地面
// 如果不在地面则修改并返回新的y坐标(一般在地下会进行更改)
func (w *World) ResolveGround(y float64) (newY float64, onGround bool) {
	if y >= w.GroundY {
		return w.GroundY, true
	}
	return y, false
}
func (w *World) UpdateMapInfo() *World {
	if w.GroundPainter == nil {
		return w
	}
	info := w.GroundPainter.GetBaseInfo()
	vp := assets.StdImagePool.GetImage(info.BirdViewKey)
	w.MapInfo = &MapInfo{
		ID:     info.ID,
		PicURI: info.BirdViewKey,
		Bound:  Bound{Left: 0, Right: float64(vp.Bounds().Dx()), Bottom: 0},
	}
	return w
}

func (w *World) BGM(uri string) *World {
	w.MapInfo.BGM = uri
	return w
}
