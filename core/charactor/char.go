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
	animatable.Rect
	//受击框生效
	Action bool
}

func (r *Rect) Right() int {
	return r.X + r.W
}

func (r *Rect) Bottom() int {
	return r.Y + r.H
}

func (r *Rect) Intersects(other *Rect) bool {
	if !r.Action || !other.Action {
		return false
	}
	if r.Right() < other.X || other.Right() < r.X {
		return false
	}
	if r.Bottom() < other.Y || r.Y < other.Bottom() {
		return false
	}
	return true
}
