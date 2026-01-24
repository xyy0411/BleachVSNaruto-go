package assets

import (
	"bytes"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/xyy0411/bleachVSnaruto/global"
)

func LoadImage(path string) *ebiten.Image {
	img, _, err := ebitenutil.NewImageFromFile(path)
	if err != nil {
		global.Logger.Fatal("load image failed:", path, err)
	}
	return img
}

func LoadWavBytes(path string) []byte {
	b, err := os.ReadFile(path)
	if err != nil {
		global.Logger.Fatal(err)
	}
	d, err := wav.DecodeWithoutResampling(bytes.NewReader(b))
	if err != nil {
		global.Logger.Fatal(err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(d)
	if err != nil {
		global.Logger.Fatal(err)
	}
	return buf.Bytes()
}

func NewPlayer(ctx *audio.Context, data []byte, volume float64) *audio.Player {
	p := ctx.NewPlayerFromBytes(data)
	p.SetVolume(volume)
	return p
}
