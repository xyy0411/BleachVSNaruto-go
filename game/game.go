package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/xyy0411/ebiten_paractice/engine"
	"image/color"
)

type Game struct {
	Engine *engine.Engine
}

func (g *Game) Update() error {
	g.Engine.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.Black)
	g.Engine.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 800, 600
}
