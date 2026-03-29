package world

import (
	"github.com/xyy0411/bleachVSnaruto/assets"
	"github.com/xyy0411/bleachVSnaruto/game_map"
)

type World struct {
	GroundY       float64
	GroundPainter game_map.MapInter
	MapInfo       *MapInfo
	Camera        *Camera
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

	w.GroundPainter.Init()
	info := w.GroundPainter.GetBaseInfo()
	vp := assets.StdImagePool.GetImage(info.BirdViewKey)
	w.MapInfo = &MapInfo{
		ID:     info.ID,
		PicURI: info.BirdViewKey,
		Bound:  Bound{Left: 0, Right: float64(vp.Bounds().Dx()), Bottom: 0},
	}
	if w.Camera != nil {
		w.Camera.ApplyZoomLimit(w.MapInfo.Bound)
		w.Camera.ClampX(w.MapInfo.Bound)
	}
	return w
}

func (w *World) FollowTargetsX(targets ...float64) {
	if w == nil || w.Camera == nil || w.MapInfo == nil || len(targets) == 0 {
		return
	}
	w.Camera.FollowTargets(w.MapInfo.Bound, targets...)
}

func (w *World) ClampBodyX(x float64) float64 {
	return w.ClampBodyRectX(x, 0)
}

func (w *World) ClampBodyRectX(x float64, width float64) float64 {
	if w == nil || w.MapInfo == nil {
		return x
	}
	if width < 0 {
		width = 0
	}

	// 右边界要减去角色当前帧宽度
	// 否则 body.X 虽然没越界 贴图右侧还是会露到地图外
	maxX := w.MapInfo.Bound.Right - width
	if maxX < w.MapInfo.Bound.Left {
		maxX = w.MapInfo.Bound.Left
	}

	if x < w.MapInfo.Bound.Left {
		return w.MapInfo.Bound.Left
	}
	if x > maxX {
		return maxX
	}
	return x
}

func (w *World) BGM(uri string) *World {
	if w.MapInfo != nil {
		w.MapInfo.BGM = uri
	}
	return w
}
