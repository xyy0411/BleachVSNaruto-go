package charactor

import (
	"github.com/xyy0411/bleachVSnaruto/core/action"
	"github.com/xyy0411/bleachVSnaruto/render/animation"
)

type Character interface {
	Update()
	GetRuntime() *Runtime
	GetAction() *action.Runtime
	GetID() string
	GetName() string
	GetData() *Data
}

type Data struct {
	MoveSpeed  float64
	JumpPower  float64
	Animations animation.Set
	Audio      AudioConfig
}

type AudioConfig struct {
	SFX    map[string]string // 事件名 -> 音频路径
	Volume float64
}
