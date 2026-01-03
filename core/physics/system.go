package physics

import (
	"github.com/xyy0411/bleachVSnaruto/core/controller"
	"github.com/xyy0411/bleachVSnaruto/core/world"
	"github.com/xyy0411/bleachVSnaruto/models"
)

type System struct {
	Controller *controller.System
	World      *world.World

	Bodies []*models.PhysicsBody

	Gravity   float64
	MoveSpeed float64
	JumpSpeed float64
	DashSpeed float64
}

func (s *System) Name() string {
	return "physics"
}

func (s *System) Update() {
	intent := s.Controller.Current

	for _, body := range s.Bodies {
		body.Dashing = false

		if intent.DashHeld && intent.MoveX != 0 && body.OnGround {
			body.Dashing = true
			body.VX = float64(intent.MoveX) * s.DashSpeed
		} else {
			body.VX = float64(intent.MoveX) * s.MoveSpeed
		}

		if intent.JumpPressed && body.OnGround {
			body.VY = -s.JumpSpeed
			body.OnGround = false
		}

		if !body.OnGround {
			body.VY += s.Gravity
		}

		body.X += body.VX
		body.Y += body.VY

		body.Y, body.OnGround = s.World.ResolveGround(body.Y)

		if body.OnGround {
			body.VY = 0
		}
	}
}
