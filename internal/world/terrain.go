package world

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/xyy0411/ebiten_paractice/internal/global"
	"github.com/xyy0411/ebiten_paractice/internal/render"
	"image/color"
)

type Terrain struct {
	// 地面宽度
	Width float64
	// 地图高度，角色的脚底会贴着这条线运动
	Height float64
	// 地面上可能存在的台阶
	Platforms []Platform
}

// Platform 用于描述简单的矩形平台
type Platform struct {
	X float64 // 左上角 X 坐标
	Y float64 // 左上角 Y 坐标
	W float64 // 平台宽度
	H float64 // 平台高度
}

func (t *Terrain) Draw(screen *ebiten.Image) {
	decoration := render.Assets.MapsAsset.Decoration
	ground := render.Assets.MapsAsset.Ground

	decOP := &ebiten.DrawImageOptions{}
	decOP.GeoM.Translate(0, float64(global.ScreenHeight-decoration.Bounds().Dy()))
	screen.DrawImage(decoration, decOP)

	groundOP := &ebiten.DrawImageOptions{}
	groundOP.GeoM.Translate(0, float64(global.ScreenHeight-ground.Bounds().Dy()))
	screen.DrawImage(ground, groundOP)
}

// GroundY 实现接口
func (t *Terrain) GroundY(x float64) float64 {
	return t.Height + 350
}

func NewTerrain(w, h float64) *Terrain {
	return &Terrain{
		Width:     w,
		Height:    h,
		Platforms: Platforms,
	}
}

func (t *Terrain) CheckPlatform(x, y, vy float64) (bool, float64) {
	// 判断是否落在平台上
	if y >= t.GroundY(x) {
		y = t.GroundY(x)
		vy = 0
		return true, y
	}

	for _, p := range Platforms {
		if x >= p.X && x <= p.X+p.W {
			if y <= p.Y && y >= p.Y {
				y = p.Y
				vy = 0
				return true, y
			}
		}
	}
	return false, y
}

// Platforms 维护所有静态平台信息
var Platforms = []Platform{{X: 300, Y: 380, W: 120, H: 20}}

// DrawPlatforms 渲染所有平台
func DrawPlatforms(screen *ebiten.Image) {
	for _, p := range Platforms {
		img := ebiten.NewImage(int(p.W), int(p.H))
		img.Fill(color.RGBA{R: 100, G: 100, B: 100, A: 255})

		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(p.X, p.Y)
		screen.DrawImage(img, op)
	}
}
