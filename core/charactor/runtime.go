package charactor

import (
	"github.com/xyy0411/bleachVSnaruto/common/state"
	"github.com/xyy0411/bleachVSnaruto/core/animatable"
	"github.com/xyy0411/bleachVSnaruto/core/audio"
	"github.com/xyy0411/bleachVSnaruto/models"
)

type Runtime struct {
	Body       *models.PhysicsBody
	BodyRect   *Rect //受击框
	EffectRect *Rect //特效框

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

// UpdataRect 更新角色框与特效框
func (r *Runtime) UpdataRect() {
	if r.State == state.Fall {
		r.BodyRect.Action = false
	}
	r.BodyRect.Left = r.Body.X
	r.BodyRect.Top = r.Body.Y
	//...特效框
}

type Events struct {
	JumpStart  bool
	JustLanded bool
}
