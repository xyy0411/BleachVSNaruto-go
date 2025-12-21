package character

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/xyy0411/ebiten_paractice/internal/animation"
	"github.com/xyy0411/ebiten_paractice/internal/render"
	"github.com/xyy0411/ebiten_paractice/internal/world"
)

// NewCharacter 创建角色并绑定默认动画
func NewCharacter(x, y float64, e *world.Terrain) *Character {
	am := animation.NewManager()

	am.Add(animation.AnimJ1, &animation.Animation{
		Frames: render.Assets.AttackFrames.J.J[0],
		Speed:  4,
		Loop:   false,
	})

	am.Add(animation.AnimJ2, &animation.Animation{
		Frames: render.Assets.AttackFrames.J.J[1],
		Speed:  4,
		Loop:   false,
	})

	am.Add(animation.AnimJ3, &animation.Animation{
		Frames: render.Assets.AttackFrames.J.J[2],
		Speed:  4,
		Loop:   false,
	})

	am.Add(animation.AnimIdle, &animation.Animation{Frames: render.Assets.IdleFrames, Speed: 10, Loop: true})
	am.Add(animation.AnimRun, &animation.Animation{Frames: render.Assets.RunFrames, Speed: 6, Loop: true})
	am.Add(animation.AnimJump, &animation.Animation{Frames: render.Assets.JumpFrames, Speed: 7, Loop: false})

	am.Play(animation.AnimIdle)

	return &Character{
		X:        x,
		Y:        y,
		State:    StateIdle,
		Speed:    4,
		Facing:   1,
		OnGround: true,
		Anim:     am,
		Env:      e,
	}
}

// Update 更新角色的输入、物理和动画
func (c *Character) Update() {
	c.basicOperations()
	c.Vx = 0
	if !c.IsAttacking() {
		c.handleInput()
	}

	c.handleAttackInput()
	c.updateActionEnd()
	c.updateState()
	c.applyPhysics()
	c.updateAnimation()
}

// Draw 绘制角色和环境调试信息
func (c *Character) Draw(screen *ebiten.Image) {
	world.DrawGroundLine(screen)
	world.DrawPlatforms(screen)

	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("X: %.2f  Vx: %.2f Y:%.2f Vy: %.2f  State: %d Frame: %d", c.X, c.Vx, c.Y, c.Vy, c.State, c.ActionFrame), 400, 300)

	frame := c.Anim.CurrentFrame()
	if frame == nil {
		return
	}

	render.DrawSprite(screen, frame, c.X, c.Y, c.Facing)
}
