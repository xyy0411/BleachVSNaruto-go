package charactor

import (
	"github.com/xyy0411/ebiten_paractice/common/state"
	"github.com/xyy0411/ebiten_paractice/core/animatable"
	"github.com/xyy0411/ebiten_paractice/models"
	"github.com/xyy0411/ebiten_paractice/render/animation"
)

type Runtime struct {
	Body *models.PhysicsBody

	State state.State
	// 朝左 1 朝右 -1
	Facing int

	AnimPlayer animation.Player
}

func (r *Runtime) AnimationPlayer() animatable.AnimationPlayer {
	return &r.AnimPlayer
}
