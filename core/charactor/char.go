package charactor

import (
	"github.com/xyy0411/bleachVSnaruto/core/animatable"
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
	SFX    map[string]string // 事件名 -> 音频路径
	Volume float64
}
