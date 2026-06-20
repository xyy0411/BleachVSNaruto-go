package ichigo

import (
	"github.com/xyy0411/bleachVSnaruto/characters"
	"github.com/xyy0411/bleachVSnaruto/common/state"
	"github.com/xyy0411/bleachVSnaruto/core/animatable"
	"github.com/xyy0411/bleachVSnaruto/core/charactor"
	"github.com/xyy0411/bleachVSnaruto/models"
)

const RoleID = "ichigo"

func init() {
	characters.AddChar(RoleID, New)
}

// Ichigo 黑崎一护角色
// 这个示例展示如何重写 BaseCharacter 的特定步骤来实现个性化逻辑
type Ichigo struct {
	*charactor.BaseCharacter

	// 特殊能力系统
	ProjectileExtension *charactor.ProjectileCharacterExtension

	// 卍解状态
	BankaiMode      bool
	BankaiTimer     int
	BankaiMaxTimer  int

	// 虚化状态
	VizardMode      bool
	VizardTimer     int
	VizardMaxTimer  int

	// 技能冷却
	GetsugaCooldown int
	GetsugaMaxCooldown int
}

func New() charactor.Character {
	body := &models.PhysicsBody{
		X:            100,
		Y:            500,
		OnGround:     true,
		DashDuration: 0.3,
		MaxJumps:     2,
	}

	data := &charactor.Data{
		MoveSpeed:  4,
		JumpPower:  12,
		Animations: buildAnimations(),
		Audio:      charactor.DefaultAudioConfig(),
	}

	rt := &charactor.Runtime{
		Body:          body,
		Facing:        1,
		PrevOnGround:  body.OnGround,
		PrevVY:        body.VY,
		PrevJumpsUsed: body.JumpsUsed,
		AnimPlayer:    animatable.Player{},
	}

	base := charactor.NewBaseCharacter(RoleID, "黑崎一护", rt, data)

	return &Ichigo{
		BaseCharacter:       base,
		ProjectileExtension: charactor.NewProjectileCharacterExtension(60), // 1秒冷却
		BankaiMode:          false,
		BankaiMaxTimer:      600, // 10秒卍解时间
		VizardMode:          false,
		VizardMaxTimer:      300, // 5秒虚化时间
		GetsugaCooldown:     0,
		GetsugaMaxCooldown:  45, // 0.75秒冷却
	}
}

// handleSpecialActions 重写：处理特殊动作
// 这是钩子方法，在状态转换之前执行
// 适合处理：技能释放、飞行物发射、模式切换等
func (i *Ichigo) handleSpecialActions() {
	body := i.Runtime.Body
	intent := i.Runtime.Intent

	// 1. 处理月牙天冲（飞行物技能）
	if intent.AttackPressed && i.GetsugaCooldown <= 0 && !body.Dashing {
		i.launchGetsuga()
		i.GetsugaCooldown = i.GetsugaMaxCooldown
	}

	// 2. 处理卍解模式切换
	if intent.StatePressed == state.Bankai && !i.BankaiMode && i.BankaiTimer == 0 {
		i.activateBankai()
	}

	// 3. 处理虚化模式切换
	if intent.StatePressed == state.Vizard && !i.VizardMode && i.VizardTimer == 0 {
		i.activateVizard()
	}
}

// handleStateTransition 重写：自定义状态转换
// 适合处理：添加自定义状态、修改状态优先级等
func (i *Ichigo) handleStateTransition() {
	body := i.Runtime.Body

	// 先调用父类的状态转换逻辑
	// 这样可以保留基础状态转换，只添加额外逻辑
	i.BaseCharacter.handleStateTransition()

	// 添加自定义状态检查
	switch {
	case i.BankaiMode:
		i.Runtime.State = state.Bankai
	case i.VizardMode:
		i.Runtime.State = state.Vizard
	}
}

// handleMovementLock 重写：修改移动锁定逻辑
// 适合处理：不同状态下的移动限制、特殊状态下的移动自由等
func (i *Ichigo) handleMovementLock() {
	body := i.Runtime.Body

	// 在卍解模式下，可以自由移动
	if i.BankaiMode {
		return
	}

	// 在虚化模式下，移动速度提升但仍然有锁定
	if i.VizardMode {
		lockMoveX := i.Runtime.State == state.JustLanded
		if lockMoveX {
			body.X -= body.VX
			body.VX = 0
		}
		return
	}

	// 其他情况使用父类的逻辑
	i.BaseCharacter.handleMovementLock()
}

// handleFacing 重写：修改朝向逻辑
// 适合处理：攻击时强制朝向、特殊状态下的朝向锁定等
func (i *Ichigo) handleFacing() {
	body := i.Runtime.Body

	// 在虚化模式下，始终朝向对手（这里简化为朝向移动方向）
	if i.VizardMode {
		if body.VX > 0 {
			i.Runtime.Facing = 1
		} else if body.VX < 0 {
			i.Runtime.Facing = -1
		}
		return
	}

	// 其他情况使用父类的逻辑
	i.BaseCharacter.handleFacing()
}

// handlePostUpdate 重写：更新后的处理
// 适合处理：更新飞行物、处理冷却、模式时间管理等
func (i *Ichigo) handlePostUpdate() {
	// 1. 更新飞行物系统
	i.ProjectileExtension.Update()

	// 2. 更新技能冷却
	if i.GetsugaCooldown > 0 {
		i.GetsugaCooldown--
	}

	// 3. 更新卍解模式时间
	if i.BankaiMode {
		i.BankaiTimer++
		if i.BankaiTimer >= i.BankaiMaxTimer {
			i.deactivateBankai()
		}
	}

	// 4. 更新虚化模式时间
	if i.VizardMode {
		i.VizardTimer++
		if i.VizardTimer >= i.VizardMaxTimer {
			i.deactivateVizard()
		}
	}

	// 调用父类的 postUpdate（如果有）
	i.BaseCharacter.handlePostUpdate()
}

// launchGetsuga 发射月牙天冲（飞行物）
func (i *Ichigo) launchGetsuga() {
	speed := 12.0
	if i.BankaiMode {
		speed = 18.0 // 卍解模式下速度更快
	}
	if i.VizardMode {
		speed = 15.0 // 虚化模式下速度提升
	}

	i.ProjectileExtension.LaunchProjectile(
		i.Runtime.Body,
		i.Runtime.Facing,
		speed,
	)
}

// activateBankai 激活卍解模式
func (i *Ichigo) activateBankai() {
	i.BankaiMode = true
	i.BankaiTimer = 0
	i.Runtime.State = state.Bankai

	// 卍解模式下提升属性
	i.Data.MoveSpeed = 6
	i.Data.JumpPower = 15
}

// deactivateBankai 取消卍解模式
func (i *Ichigo) deactivateBankai() {
	i.BankaiMode = false
	i.BankaiTimer = 0

	// 恢复属性
	i.Data.MoveSpeed = 4
	i.Data.JumpPower = 12
}

// activateVizard 激活虚化模式
func (i *Ichigo) activateVizard() {
	i.VizardMode = true
	i.VizardTimer = 0
	i.Runtime.State = state.Vizard
}

// deactivateVizard 取消虚化模式
func (i *Ichigo) deactivateVizard() {
	i.VizardMode = false
	i.VizardTimer = 0
}

// buildAnimations 构建动画配置
func buildAnimations() animatable.Animations {
	return animatable.Animations{
		ByState: map[state.State]*animatable.Animation{
			state.Idle: {
				FPS:       12,
				Loop:      true,
				FramesKeys: []string{"idle_0001.png", "idle_0002.png", "idle_0003.png"},
			},
			state.Run: {
				FPS:       8,
				Loop:      true,
				FramesKeys: []string{"run_0001.png", "run_0002.png", "run_0003.png"},
			},
			state.Jump: {
				FPS:       10,
				Loop:      true,
				FramesKeys: []string{"jump_0001.png", "jump_0002.png"},
			},
			state.JumpStart: {
				FPS:       10,
				Loop:      false,
				FramesKeys: []string{"jump_start_0001.png"},
			},
			state.JustLanded: {
				FPS:       15,
				Loop:      false,
				FramesKeys: []string{"land_0001.png"},
			},
			state.Dash: {
				FPS:       12,
				Loop:      false,
				FramesKeys: []string{"dash_0001.png", "dash_0002.png"},
			},
			state.Bankai: {
				FPS:       12,
				Loop:      true,
				FramesKeys: []string{"bankai_0001.png", "bankai_0002.png"},
			},
			state.Vizard: {
				FPS:       15,
				Loop:      true,
				FramesKeys: []string{"vizard_0001.png", "vizard_0002.png"},
			},
		},
	}
}