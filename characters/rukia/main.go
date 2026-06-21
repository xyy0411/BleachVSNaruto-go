package rukia

import (
	"github.com/xyy0411/bleachVSnaruto/characters"
	"github.com/xyy0411/bleachVSnaruto/core/animatable"
	"github.com/xyy0411/bleachVSnaruto/core/charactor"
	"github.com/xyy0411/bleachVSnaruto/models"
)

const RoleID = "rukia"

func init() {
	characters.AddChar(RoleID, New)
}

type Rukia struct {
	*charactor.BaseCharacter
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

	base := charactor.NewBaseCharacter(RoleID, "朽木露琪亚", rt, data)

	return &Rukia{
		BaseCharacter: base,
	}
}
