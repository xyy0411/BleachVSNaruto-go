package input

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/xyy0411/bleachVSnaruto/models"
)

type KeyboardSource struct{}

func (k *KeyboardSource) Read() models.InputFrame {
	return models.InputFrame{
		Left:   ebiten.IsKeyPressed(ebiten.KeyA),
		Right:  ebiten.IsKeyPressed(ebiten.KeyD),
		Up:     ebiten.IsKeyPressed(ebiten.KeyW),
		Down:   ebiten.IsKeyPressed(ebiten.KeyS),
		Attack: ebiten.IsKeyPressed(ebiten.KeyJ),
		Dash:   ebiten.IsKeyPressed(ebiten.KeyL),
		Jump:   ebiten.IsKeyPressed(ebiten.KeyK),
	}
}
