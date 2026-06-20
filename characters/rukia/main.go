package rukia

import (
	"github.com/xyy0411/bleachVSnaruto/characters"
	"github.com/xyy0411/bleachVSnaruto/common/state"
	"github.com/xyy0411/bleachVSnaruto/core/animatable"
	"github.com/xyy0411/bleachVSnaruto/core/charactor"
	"github.com/xyy0411/bleachVSnaruto/core/event"
	"github.com/xyy0411/bleachVSnaruto/models"
)

const RoleID = "rukia"

func init() {
	characters.AddChar(RoleID, New)
}

type Rukia struct {
	id       string
	name     string
	Runtime  *charactor.Runtime
	Data     *charactor.Data
	EventBus *event.Bus
}

func New() charactor.Character {
	body := &models.PhysicsBody{
		X:            100,
		Y:            500,
		OnGround:     true,
		DashDuration: 0.3,
		MaxJumps:     2,
	}

	data := &charactor.Data{
		MoveSpeed:  4,
		JumpPower:  12,
		Animations: buildAnimations(),
		Audio:      charactor.DefaultAudioConfig(),
	}

	player := animatable.Player{}

	rt := &charactor.Runtime{
		Body:          body,
		Facing:        1,
		PrevOnGround:  body.OnGround,
		PrevVY:        body.VY,
		PrevJumpsUsed: body.JumpsUsed,
		AnimPlayer:    player,
	}

	return &Rukia{
		id:      "rukia",
		name:    "朽木露琪亚",
		Runtime: rt,
		Data:    data,
	}
}

func (r Rukia) GetData() *charactor.Data {
	return r.Data
}

func (r Rukia) GetRuntime() *charactor.Runtime {
	return r.Runtime
}

func (r Rukia) GetID() string {
	return r.id
}

func (r Rukia) GetName() string {
	return r.name
}

func (r *Rukia) SetEventBus(bus *event.Bus) {
	r.EventBus = bus
}

func (r Rukia) Update() {
	physicsBody := r.Runtime.Body
	events := charactor.Events{}
	if !r.Runtime.PrevOnGround && physicsBody.OnGround && r.Runtime.PrevVY > 0 {
		events.JustLanded = true
	}
	if physicsBody.JumpsUsed > r.Runtime.PrevJumpsUsed && physicsBody.VY < 0 {
		events.JumpStart = true
	}
	r.Runtime.Events = events

	jumpStartAnim := r.Data.Animations.ByState[state.JumpStart]
	justLandedAnim := r.Data.Animations.ByState[state.JustLanded]

	jumpStartLocked := jumpStartAnim != nil &&
		!jumpStartAnim.Loop &&
		r.Runtime.AnimPlayer.Current == jumpStartAnim &&
		r.Runtime.AnimPlayer.Frame < int64(len(jumpStartAnim.FramesKeys)-1)
	justLandedLocked := justLandedAnim != nil &&
		!justLandedAnim.Loop &&
		r.Runtime.AnimPlayer.Current == justLandedAnim &&
		r.Runtime.AnimPlayer.Frame < int64(len(justLandedAnim.FramesKeys)-1)

	switch {
	case justLandedLocked:
		r.Runtime.State = state.JustLanded
	case jumpStartLocked:
		r.Runtime.State = state.JumpStart
	case events.JustLanded:
		r.Runtime.State = state.JustLanded
	case events.JumpStart:
		r.Runtime.State = state.JumpStart
	case physicsBody.Dashing:
		r.Runtime.State = state.Dash
	case physicsBody.OnGround:
		if physicsBody.VX != 0 {
			r.Runtime.State = state.Run
		} else {
			r.Runtime.State = state.Idle
		}
	default:
		r.Runtime.State = state.Jump
	}

	lockMoveX := r.Runtime.State == state.JustLanded || r.Runtime.State == state.JumpStart
	if lockMoveX {
		physicsBody.X -= physicsBody.VX
		physicsBody.VX = 0
	}

	if physicsBody.OnGround && !lockMoveX && !physicsBody.Dashing {
		if physicsBody.VX > 0 {
			r.Runtime.Facing = 1
		} else if physicsBody.VX < 0 {
			r.Runtime.Facing = -1
		}
	}

	r.Runtime.AnimPlayer.Play(r.Data.Animations.ByState[r.Runtime.State])

	r.Runtime.PrevOnGround = physicsBody.OnGround
	r.Runtime.PrevVY = physicsBody.VY
	r.Runtime.PrevJumpsUsed = physicsBody.JumpsUsed
}
