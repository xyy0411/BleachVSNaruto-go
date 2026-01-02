package draw

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"golang.org/x/image/colornames"
)

func Ground(screen *ebiten.Image, groundY float64) {
	vector.StrokeLine(screen, 0, float32(groundY-2), 800, float32(groundY), 2, colornames.Aliceblue, true)
}
