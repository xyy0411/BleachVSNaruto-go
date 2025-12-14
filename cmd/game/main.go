package main

import (
	"log"

	"ebiten_paractice/internal/game"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	g := game.New()

	ebiten.SetWindowSize(800, 600)
	ebiten.SetWindowTitle("Ebiten 角色移动 + Idle/Run Demo")

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
