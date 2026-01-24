package physics

import (
	"github.com/xyy0411/bleachVSnaruto/core/controller"
	gametime "github.com/xyy0411/bleachVSnaruto/core/time"
	"github.com/xyy0411/bleachVSnaruto/core/world"
	"github.com/xyy0411/bleachVSnaruto/models"
)

type System struct {
	Controller *controller.System
	World      *world.World
	Time       *gametime.Time

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
	delta := s.Time.Delta

	for _, body := range s.Bodies {
		if body.MaxJumps <= 0 {
			body.MaxJumps = 1
		}

		// 更新冲刺计时，并在未结束前保持冲刺状态
		if body.DashTimer > 0 {
			body.DashTimer -= delta
			if body.DashTimer <= 0 {
				body.DashTimer = 0
				body.Dashing = false
			} else {
				body.Dashing = true
			}
		} else {
			body.Dashing = false
		}

		// 只有在冲刺结束后且重新按下冲刺键时才可重新冲刺
		if intent.MoveX != 0 {
			body.DashDirection = intent.MoveX
		} else if body.DashDirection == 0 {
			body.DashDirection = 1
		}

		if !body.Dashing && intent.DashPressed && body.OnGround {
			body.Dashing = true
			body.DashTimer = body.DashDuration
			if intent.MoveX != 0 {
				body.DashDirection = intent.MoveX
			}
		}

		if body.Dashing {
			body.VX = (float64(body.DashDirection) * s.DashSpeed * s.MoveSpeed) / 3.0
		} else {
			body.VX = float64(intent.MoveX) * s.MoveSpeed
		}

		canJump := intent.JumpPressed && (body.OnGround || body.JumpsUsed < body.MaxJumps)
		if canJump {
			body.VY = -s.JumpSpeed
			body.OnGround = false
			body.JumpsUsed++
		}

		if !body.OnGround {
			body.VY += s.Gravity
		}

		body.X += body.VX
		body.Y += body.VY

		body.Y, body.OnGround = s.World.ResolveGround(body.Y)

		if body.OnGround {
			body.VY = 0
			body.JumpsUsed = 0
		}
	}
}
