package world

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/xyy0411/bleachVSnaruto/assets"
	"github.com/xyy0411/bleachVSnaruto/core/charactor"
	"github.com/xyy0411/bleachVSnaruto/game_map"
)

// World 保存地图、地面和摄像机等战斗场景环境信息。
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

// UpdateMapInfo 初始化地图信息，并在可用时同步修正摄像机状态。
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

// FollowTargetsX 使用多个目标点更新摄像机的水平跟随。
func (w *World) FollowTargetsX(targets ...float64) {
	if w == nil || w.Camera == nil || w.MapInfo == nil || len(targets) == 0 {
		return
	}
	w.Camera.FollowTargets(w.MapInfo.Bound, targets...)
}

// ClampBodyX 将角色横坐标限制在地图边界内。
func (w *World) ClampBodyX(x float64) float64 {
	return w.ClampBodyRectX(x, 0)
}

// ClampBodyRectX 将带宽度的物体横坐标限制在地图边界内。
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

// BGM 设置当前地图的背景音乐资源标识。
func (w *World) BGM(uri string) *World {
	if w.MapInfo != nil {
		w.MapInfo.BGM = uri
	}
	return w
}

// CameraParams 返回摄像机的当前参数（X位置和缩放倍率）
func (w *World) CameraParams() (cameraX float64, cameraZoom float64) {
	if w.Camera == nil {
		return 0, 1
	}
	return w.Camera.X, w.Camera.Scale()
}

// Draw 绘制世界场景（地面和角色）
func (w *World) Draw(screen *ebiten.Image, actors []charactor.Character) {
	cameraX, cameraZoom := w.CameraParams()

	if w.GroundPainter != nil {
		w.GroundPainter.Draw(screen, cameraX, cameraZoom, w.GroundY)
	}

	for _, actor := range actors {
		rt := actor.GetRuntime()
		if rt == nil {
			continue
		}

		frame := rt.AnimPlayer.CurrentFrame()
		if frame == nil {
			continue
		}

		op := &ebiten.DrawImageOptions{}
		if rt.Facing == -1 {
			op.GeoM.Scale(-1, 1)
			op.GeoM.Translate(float64(frame.Bounds().Dx()), 0)
		}

		op.GeoM.Scale(cameraZoom, cameraZoom)
		drawX, drawY := w.Camera.WorldToScreen(frameDrawPosition(rt))
		op.GeoM.Translate(drawX, drawY)
		screen.DrawImage(frame, op)
	}
}

func frameDrawPosition(rt *charactor.Runtime) (float64, float64) {
	frame := rt.AnimPlayer.CurrentFrame()
	if frame == nil {
		return rt.Body.X, rt.Body.Y
	}

	meta := rt.AnimPlayer.CurrentFrameMeta()
	originX := 0.0
	originY := float64(frame.Bounds().Dy())
	if meta != nil {
		originX = meta.Origin.X
		originY = meta.Origin.Y
	}

	drawX := rt.Body.X - originX
	if rt.Facing == -1 {
		drawX = rt.Body.X - (float64(frame.Bounds().Dx()) - originX)
	}

	return drawX, rt.Body.Y - originY
}
