package zangetsu

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/xyy0411/bleachVSnaruto/assets"
	"github.com/xyy0411/bleachVSnaruto/game_map"
)

var (
	terrainuri    = "./assets/maps/zangetsu_haka/Symbol 8.png"
	decorationuri = "./assets/maps/zangetsu_haka/10.png"
	backgrounduri = "./assets/maps/zangetsu_haka/2.png"
)

type zangetsu struct {
	game_map.BaseInfo
}

func init() {
	game_map.StdRegistry.RegisterMap("zangetsu", new(zangetsu))
}

func (z *zangetsu) Init() {
	assets.StdImagePool.LoadImage(terrainuri)
	assets.StdImagePool.LoadImage(decorationuri)
	assets.StdImagePool.LoadImage(backgrounduri)
}
func (z *zangetsu) Draw(screen *ebiten.Image, groundY float64) {
	terrain := assets.StdImagePool.GetImage(terrainuri)
	decorationImg := assets.StdImagePool.GetImage(decorationuri)
	background := assets.StdImagePool.GetImage(backgrounduri)
	bgW := background.Bounds().Dx()
	bgH := background.Bounds().Dy()

	tilesX := (screen.Bounds().Dx() / bgW) + 1
	tilesY := (screen.Bounds().Dy() / bgH) + 1

	for y := range tilesY {
		for x := range tilesX {
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

func (z *zangetsu) GetBaseInfo() game_map.BaseInfo {
	birdViewKey := "zangetsu_view"
	terrain := assets.StdImagePool.GetImage(terrainuri)
	decorationImg := assets.StdImagePool.GetImage(decorationuri)
	background := assets.StdImagePool.GetImage(backgrounduri)
	bgW := background.Bounds().Dx()
	bgH := background.Bounds().Dy()

	screen := ebiten.NewImage(bgW, bgH)
	screen.DrawImage(background, &ebiten.DrawImageOptions{})

	terrainOP := &ebiten.DrawImageOptions{}
	terrainOP.GeoM.Translate(0, float64(600-terrain.Bounds().Dy()))
	screen.DrawImage(terrain, terrainOP)
	decorationOP := &ebiten.DrawImageOptions{}
	decorationOP.GeoM.Translate(0, float64(600-decorationImg.Bounds().Dy()))
	screen.DrawImage(decorationImg, decorationOP)
	assets.StdImagePool.PostImage(birdViewKey, screen)
	z.BirdViewKey = birdViewKey
	return z.BaseInfo
}
