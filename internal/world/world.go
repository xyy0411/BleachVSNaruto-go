package world

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image/color"
)

// DrawGroundLine 绘制地面基准线
func DrawGroundLine(screen *ebiten.Image) {
	ebitenutil.DrawLine(screen, 0, GroundY, 800, GroundY, color.RGBA{R: 255, A: 255})
}
