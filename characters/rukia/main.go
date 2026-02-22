package rukia

import (
	"github.com/xyy0411/bleachVSnaruto/common/state"
	"github.com/xyy0411/bleachVSnaruto/core/action"
	"github.com/xyy0411/bleachVSnaruto/core/audio"
	"github.com/xyy0411/bleachVSnaruto/core/charactor"
	"github.com/xyy0411/bleachVSnaruto/models"
	"github.com/xyy0411/bleachVSnaruto/render/animation"
)

type Rukia struct {
	id   string
	name string

	Runtime *charactor.Runtime
	Data    *charactor.Data
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

	player := animation.Player{}

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
		name:    "rukia",
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

func (r Rukia) Update() {
	body := r.Runtime.Body
	events := charactor.Events{}
	if !r.Runtime.PrevOnGround && body.OnGround && r.Runtime.PrevVY > 0 {
		events.JustLanded = true
	}
	if body.JumpsUsed > r.Runtime.PrevJumpsUsed && body.VY < 0 {
		events.JumpStart = true
	}
	r.Runtime.Events = events
	if events.JumpStart {
		if path := r.Data.Audio.SFX[audio.EventJump]; path != "" && audio.Default != nil {
			audio.Default.Play(path, r.Data.Audio.Volume)
		}
	}

	jumpStartAnim := r.Data.Animations.ByState[state.JumpStart]
	justLandedAnim := r.Data.Animations.ByState[state.JustLanded]

	jumpStartLocked := jumpStartAnim != nil &&
		!jumpStartAnim.Loop &&
		r.Runtime.AnimPlayer.Current == jumpStartAnim &&
		r.Runtime.AnimPlayer.Frame < int64(len(jumpStartAnim.Frames)-1)
	justLandedLocked := justLandedAnim != nil &&
		!justLandedAnim.Loop &&
		r.Runtime.AnimPlayer.Current == justLandedAnim &&
		r.Runtime.AnimPlayer.Frame < int64(len(justLandedAnim.Frames)-1)

	switch {
	case justLandedLocked:
		r.Runtime.State = state.JustLanded
	case jumpStartLocked:
		r.Runtime.State = state.JumpStart
	case events.JustLanded:
		r.Runtime.State = state.JustLanded
	case events.JumpStart:
		r.Runtime.State = state.JumpStart
	case body.Dashing:
		r.Runtime.State = state.Dash
	case body.OnGround:
		if body.VX != 0 {
			r.Runtime.State = state.Run
		} else {
			r.Runtime.State = state.Idle
		}
	default:
		r.Runtime.State = state.Jump
	}

	lockMoveX := r.Runtime.State == state.JustLanded || r.Runtime.State == state.JumpStart
	if lockMoveX {
		body.X -= body.VX
		body.VX = 0
	}

	if body.OnGround && !lockMoveX && !body.Dashing {
		if body.VX > 0 {
			r.Runtime.Facing = 1
		} else if body.VX < 0 {
			r.Runtime.Facing = -1
		}
	}

	r.Runtime.AnimPlayer.Play(r.Data.Animations.ByState[r.Runtime.State])

	r.Runtime.PrevOnGround = body.OnGround
	r.Runtime.PrevVY = body.VY
	r.Runtime.PrevJumpsUsed = body.JumpsUsed
}

func (r Rukia) GetAction() *action.Runtime {
	return &action.Runtime{}
}
