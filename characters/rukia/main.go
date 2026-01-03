package rukia

import (
	"github.com/xyy0411/bleachVSnaruto/common/state"
	"github.com/xyy0411/bleachVSnaruto/core/action"
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
		X:        100,
		Y:        500,
		OnGround: true,
	}

	data := &charactor.Data{
		MoveSpeed:  4,
		JumpPower:  12,
		Animations: buildAnimations(),
	}

	player := animation.Player{}

	rt := &charactor.Runtime{
		Body:       body,
		Facing:     1,
		AnimPlayer: player,
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

func (r Rukia) Update() {
	body := r.Runtime.Body

	if body.Dashing {
		r.Runtime.State = state.Dash
	} else if body.OnGround {
		if body.VX != 0 {
			r.Runtime.State = state.Run
		} else {
			r.Runtime.State = state.Idle
		}
	} else {
		r.Runtime.State = state.Jump
	}

	if body.VX > 0 {
		r.Runtime.Facing = 1
	} else if body.VX < 0 {
		r.Runtime.Facing = -1
	}

	if anim := r.Data.Animations.ByState[r.Runtime.State]; anim != nil {
		r.Runtime.AnimPlayer.Play(anim)
	}
}

func (r Rukia) GetAction() *action.Runtime {
	return &action.Runtime{}
}
