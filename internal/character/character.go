package character

import (
	"fmt"

	"ebiten_paractice/internal/animation"
	"ebiten_paractice/internal/app"
	"ebiten_paractice/internal/physics"
	"ebiten_paractice/internal/render"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// Character 角色实体
type Character struct {
	Health float64 // 角色当前生命值
	Height float64 // 角色身高
	Width  float64 // 角色宽度

	X float64 // 角色位置 X
	Y float64 // 角色位置 Y

	Speed  float64 // 水平移动速度
	Facing int     // 朝向，1 向右，-1 向左

	State CharState // 当前动作状态

	Vx float64 // 水平速度分量
	Vy float64 // 垂直速度分量

	Skills map[SkillKey]Skill // 技能映射表

	OnGround  bool // 是否在地面上
	JumpCount int  // 已跳跃次数

	Anim *animation.Manager // 动画管理器

	Attacking   bool // 是否处于攻击过程
	AttackTimer int  // 攻击帧计数
}

// NewCharacter 创建角色并绑定默认动画
func NewCharacter(x, y float64) *Character {
	am := animation.NewManager()

	am.Add(animation.AnimIdle, &animation.Animation{Frames: render.Assets.IdleFrames, Speed: 10, Loop: true})
	am.Add(animation.AnimAttack, &animation.Animation{Frames: render.Assets.AttackFrames, Speed: 4, Loop: false})
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
	}
}

// Update 更新角色的输入、物理和动画
func (c *Character) Update() {
	c.Vx = 0
	c.handleInput()
	c.updateState()
	c.applyPhysics()
	c.updateAnimation()
}

// Draw 绘制角色和环境调试信息
func (c *Character) Draw(screen *ebiten.Image) {
	app.DrawGroundLine(screen)
	app.DrawPlatforms(screen)

	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("X: %.2f  Vx: %.2f Y:%.2f Vy: %.2f  State: %d", c.X, c.Vx, c.Y, c.Vy, c.State), 400, 300)

	frame := c.Anim.CurrentFrame()
	if frame == nil {
		return
	}

	render.DrawSprite(screen, frame, c.X, c.Y, c.Facing)
}

func (c *Character) updateState() {
	if c.State == StateAttack {
		c.AttackTimer++
		c.Vx = 0

		if c.AttackTimer >= AttackDuration {
			c.AttackTimer = 0
			c.Attacking = false

			if c.OnGround {
				c.State = StateIdle
			} else {
				c.State = StateJump
			}
		}
		return
	}

	if !c.OnGround {
		c.State = StateJump
		return
	}

	c.JumpCount = 0

	if c.Vx != 0 {
		c.State = StateRun
	} else {
		c.State = StateIdle
	}
}

func (c *Character) applyPhysics() {
	prevY := c.Y

	if !c.OnGround {
		c.Vy += physics.Gravity
	}

	c.X += c.Vx
	c.Y += c.Vy
	c.OnGround = false

	if c.Y >= app.GroundY {
		c.Y = app.GroundY
		c.Vy = 0
		c.OnGround = true
	}

	for _, p := range app.Platforms {
		if c.X >= p.X && c.X <= p.X+p.W {
			if prevY <= p.Y && c.Y >= p.Y {
				c.Y = p.Y
				c.Vy = 0
				c.OnGround = true
			}
		}
	}
}

func (c *Character) updateAnimation() {
	switch c.State {
	case StateIdle:
		c.Anim.Play(animation.AnimIdle)
	case StateRun:
		c.Anim.Play(animation.AnimRun)
	case StateJump:
		c.Anim.Play(animation.AnimJump)
	case StateAttack:
		c.Anim.Play(animation.AnimAttack)
	}

	c.Anim.Update()

	if c.State == StateAttack && c.Anim.IsFinished() {
		c.State = StateIdle
	}
}

// UseSkill 预留给技能系统的入口
func (c *Character) UseSkill(_ Skill) {}
