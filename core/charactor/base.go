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
	b.UpdateEvents()
	b.HandleSpecialActions()
	b.HandleStateTransition()
	b.HandleAudioEvents()
	b.HandleMovementLock()
	b.HandleFacing()
	b.PlayAnimation()
	b.SavePrevState()
	b.HandlePostUpdate()
}

// UpdateEvents 更新角色事件
func (b *BaseCharacter) UpdateEvents() {
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

// HandleSpecialActions 处理特殊动作
// 例如：飞行物、特殊技能等
func (b *BaseCharacter) HandleSpecialActions() {
	// 处理技能动画帧 origin 变化导致的位移
	if b.isSkillState() {
		anim := b.Data.Animations.ByState[b.Runtime.State]
		if anim != nil && b.Runtime.AnimPlayer.Frame < int64(len(anim.Frames)) {
			frame := anim.Frames[b.Runtime.AnimPlayer.Frame]
			body := b.Runtime.Body
			// 计算当前帧与上一帧 origin.x 的差值
			deltaX := frame.Origin.X - b.Runtime.PrevOriginX
			// 根据角色朝向应用位移
			body.X += deltaX * float64(b.Runtime.Facing)
			// 记录当前帧的 origin.x
			b.Runtime.PrevOriginX = frame.Origin.X
		}
	} else {
		// 非技能状态时重置 PrevOriginX
		b.Runtime.PrevOriginX = 0
	}
}

// HandleStateTransition 处理状态转换
func (b *BaseCharacter) HandleStateTransition() {
	body := b.Runtime.Body

	jumpStartLocked := b.isAnimLocked(state.JumpStart)
	justLandedLocked := b.isAnimLocked(state.JustLanded)
	skillAnimLocked := b.isSkillAnimLocked()

	switch {
	case justLandedLocked:
		b.Runtime.State = state.JustLanded
	case jumpStartLocked:
		b.Runtime.State = state.JumpStart
	case b.Runtime.Events.JustLanded:
		b.Runtime.State = state.JustLanded
	case b.Runtime.Events.JumpStart:
		b.Runtime.State = state.JumpStart
	case skillAnimLocked:
		b.Runtime.Intent.ComboCommand = ""
	case b.Runtime.Intent.ComboCommand != "":
		b.transitionByCombo()
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

// isAnimLocked 检查指定状态的动画是否正在播放且未结束
func (b *BaseCharacter) isAnimLocked(s state.State) bool {
	anim := b.Data.Animations.ByState[s]
	if anim == nil {
		return false
	}
	return !anim.Loop &&
		b.Runtime.AnimPlayer.Current == anim &&
		b.Runtime.AnimPlayer.Frame < int64(len(anim.FramesKeys)-1)
}

// isSkillAnimLocked 检查是否有技能动画正在播放且未结束
func (b *BaseCharacter) isSkillAnimLocked() bool {
	for _, s := range state.GetSkill() {
		if b.isAnimLocked(s) {
			return true
		}
	}
	return false
}

// transitionByCombo 根据 ComboCommand 进行状态转换
func (b *BaseCharacter) transitionByCombo() {
	cmd := b.Runtime.Intent.ComboCommand

	stateMap := map[string]state.State{
		"wj": state.WJ, "sj": state.SJ, "kj": state.KJ,
		"wu": state.WU, "su": state.SU, "ku": state.KU,
		"wi": state.WI, "si": state.SI, "ki": state.KI,
		"i": state.NormalI,
	}

	if s, ok := stateMap[cmd]; ok {
		if b.Data.Animations.ByState[s] != nil {
			b.Runtime.State = s
		}
	}
	b.Runtime.Intent.ComboCommand = ""
}

// HandleAudioEvents 处理音频事件
func (b *BaseCharacter) HandleAudioEvents() {
	body := b.Runtime.Body

	// 落地音效
	if b.Runtime.Events.JustLanded {
		b.publishAudio(audio.EventJustLanded)
	}

	// 跳跃音效
	if b.Runtime.Events.JumpStart {
		b.publishAudio(audio.EventJumpStart)
	}

	// 冲刺音效
	if !b.Runtime.PrevDashed && body.Dashing {
		b.publishAudio(audio.EventDash)
	}

	// 跑步音效
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

// HandleMovementLock 处理移动锁定
func (b *BaseCharacter) HandleMovementLock() {
	body := b.Runtime.Body

	// 技能状态时锁定移动
	if b.isSkillState() {
		body.X -= body.VX
		body.VX = 0
		return
	}

	// 落地和跳跃起手时锁定移动
	lockMoveX := b.Runtime.State == state.JustLanded || b.Runtime.State == state.JumpStart
	if lockMoveX {
		body.X -= body.VX
		body.VX = 0
	}
}

// isSkillState 检查当前是否处于技能状态
func (b *BaseCharacter) isSkillState() bool {
	for _, s := range state.GetSkill() {
		if b.Runtime.State == s {
			return true
		}
	}
	return false
}

// HandleFacing 处理朝向
func (b *BaseCharacter) HandleFacing() {
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

// PlayAnimation 播放动画
func (b *BaseCharacter) PlayAnimation() {
	b.Runtime.AnimPlayer.Play(b.Data.Animations.ByState[b.Runtime.State])
}

// SavePrevState 保存前一帧状态
func (b *BaseCharacter) SavePrevState() {
	body := b.Runtime.Body
	b.Runtime.PrevOnGround = body.OnGround
	b.Runtime.PrevVY = body.VY
	b.Runtime.PrevJumpsUsed = body.JumpsUsed
	b.Runtime.PrevDashed = body.Dashing
}

// HandlePostUpdate 更新后的处理
// 例如：更新飞行物位置、处理特殊技能冷却等
func (b *BaseCharacter) HandlePostUpdate() {
	// 默认实现：什么都不做
	// 角色可以重写此方法来处理更新后的逻辑
}
