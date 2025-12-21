package character

import (
	"ebiten_paractice/internal/physics"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

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

	if inpututil.IsKeyJustPressed(ebiten.KeyJ) {
		dir := getDirInput(c)
		key := SkillKey{dir, AttackJ}
		if skill, ok := c.Skills[key]; ok {
			c.UseSkill(skill)
		}
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyJ) && c.State != StateAttack {
		c.State = StateAttack
		c.Attacking = true
		c.AttackTimer = 0
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
