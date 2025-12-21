package app

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// GroundY 地面高度，角色的脚底会贴着这条线运动
const GroundY = 500

// Platform 用于描述简单的矩形平台
type Platform struct {
	X float64 // 左上角 X 坐标
	Y float64 // 左上角 Y 坐标
	W float64 // 平台宽度
	H float64 // 平台高度
}

// Platforms 维护所有静态平台信息
var Platforms = []Platform{{X: 300, Y: 380, W: 120, H: 20}}

// DrawGroundLine 绘制地面基准线
func DrawGroundLine(screen *ebiten.Image) {
	ebitenutil.DrawLine(screen, 0, GroundY, 800, GroundY, color.RGBA{R: 255, A: 255})
}

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
