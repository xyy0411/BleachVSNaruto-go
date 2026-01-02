package charactor

import (
	"github.com/xyy0411/ebiten_paractice/core/action"
	"github.com/xyy0411/ebiten_paractice/render/animation"
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
}
