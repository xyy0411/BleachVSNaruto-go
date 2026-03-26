package input

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/xyy0411/bleachVSnaruto/models"
)

type KeyboardSource struct{}

func (k *KeyboardSource) Read() models.InputFrame {
	return models.InputFrame{
		Left:       ebiten.IsKeyPressed(ebiten.KeyA),
		Right:      ebiten.IsKeyPressed(ebiten.KeyD),
		Up:         ebiten.IsKeyPressed(ebiten.KeyW),
		Down:       ebiten.IsKeyPressed(ebiten.KeyS),
		Attack:     ebiten.IsKeyPressed(ebiten.KeyJ),
		Dash:       ebiten.IsKeyPressed(ebiten.KeyL),
		Jump:       ebiten.IsKeyPressed(ebiten.KeyK),
		LangAtk:    ebiten.IsKeyPressed(ebiten.KeyU),
		Outbreak:   ebiten.IsKeyPressed(ebiten.KeyI),
		Assistance: ebiten.IsKeyPressed(ebiten.KeyO),
	}
}

type KeyboardSourceWithTwo struct{}

func (k *KeyboardSourceWithTwo) Read() models.InputFrame {
	if ebiten.IsKeyPressed(ebiten.KeyDigit3) {
		fmt.Println("3")
	}
	return models.InputFrame{
		Left:       ebiten.IsKeyPressed(ebiten.KeyArrowLeft),
		Right:      ebiten.IsKeyPressed(ebiten.KeyArrowRight),
		Up:         ebiten.IsKeyPressed(ebiten.KeyArrowUp),
		Down:       ebiten.IsKeyPressed(ebiten.KeyArrowDown),
		Attack:     ebiten.IsKeyPressed(ebiten.KeyNumpad1),
		Dash:       ebiten.IsKeyPressed(ebiten.KeyNumpad2),
		Jump:       ebiten.IsKeyPressed(ebiten.KeyNumpad3),
		LangAtk:    ebiten.IsKeyPressed(ebiten.KeyNumpad4),
		Outbreak:   ebiten.IsKeyPressed(ebiten.KeyNumpad5),
		Assistance: ebiten.IsKeyPressed(ebiten.KeyNumpad6),
	}
}
