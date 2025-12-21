package character

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/xyy0411/ebiten_paractice/internal/physics"
)

func (c *Character) handleAttackInput() {
	if !ebiten.IsKeyPressed(ebiten.KeyJ) {
		return
	}

	if c.Action == ActionIdle || c.Action == ActionRun {
		c.SwitchAction(ActionJ1)
		return
	}

	def, ok := actionTable[c.Action]
	if !ok || def.CancelStart < 0 {
		return
	}

	// 是否在可取消帧
	if c.ActionFrame >= def.CancelStart && c.ActionFrame <= def.CancelEnd {
		c.SwitchAction(def.NextOnJ)
	}

	if c.ActionFrame >= def.TotalFrames {
		c.SwitchAction(ActionIdle)
	}
}

func (c *Character) updateActionEnd() {
	def, ok := actionTable[c.Action]
	if !ok {
		return
	}

	if c.ActionFrame >= def.TotalFrames {
		c.SwitchAction(ActionIdle)
	}
}

func (c *Character) handleInput() {
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		c.Vx = -c.Speed
		c.Facing = -1
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		c.Vx = c.Speed
		c.Facing = 1
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyK) && c.JumpCount < 2 {
		c.Vy = physics.JumpSpeed
		c.OnGround = false
		c.JumpCount++
	}

	if ebiten.IsKeyPressed(ebiten.KeyL) {
		c.Vx += float64(c.Facing) * c.Speed * 4
	}
}

func getDirInput(c *Character) DirInput {
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		return DirUp
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		return DirDown
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) && c.Facing == -1 || ebiten.IsKeyPressed(ebiten.KeyD) && c.Facing == 1 {
		return DirForward
	}
	return DirNone
}
