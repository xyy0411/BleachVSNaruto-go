package audio

import (
	"time"

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
	if len(b) == 0 {
		return
	}
	p := assets.NewPlayer(s.Ctx, b, volume)
	p.Play()
	go closePlayerWhenFinished(p)
}

// closePlayerWhenFinished 等待一次性音效播放结束后主动释放底层资源
func closePlayerWhenFinished(p *audio.Player) {
	ticker := time.NewTicker(100 * time.Millisecond)
	defer ticker.Stop()

	for range ticker.C {
		if p.IsPlaying() {
			continue
		}
		_ = p.Close()
		return
	}
}
