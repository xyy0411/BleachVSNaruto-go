package app

import (
	"github.com/xyy0411/ebiten_paractice/internal/character"
	"github.com/xyy0411/ebiten_paractice/internal/world"
)

type World struct {
	Players []*character.Character
	Map     *world.Terrain
}
