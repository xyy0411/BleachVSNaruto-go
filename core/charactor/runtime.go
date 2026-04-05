package charactor

import (
	"github.com/xyy0411/bleachVSnaruto/common/state"
	"github.com/xyy0411/bleachVSnaruto/core/animatable"
	"github.com/xyy0411/bleachVSnaruto/core/audio"
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
	PrevDashed    bool
	Events        Events

	// 用于处理跑步时的音效
	RunStepTimer float64
	// 记录事件上次音效播放到哪个片段
	LastAudioVariant map[audio.Event]int

	AnimPlayer  animatable.Player
	AudioEvents []audio.Event
}

func (r *Runtime) AnimationPlayer() animatable.AnimationPlayer {
	return &r.AnimPlayer
}

type Events struct {
	JumpStart  bool
	JustLanded bool
}
