package audio

import (
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/xyy0411/bleachVSnaruto/assets"
)

type System struct {
	Ctx   *audio.Context
	Cache map[string][]byte
}

var Default *System

func Init(ctx *audio.Context) {
	Default = &System{
		Ctx:   ctx,
		Cache: map[string][]byte{},
	}
}

func (s *System) getBytes(path string) []byte {
	if b, ok := s.Cache[path]; ok {
		return b
	}
	// 如果文件不存在的话那么就直接播放空的音效
	b, _ := assets.LoadWavBytes(path, s.Ctx.SampleRate())
	s.Cache[path] = b
	return b
}

func (s *System) Play(path string, volume float64) {
	b := s.getBytes(path)
	p := assets.NewPlayer(s.Ctx, b, volume)
	p.Play()
}
