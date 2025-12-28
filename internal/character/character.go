package character

import (
	"github.com/xyy0411/ebiten_paractice/internal/animation"
	"github.com/xyy0411/ebiten_paractice/internal/physics"
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

	Action           Action // 当前动作
	ActionFrame      int64  // 动作帧数
	ActionStartFrame int64  // 动作开始帧数

	Env Environment // 环境信息
}

func (c *Character) basicOperations() {
	c.ActionFrame = GlobalFrame - c.ActionStartFrame
}

func (c *Character) updateState() {

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

func (c *Character) applyActionDamping() {
	def, ok := actionTable[c.Action]
	if !ok {
		return
	}

	switch def.Phase(c.ActionFrame) {
	case PhaseStartup:
		c.Vx *= def.DampingStartup
	case PhaseActive:
		c.Vx *= def.DampingActive
	case PhaseRecovery:
		c.Vx *= def.DampingRecovery
	}
}

func (c *Character) CanHit() bool {
	def, ok := actionTable[c.Action]
	if !ok {
		return false
	}

	return c.ActionFrame >= def.HitStart &&
		c.ActionFrame <= def.HitEnd
}

func (c *Character) applyPhysics() {
	// prevY := c.Y

	c.applyActionDamping()

	if c.Env == nil {
		return
	}

	if !c.OnGround {
		c.Vy += physics.Gravity
	}

	c.X += c.Vx
	c.Y += c.Vy
	c.OnGround = false

	onGround, y := c.Env.CheckPlatform(c.X, c.Y, c.Vy)
	if onGround {
		c.Y = y
		c.Vy = 0
		c.OnGround = true
	}
}

func (c *Character) updateAnimation() {
	defer c.Anim.Update()

	if c.IsAttacking() {
		c.Anim.Play(actionToAnim(c.Action))
		return
	}

	switch c.State {
	case StateIdle:
		c.Anim.Play(animation.AnimIdle)
	case StateRun:
		c.Anim.Play(animation.AnimRun)
	case StateJump:
		c.Anim.Play(animation.AnimJump)
	}

	if c.State == StateAttack && c.Anim.IsFinished() {
		c.State = StateIdle
	}
}

// UseSkill 预留给技能系统的入口
func (c *Character) UseSkill(_ Skill) {}

func (c *Character) SwitchAction(a Action) {
	c.Action = a
	c.ActionStartFrame = GlobalFrame
}

func (c *Character) IsAttacking() bool {
	return c.Action >= ActionJ1 && c.Action <= ActionJ3
}
