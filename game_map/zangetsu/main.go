package zangetsu

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/xyy0411/bleachVSnaruto/assets"
	"github.com/xyy0411/bleachVSnaruto/game_map"
)

const (
	logicalHeight = 600
)

var (
	// Name 是斩月地图的注册名
	Name          = "zangetsu"
	terrainuri    = "./assets/maps/zangetsu_haka/Symbol 8.png"
	decorationuri = "./assets/maps/zangetsu_haka/10.png"
	backgrounduri = "./assets/maps/zangetsu_haka/2.png"
)

type zangetsu struct {
	game_map.BaseInfo
}

func init() {
	game_map.StdRegistry.RegisterMap(Name, new(zangetsu))
}

func (z *zangetsu) Init() {
	assets.StdImagePool.LoadImageArray(terrainuri, decorationuri, backgrounduri)
}

func (z *zangetsu) Draw(screen *ebiten.Image, cameraX float64, zoom float64, _ float64) {
	worldImage := assets.StdImagePool.GetImage(z.GetBaseInfo().BirdViewKey, true)
	offsetY := (float64(screen.Bounds().Dy()) - float64(screen.Bounds().Dy())*zoom) / 2
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(zoom, zoom)
	op.GeoM.Translate(-cameraX*zoom, offsetY)
	screen.DrawImage(worldImage, op)
}

func (z *zangetsu) GetBaseInfo() game_map.BaseInfo {
	if z.BirdViewKey != "" {
		return z.BaseInfo
	}

	birdViewKey := "zangetsu_view"
	terrain := assets.StdImagePool.GetImage(terrainuri)
	decorationImg := assets.StdImagePool.GetImage(decorationuri)
	background := assets.StdImagePool.GetImage(backgrounduri)

	worldWidth := terrain.Bounds().Dx()
	bgWidth := background.Bounds().Dx() - 1

	screen := ebiten.NewImage(worldWidth, logicalHeight)
	for x := 0; x < worldWidth; x += bgWidth {
		backgroundOP := &ebiten.DrawImageOptions{}
		backgroundOP.GeoM.Translate(float64(x), 0)
		screen.DrawImage(background, backgroundOP)
	}

	terrainOP := &ebiten.DrawImageOptions{}
	terrainOP.GeoM.Translate(0, float64(logicalHeight-terrain.Bounds().Dy()))
	screen.DrawImage(terrain, terrainOP)

	decorationOP := &ebiten.DrawImageOptions{}
	decorationOP.GeoM.Translate(0, float64(logicalHeight-decorationImg.Bounds().Dy()))
	screen.DrawImage(decorationImg, decorationOP)

	assets.StdImagePool.PostImage(birdViewKey, screen, true)
	z.BirdViewKey = birdViewKey
	z.ID = Name
	return z.BaseInfo
}
