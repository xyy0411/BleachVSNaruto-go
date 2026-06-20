package charactor

import (
	"github.com/xyy0411/bleachVSnaruto/core/animatable"
	"github.com/xyy0411/bleachVSnaruto/core/audio"
	"github.com/xyy0411/bleachVSnaruto/core/event"
)

type Character interface {
	Update()
	GetRuntime() *Runtime
	GetID() string
	GetName() string
	GetData() *Data
	SetEventBus(bus *event.Bus)
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
