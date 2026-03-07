package zangetsu

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/xyy0411/bleachVSnaruto/assets"
)

type zangetsu struct {
	terrain    *ebiten.Image
	decoration *ebiten.Image
	background *ebiten.Image
}

func Init(screen *ebiten.Image, groundY float64) {
	terrain := assets.StdImagePool.GetImage("./assets/maps/zangetsu_haka/Symbol 8.png")
	decorationImg := assets.StdImagePool.GetImage("./assets/maps/zangetsu_haka/10.png")
	background := assets.StdImagePool.GetImage("./assets/maps/zangetsu_haka/2.png")
	bgW := background.Bounds().Dx()
	bgH := background.Bounds().Dy()

	screenW, screenH := screen.Size()
	tilesX := (screenW / bgW) + 1
	tilesY := (screenH / bgH) + 1

	for y := 0; y < tilesY; y++ {
		for x := 0; x < tilesX; x++ {
			opts := &ebiten.DrawImageOptions{}
			opts.GeoM.Translate(float64(x*bgW), float64(y*bgH))
			screen.DrawImage(background, opts)
		}
	}
	terrainOP := &ebiten.DrawImageOptions{}
	terrainOP.GeoM.Translate(0, float64(600-terrain.Bounds().Dy()))
	screen.DrawImage(terrain, terrainOP)
	decorationOP := &ebiten.DrawImageOptions{}
	decorationOP.GeoM.Translate(0, float64(600-decorationImg.Bounds().Dy()))
	screen.DrawImage(decorationImg, decorationOP)
}
