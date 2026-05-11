package charactor

import (
	"github.com/xyy0411/bleachVSnaruto/common/state"
	"github.com/xyy0411/bleachVSnaruto/core/animatable"
	"github.com/xyy0411/bleachVSnaruto/core/audio"
	"github.com/xyy0411/bleachVSnaruto/models"
)

// Runtime 保存角色在运行时的物理状态 动画状态和音频数据
type Runtime struct {
	Body *models.PhysicsBody
	// Intent 保存当前帧由控制层生成的角色行为意图
	Intent models.Intent

	State state.State
	// Facing 表示角色朝向 -1 为朝左 1 为朝右
	Facing        int
	PrevOnGround  bool
	PrevVY        float64
	PrevJumpsUsed int
	PrevDashed    bool
	Events        Events

	// RunStepTimer 控制跑步音效的触发节奏
	RunStepTimer float64
	// LastAudioVariant 记录每类音效上次播放到的片段索引
	LastAudioVariant map[audio.Event]int

	AnimPlayer  animatable.Player
	AudioEvents []audio.Event
}

// AnimationPlayer 返回角色当前使用的动画播放器
func (r *Runtime) AnimationPlayer() animatable.AnimationPlayer {
	return &r.AnimPlayer
}

// Events 保存角色本帧产生的瞬时事件
type Events struct {
	JumpStart  bool
	JustLanded bool
}
