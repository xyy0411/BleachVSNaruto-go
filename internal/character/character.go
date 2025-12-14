package character

import (
	"fmt"

	"ebiten_paractice/internal/animation"
	"ebiten_paractice/internal/assets"
	"ebiten_paractice/internal/world"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// Action 描述角色的动作定义
type Action int

const (
	ActionIdle Action = iota
	ActionRun
	ActionJump
	ActionBlink
	ActionAttackJ
	ActionAttackADJ
	ActionAttackSJ
	ActionAttackWJ
	ActionAttackU
	ActionAttackADU
	ActionAttackWU
)

// ActionDef 控制动作帧与取消窗口
type ActionDef struct {
	TotalFrames   int            // 总帧数
	StartupEnd    int            // 起手帧结束位置
	ActiveEnd     int            // 攻击判定结束位置
	RecoveryEnd   int            // 硬直结束位置
	CancelWindows []CancelWindow // 允许取消的区间
}

// CancelWindow 定义动作取消的可行区间
type CancelWindow struct {
	Start int      // 取消窗口起始帧
	End   int      // 取消窗口结束帧
	To    []Action // 可取消到的动作列表
}

// CharState 角色状态枚举
type CharState int

const (
	StateIdle CharState = iota
	StateRun
	StateJump
	StateAttack
)

// DirInput 表示方向键输入
type DirInput int

const (
	DirNone DirInput = iota
	DirUp
	DirDown
	DirForward
)

// AttackKey 普通攻击按键
type AttackKey int

const (
	AttackJ AttackKey = iota
	AttackU
	AttackI
)

const (
	Gravity        = 0.6 // 重力加速度，每帧速度的变化量
	JumpSpeed      = -12 // 跳跃初速度
	AttackDuration = 20  // 攻击持续帧数
)

// SkillKey 用于匹配招式表
type SkillKey struct {
	Dir DirInput  // 方向输入
	Key AttackKey // 具体按键
}

// Skill 预留的技能定义结构
type Skill struct {
	Frame int // 触发帧数
}

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

	am.Add(animation.AnimIdle, &animation.Animation{Frames: assets.Assets.IdleFrames, Speed: 10, Loop: true})
	am.Add(animation.AnimAttack, &animation.Animation{Frames: assets.Assets.AttackFrames, Speed: 4, Loop: false})
	am.Add(animation.AnimRun, &animation.Animation{Frames: assets.Assets.RunFrames, Speed: 6, Loop: true})
	am.Add(animation.AnimJump, &animation.Animation{Frames: assets.Assets.JumpFrames, Speed: 7, Loop: false})

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
	world.DrawGroundLine(screen)
	world.DrawPlatforms(screen)

	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("X: %.2f  Vx: %.2f Y:%.2f Vy: %.2f  State: %d", c.X, c.Vx, c.Y, c.Vy, c.State), 400, 300)

	frame := c.Anim.CurrentFrame()
	if frame == nil {
		return
	}

	drawSprite(screen, frame, c.X, c.Y, c.Facing)
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
		c.Vy = JumpSpeed
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
		c.Vy += Gravity
	}

	c.X += c.Vx
	c.Y += c.Vy
	c.OnGround = false

	if c.Y >= world.GroundY {
		c.Y = world.GroundY
		c.Vy = 0
		c.OnGround = true
	}

	for _, p := range world.Platforms {
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

// drawSprite 根据朝向绘制一帧
func drawSprite(screen *ebiten.Image, img *ebiten.Image, x, y float64, facing int) {
	op := &ebiten.DrawImageOptions{}

	if facing == -1 {
		w := img.Bounds().Dx()
		op.GeoM.Scale(-1, 1)
		op.GeoM.Translate(float64(w), 0)
	}

	op.GeoM.Translate(x, y-float64(img.Bounds().Dy()))
	screen.DrawImage(img, op)
}
