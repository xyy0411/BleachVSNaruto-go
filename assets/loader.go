package assets

import (
	"bytes"
	"os"

	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"
)

func LoadWavBytes(path string, sampleRate int) ([]byte, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	d, err := wav.DecodeWithSampleRate(sampleRate, bytes.NewReader(b))
	if err != nil {
		return nil, err
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(d)
	return buf.Bytes(), err
}

func NewPlayer(ctx *audio.Context, data []byte, volume float64) *audio.Player {
	p := ctx.NewPlayerFromBytes(data)
	p.SetVolume(volume)
	return p
}
