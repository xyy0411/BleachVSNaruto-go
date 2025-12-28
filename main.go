package main

import (
	"github.com/xyy0411/ebiten_paractice/internal/app"
	"github.com/xyy0411/ebiten_paractice/internal/global"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	g := app.New()

	ebiten.SetWindowSize(global.ScreenWidth*global.WindowScale, global.ScreenHeight*global.WindowScale)
	ebiten.SetWindowTitle("死神VS火影 Demo")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}

}
