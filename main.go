package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"image/color"
)

type Action int

const (
	ActionIdle Action = iota
	ActionRun
	ActionJump
	// ActionBlink 瞬步
	ActionBlink
	ActionAttackJ
	// ActionAttackADJ 投技
	ActionAttackADJ
	ActionAttackSJ
	ActionAttackWJ
	ActionAttackU
	// ActionAttackADU u投
	ActionAttackADU
	ActionAttackWU
)

type ActionDef struct {
	TotalFrames int

	StartupEnd  int // 前摇结束帧
	ActiveEnd   int // 命中结束帧
	RecoveryEnd int // 后摇结束帧（= TotalFrames）

	CancelWindows []CancelWindow
}

type CancelWindow struct {
	Start int
	End   int
	To    []Action // 允许取消到哪些动作
}

// CharState 角色状态
type CharState int

const (
	StateIdle CharState = iota // 站立
	StateRun                   // 跑动
	StateJump                  // 跳跃 / 空中状态
	StateAttack
)

type DirInput int

const (
	DirNone DirInput = iota
	DirUp
	DirDown
	DirForward
)

type AttackKey int

const (
	AttackJ AttackKey = iota
	AttackU
	AttackI
)

const (
	// Gravity 重力加速度
	// 重力为每帧速度的变化量
	Gravity = 0.6
	// JumpSpeed 跳跃速度
	JumpSpeed = -12
	// GroundY 地面高度
	GroundY = 500
	// AttackDuration 总共 20 帧
	AttackDuration = 20
)

type SkillKey struct {
	Dir DirInput
	Key AttackKey
}

type Skill struct {
	Frame int
	// other...
}

// Character 角色
type Character struct {
	Health float64 // 血量
	Height float64 // 身高
	Width  float64 // 宽度

	X, Y   float64   // 位置
	Speed  float64   // 移动速度
	Facing int       // 面向方向：1 右，-1 左
	State  CharState // 当前状态

	Vx float64 // 水平速度
	Vy float64 // 垂直速度（跳跃 / 下落）

	Skills map[SkillKey]Skill

	OnGround  bool // 是否在地面上
	JumpCount int  // 记录跳跃的次数用来判断是否可以进行二段跳

	Anim *AnimationManager

	Attacking   bool
	AttackTimer int
}

func drawGroundLine(screen *ebiten.Image) {
	ebitenutil.DrawLine(
		screen,
		0, GroundY,
		800, GroundY,
		color.RGBA{R: 255, A: 255},
	)
}

func getDirInput(c *Character) DirInput {
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		return DirUp
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		return DirDown
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) && c.Facing == -1 ||
		ebiten.IsKeyPressed(ebiten.KeyD) && c.Facing == 1 {
		return DirForward
	}
	return DirNone
}

func (c *Character) UseSkill(s Skill) {

}

func (c *Character) handleInput() {
	// 左右移动
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		c.Vx = -c.Speed
		c.Facing = -1
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		c.Vx = c.Speed
		c.Facing = 1
	}

	if ebiten.IsKeyPressed(ebiten.KeyW) {

	}

	if ebiten.IsKeyPressed(ebiten.KeyS) {

	}

	// 跳跃
	if inpututil.IsKeyJustPressed(ebiten.KeyK) && c.JumpCount < 2 {
		c.Vy = JumpSpeed
		c.OnGround = false
		c.JumpCount++
	}

	// 带普通攻击的技能组合
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

	// 瞬步也就是突然位移
	if ebiten.IsKeyPressed(ebiten.KeyL) {
		c.Vx += float64(c.Facing) * c.Speed * 4
	}
}

func (c *Character) updateState() {
	// 攻击状态优先
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

	// 空中状态
	if !c.OnGround {
		c.State = StateJump
		return
	}

	// 此时已经在地面了所以清零跳跃次数
	c.JumpCount = 0

	// 地面状态
	if c.Vx != 0 {
		c.State = StateRun
	} else {
		c.State = StateIdle
	}
}

func (c *Character) applyPhysics() {
	// 记录上一帧脚底位置
	prevY := c.Y

	// 重力
	if !c.OnGround {
		c.Vy += Gravity
	}

	// 位置更新
	c.X += c.Vx
	c.Y += c.Vy

	// 先假设不在地面
	c.OnGround = false

	if c.Y >= GroundY {
		c.Y = GroundY
		c.Vy = 0
		c.OnGround = true
	}

	for _, p := range platforms {
		// 横向范围内（脚底在平台上方）
		if c.X >= p.X && c.X <= p.X+p.W {
			// 必须是“从上往下落”
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
		c.Anim.Play(AnimIdle)

	case StateRun:
		c.Anim.Play(AnimRun)

	case StateJump:
		c.Anim.Play(AnimJump)

	case StateAttack:
		c.Anim.Play(AnimAttack)
	}

	c.Anim.Update()

	// 攻击动画播完，回到正常状态
	if c.State == StateAttack && c.Anim.IsFinished() {
		c.State = StateIdle
	}
}

// Update 更新角色逻辑（根据按键修改位置和状态）
func (c *Character) Update() {
	c.Vx = 0
	c.handleInput()
	c.updateState()
	c.applyPhysics()
	/*	if c.State == StateAttack {
		return
	}*/

	c.updateAnimation()
}

// Draw 绘制角色
func (c *Character) Draw(screen *ebiten.Image) {
	// 地面
	drawGroundLine(screen)

	// 台阶
	drawPlatforms(screen)

	ebitenutil.DebugPrintAt(screen, fmt.Sprintf(
		"X: %.2f  Vx: %.2f Y:%.2f Vy: %.2f  State: %d",
		c.X, c.Vx, c.Y, c.Vy, c.State,
	), 400, 300)

	frame := c.Anim.CurrentFrame()
	if frame == nil {
		return
	}

	drawSprite(screen, frame, c.X, c.Y, c.Facing)

}

// Game 游戏主结构
type Game struct {
	player *Character
}

// NewGame 初始化
func NewGame() *Game {
	loadAssets()

	player := NewCharacter(200, GroundY)

	return &Game{player: player}
}

func NewCharacter(x, y float64) *Character {
	am := NewAnimationManager()

	am.Add(AnimIdle, &Animation{
		Frames: assets.IdleFrames,
		Speed:  10,
		Loop:   true,
	})

	am.Add(AnimAttack, &Animation{
		Frames: assets.AttackFrames,
		Speed:  4,
		Loop:   false,
	})

	am.Add(AnimRun, &Animation{
		Frames: assets.RunFrames,
		Speed:  6,
		Loop:   true,
	})

	am.Add(AnimJump, &Animation{
		Frames: assets.JumpFrames,
		Speed:  7,
		Loop:   false,
	})

	am.Play(AnimIdle)

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

type Assets struct {
	IdleFrames   []*ebiten.Image
	RunFrames    []*ebiten.Image
	AttackFrames []*ebiten.Image
	JumpFrames   []*ebiten.Image
	LandFrames   []*ebiten.Image
}

var assets Assets

// Update 每帧更新逻辑
func (g *Game) Update() error {
	g.player.Update()
	return nil
}

// Draw 每帧绘制
func (g *Game) Draw(screen *ebiten.Image) {
	// 填个背景颜色
	screen.Fill(color.RGBA{R: 30, G: 30, B: 40, A: 255})

	// 画角色
	g.player.Draw(screen)

	// 左上角输出点调试文字
	ebitenutil.DebugPrint(screen, "← → 移动角色，蓝色=Idle，绿色=Run")

	// 显示 FPS / TPS
	ebitenutil.DebugPrint(
		screen,
		"FPS: "+fmt.Sprintf("%.2f", ebiten.ActualFPS())+
			"\nTPS: "+fmt.Sprintf("%.2f", ebiten.ActualTPS()),
	)

}

// Layout 设置逻辑画布大小
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 800, 600
}

func main() {

	game := NewGame()

	ebiten.SetWindowSize(800, 600)
	ebiten.SetWindowTitle("Ebiten 角色移动 + Idle/Run Demo")

	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
