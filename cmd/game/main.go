package main

import (
	"github.com/xyy0411/ebiten_paractice/internal/app"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	g := app.New()

	ebiten.SetWindowSize(800, 600)
	ebiten.SetWindowTitle("Ebiten 角色移动 + Idle/Run Demo")

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
