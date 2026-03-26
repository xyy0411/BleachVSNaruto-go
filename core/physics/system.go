package physics

import (
	"github.com/xyy0411/bleachVSnaruto/core/controller"
	gametime "github.com/xyy0411/bleachVSnaruto/core/time"
	"github.com/xyy0411/bleachVSnaruto/core/world"
)

// 这里假设Controller与Bodies对应
type System struct {
	Controller []*controller.System
	World      *world.World
	Time       *gametime.Time

	Gravity   float64
	MoveSpeed float64
	JumpSpeed float64
	DashSpeed float64
}

func (s *System) Name() string {
	return "physics"
}

func (s *System) Update() {
	for i, player := range s.Controller {
		intent := player.Current
		delta := s.Time.Delta
		body := player.Body
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
			body.DashDirection = -2*i + 1
		}

		if !body.Dashing && intent.DashPressed && body.OnGround {
			body.Dashing = true
			body.DashTimer = body.DashDuration
			if intent.MoveX != 0 {
				body.DashDirection = intent.MoveX
			}
		}

		if body.Dashing {
			baseVX := (float64(body.DashDirection) * s.DashSpeed * s.MoveSpeed) / 2.0
			if body.DashDuration > 0 {
				progress := 1 - (body.DashTimer / body.DashDuration) // 0..1
				var speedScale float64
				if progress < 0.5 {
					speedScale = progress * 2 // accelerate
				} else {
					speedScale = (1 - progress) * 2 // decelerate
				}
				body.VX = baseVX * speedScale
				minDashVX := float64(body.DashDirection) * (s.MoveSpeed * 1.01)
				if body.VX*float64(body.DashDirection) < s.MoveSpeed*1.01 {
					body.VX = minDashVX
				}
			} else {
				body.VX = baseVX
			}
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
