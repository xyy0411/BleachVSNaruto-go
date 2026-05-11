package game

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/xyy0411/bleachVSnaruto/debugview"
	"github.com/xyy0411/bleachVSnaruto/engine"
)

type Game struct {
	Engine *engine.Engine
	Debug  *debugview.Panel
}

func (g *Game) Update() error {
	g.Engine.Update()
	if g.Debug != nil {
		g.Debug.Update()
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.Black)
	g.Engine.Draw(screen)
	if g.Debug != nil {
		g.Debug.Draw(screen, g.Engine)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 800, 600 // 返回固定的窗口尺寸 800x600
}
