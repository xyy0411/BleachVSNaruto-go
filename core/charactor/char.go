package charactor

import (
	"github.com/xyy0411/bleachVSnaruto/core/animatable"
	"github.com/xyy0411/bleachVSnaruto/core/audio"
)

type Character interface {
	Update()
	GetRuntime() *Runtime
	GetID() string
	GetName() string
	GetData() *Data
}

type Data struct {
	MoveSpeed  float64
	JumpPower  float64
	Animations animatable.Set
	Audio      AudioConfig
}

type AudioConfig struct {
	SFX    map[audio.Event][]string // 音效事件 -> 音频路径
	Volume float64
}
type Rect struct {
	Left   float64
	Top    float64
	Width  float64
	Height float64
	//受击框生效
	Action bool
}

func (r *Rect) Right() float64 {
	return r.Left + r.Width
}

func (r *Rect) Bottom() float64 {
	return r.Top + r.Height
}

func (r *Rect) Intersects(other *Rect) bool {
	return r.Left < other.Right() &&
		r.Right() > other.Left &&
		r.Top < other.Bottom() &&
		r.Bottom() > other.Top
}