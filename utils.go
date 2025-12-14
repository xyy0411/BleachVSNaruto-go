package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image/color"
	"log"
)

type Platform struct {
	X, Y float64 // 左上角
	W, H float64
}

var platforms = []Platform{
	{X: 300, Y: 380, W: 120, H: 20},
}

func drawPlatforms(screen *ebiten.Image) {
	for _, p := range platforms {
		img := ebiten.NewImage(int(p.W), int(p.H))
		img.Fill(color.RGBA{R: 100, G: 100, B: 100, A: 255})

		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(p.X, p.Y)
		screen.DrawImage(img, op)
	}
}

func drawSprite(
	screen *ebiten.Image,
	img *ebiten.Image,
	x, y float64,
	facing int,
) {
	op := &ebiten.DrawImageOptions{}

	if facing == -1 {
		w := img.Bounds().Dx()
		op.GeoM.Scale(-1, 1)
		op.GeoM.Translate(float64(w), 0)
	}

	op.GeoM.Translate(x, y-float64(img.Bounds().Dy()))
	screen.DrawImage(img, op)
}

func loadImage(path string) *ebiten.Image {
	img, _, err := ebitenutil.NewImageFromFile(path)
	if err != nil {
		log.Fatal("加载图片失败:", path, err)
	}
	return img
}

func loadAssets() {
	assets.IdleFrames = []*ebiten.Image{
		loadImage("character/shapes/8.png"),
		loadImage("character/shapes/10.png"),
		loadImage("character/shapes/12.png"),
		loadImage("character/shapes/14.png"),
	}
	assets.RunFrames = []*ebiten.Image{
		loadImage("character/shapes/17.png"),
		loadImage("character/shapes/19.png"),
		loadImage("character/shapes/21.png"),
		loadImage("character/shapes/23.png"),
		loadImage("character/shapes/25.png"),
		loadImage("character/shapes/27.png"),
		loadImage("character/shapes/29.png"),
		// loadImage("character/shapes/31.png"),
	}
	assets.JumpFrames = []*ebiten.Image{
		loadImage("character/shapes/39.png"),
		loadImage("character/shapes/41.png"),
		loadImage("character/shapes/43.png"),
	}
	assets.LandFrames = []*ebiten.Image{
		loadImage("character/shapes/45.png"),
		loadImage("character/shapes/47.png"),
		loadImage("character/shapes/49.png"),
	}

	assets.AttackFrames = []*ebiten.Image{}
}
