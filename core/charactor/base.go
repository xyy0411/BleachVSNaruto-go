package charactor

import (
	"github.com/xyy0411/bleachVSnaruto/common/state"
	"github.com/xyy0411/bleachVSnaruto/core/audio"
	"github.com/xyy0411/bleachVSnaruto/core/event"
)

type BaseCharacter struct {
	id      string
	name    string
	Runtime *Runtime
	Data    *Data
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

func (b *BaseCharacter) Update() {
	body := b.Runtime.Body

	b.updateEvents()

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
		b.publishAudio(audio.EventJustLanded)
	case b.Runtime.Events.JumpStart:
		b.Runtime.State = state.JumpStart
		b.publishAudio(audio.EventJumpStart)
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

	if !b.Runtime.PrevDashed && body.Dashing {
		b.publishAudio(audio.EventDash)
	}

	if body.OnGround && body.VX != 0 && b.Runtime.State == state.Run {
		b.Runtime.RunStepTimer++
		if b.Runtime.RunStepTimer >= 10 {
			b.publishAudio(audio.EventRunStep)
			b.Runtime.RunStepTimer = 0
		}
	} else {
		b.Runtime.RunStepTimer = 0
	}

	b.handleMovementLock()
	b.handleFacing()

	b.Runtime.AnimPlayer.Play(b.Data.Animations.ByState[b.Runtime.State])

	b.savePrevState()
}

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

func (b *BaseCharacter) handleMovementLock() {
	body := b.Runtime.Body
	lockMoveX := b.Runtime.State == state.JustLanded || b.Runtime.State == state.JumpStart
	if lockMoveX {
		body.X -= body.VX
		body.VX = 0
	}
}

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

func (b *BaseCharacter) savePrevState() {
	body := b.Runtime.Body
	b.Runtime.PrevOnGround = body.OnGround
	b.Runtime.PrevVY = body.VY
	b.Runtime.PrevJumpsUsed = body.JumpsUsed
	b.Runtime.PrevDashed = body.Dashing
}
