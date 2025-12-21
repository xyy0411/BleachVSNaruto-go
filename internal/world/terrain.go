package world

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
)

type Terrain struct {
	Ground    float64
	Platforms []Platform
}

// GroundY 地面高度，角色的脚底会贴着这条线运动
const GroundY = 500

// Platform 用于描述简单的矩形平台
type Platform struct {
	X float64 // 左上角 X 坐标
	Y float64 // 左上角 Y 坐标
	W float64 // 平台宽度
	H float64 // 平台高度
}

// GroundY 实现接口
func (t *Terrain) GroundY(x float64) float64 {
	return GroundY
}

func NewTerrain() *Terrain {
	return &Terrain{
		Ground:    GroundY,
		Platforms: Platforms,
	}
}

func (t *Terrain) CheckPlatform(x, y, vy float64) (bool, float64) {
	// 判断是否落在平台上
	if y >= GroundY {
		y = GroundY
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
