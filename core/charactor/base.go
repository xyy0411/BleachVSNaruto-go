package charactor

import (
	"github.com/xyy0411/bleachVSnaruto/common/state"
	"github.com/xyy0411/bleachVSnaruto/core/audio"
	"github.com/xyy0411/bleachVSnaruto/core/event"
)

type BaseCharacter struct {
	id       string
	name     string
	Runtime  *Runtime
	Data     *Data
	EventBus *event.Bus
}

func NewBaseCharacter(id, name string, runtime *Runtime, data *Data) *BaseCharacter {
	return &BaseCharacter{
		id:      id,
		name:    name,
		Runtime: runtime,
		Data:    data,
	}
}

func (b *BaseCharacter) GetData() *Data {
	return b.Data
}

func (b *BaseCharacter) GetRuntime() *Runtime {
	return b.Runtime
}

func (b *BaseCharacter) GetID() string {
	return b.id
}

func (b *BaseCharacter) GetName() string {
	return b.name
}

func (b *BaseCharacter) SetEventBus(bus *event.Bus) {
	b.EventBus = bus
}

// publishAudio 安全地发布音频事件，如果 EventBus 为 nil 则什么都不做
func (b *BaseCharacter) publishAudio(audioEvent audio.Event) {
	if b.EventBus != nil {
		b.EventBus.PublishAudio(b.id, audioEvent)
	}
}

// Update 使用模板方法模式，定义更新流程的骨架
// 角色可以重写特定步骤来实现个性化逻辑
func (b *BaseCharacter) Update() {
	// 步骤1：更新事件
	b.updateEvents()

	// 步骤2：处理特殊动作（钩子方法，角色可以重写）
	b.handleSpecialActions()

	// 步骤3：状态转换
	b.handleStateTransition()

	// 步骤4：处理音频事件
	b.handleAudioEvents()

	// 步骤5：处理移动锁定
	b.handleMovementLock()

	// 步骤6：处理朝向
	b.handleFacing()

	// 步骤7：播放动画
	b.playAnimation()

	// 步骤8：保存前一帧状态
	b.savePrevState()

	// 步骤9：更新后的处理（钩子方法，角色可以重写）
	b.handlePostUpdate()
}

// updateEvents 更新角色事件（可重写）
func (b *BaseCharacter) updateEvents() {
	body := b.Runtime.Body
	events := Events{}
	if !b.Runtime.PrevOnGround && body.OnGround && b.Runtime.PrevVY > 0 {
		events.JustLanded = true
	}
	if body.JumpsUsed > b.Runtime.PrevJumpsUsed && body.VY < 0 {
		events.JumpStart = true
	}
	b.Runtime.Events = events
}

// handleSpecialActions 处理特殊动作（钩子方法，角色可以重写）
// 例如：飞行物、特殊技能等
func (b *BaseCharacter) handleSpecialActions() {
	// 默认实现：什么都不做
	// 角色可以重写此方法来处理特殊动作
}

// handleStateTransition 处理状态转换（可重写）
func (b *BaseCharacter) handleStateTransition() {
	body := b.Runtime.Body

	jumpStartAnim := b.Data.Animations.ByState[state.JumpStart]
	justLandedAnim := b.Data.Animations.ByState[state.JustLanded]

	jumpStartLocked := jumpStartAnim != nil &&
		!jumpStartAnim.Loop &&
		b.Runtime.AnimPlayer.Current == jumpStartAnim &&
		b.Runtime.AnimPlayer.Frame < int64(len(jumpStartAnim.FramesKeys)-1)
	justLandedLocked := justLandedAnim != nil &&
		!justLandedAnim.Loop &&
		b.Runtime.AnimPlayer.Current == justLandedAnim &&
		b.Runtime.AnimPlayer.Frame < int64(len(justLandedAnim.FramesKeys)-1)

	switch {
	case justLandedLocked:
		b.Runtime.State = state.JustLanded
	case jumpStartLocked:
		b.Runtime.State = state.JumpStart
	case b.Runtime.Events.JustLanded:
		b.Runtime.State = state.JustLanded
	case b.Runtime.Events.JumpStart:
		b.Runtime.State = state.JumpStart
	case body.Dashing:
		b.Runtime.State = state.Dash
	case body.OnGround:
		if body.VX != 0 {
			b.Runtime.State = state.Run
		} else {
			b.Runtime.State = state.Idle
		}
	default:
		b.Runtime.State = state.Jump
	}
}

// handleAudioEvents 处理音频事件（可重写）
func (b *BaseCharacter) handleAudioEvents() {
	body := b.Runtime.Body

	// 处理落地音效
	if b.Runtime.Events.JustLanded {
		b.publishAudio(audio.EventJustLanded)
	}

	// 处理跳跃音效
	if b.Runtime.Events.JumpStart {
		b.publishAudio(audio.EventJumpStart)
	}

	// 处理冲刺音效
	if !b.Runtime.PrevDashed && body.Dashing {
		b.publishAudio(audio.EventDash)
	}

	// 处理跑步音效
	if body.OnGround && body.VX != 0 && b.Runtime.State == state.Run {
		b.Runtime.RunStepTimer++
		if b.Runtime.RunStepTimer >= 10 {
			b.publishAudio(audio.EventRunStep)
			b.Runtime.RunStepTimer = 0
		}
	} else {
		b.Runtime.RunStepTimer = 0
	}
}

// handleMovementLock 处理移动锁定（可重写）
func (b *BaseCharacter) handleMovementLock() {
	body := b.Runtime.Body
	lockMoveX := b.Runtime.State == state.JustLanded || b.Runtime.State == state.JumpStart
	if lockMoveX {
		body.X -= body.VX
		body.VX = 0
	}
}

// handleFacing 处理朝向（可重写）
func (b *BaseCharacter) handleFacing() {
	body := b.Runtime.Body
	lockMoveX := b.Runtime.State == state.JustLanded || b.Runtime.State == state.JumpStart

	if body.OnGround && !lockMoveX && !body.Dashing {
		if body.VX > 0 {
			b.Runtime.Facing = 1
		} else if body.VX < 0 {
			b.Runtime.Facing = -1
		}
	}
}

// playAnimation 播放动画（可重写）
func (b *BaseCharacter) playAnimation() {
	b.Runtime.AnimPlayer.Play(b.Data.Animations.ByState[b.Runtime.State])
}

// savePrevState 保存前一帧状态（可重写）
func (b *BaseCharacter) savePrevState() {
	body := b.Runtime.Body
	b.Runtime.PrevOnGround = body.OnGround
	b.Runtime.PrevVY = body.VY
	b.Runtime.PrevJumpsUsed = body.JumpsUsed
	b.Runtime.PrevDashed = body.Dashing
}

// handlePostUpdate 更新后的处理（钩子方法，角色可以重写）
// 例如：更新飞行物位置、处理特殊技能冷却等
func (b *BaseCharacter) handlePostUpdate() {
	// 默认实现：什么都不做
	// 角色可以重写此方法来处理更新后的逻辑
}
