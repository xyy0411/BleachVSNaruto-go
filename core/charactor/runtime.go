package charactor

import (
	"github.com/xyy0411/bleachVSnaruto/common/state"
	"github.com/xyy0411/bleachVSnaruto/core/animatable"
	"github.com/xyy0411/bleachVSnaruto/models"
)

type Runtime struct {
	Body *models.PhysicsBody

	State state.State
	// 朝左 -1 朝右 1
	Facing        int
	PrevOnGround  bool
	PrevVY        float64
	PrevJumpsUsed int
	Events        Events

	AnimPlayer animatable.Player
}

func (r *Runtime) AnimationPlayer() animatable.AnimationPlayer {
	return &r.AnimPlayer
}

type Events struct {
	JumpStart  bool
	JustLanded bool
}
