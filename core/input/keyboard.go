package input

import (
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
	return models.InputFrame{
		Left:       ebiten.IsKeyPressed(ebiten.KeyArrowLeft),
		Right:      ebiten.IsKeyPressed(ebiten.KeyArrowRight),
		Up:         ebiten.IsKeyPressed(ebiten.KeyArrowUp),
		Down:       ebiten.IsKeyPressed(ebiten.KeyArrowDown),
		Attack:     ebiten.IsKeyPressed(ebiten.KeyDigit1),
		Dash:       ebiten.IsKeyPressed(ebiten.KeyDigit2),
		Jump:       ebiten.IsKeyPressed(ebiten.KeyDigit3),
		LangAtk:    ebiten.IsKeyPressed(ebiten.KeyDigit4),
		Outbreak:   ebiten.IsKeyPressed(ebiten.KeyDigit5),
		Assistance: ebiten.IsKeyPressed(ebiten.KeyDigit6),
	}
}
