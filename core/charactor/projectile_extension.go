package charactor

import (
	"github.com/xyy0411/bleachVSnaruto/models"
)

// Projectile 飞行物数据结构
type Projectile struct {
	X, Y       float64
	VX, VY     float64
	Facing     int
	Active     bool
	Damage     int
	Lifetime   int
	MaxLifetime int
}

// ProjectileSystem 飞行物系统
type ProjectileSystem struct {
	Projectiles []*Projectile
}

// NewProjectileSystem 创建新的飞行物系统
func NewProjectileSystem() *ProjectileSystem {
	return &ProjectileSystem{
		Projectiles: make([]*Projectile, 0),
	}
}

// LaunchProjectile 发射飞行物
func (ps *ProjectileSystem) LaunchProjectile(x, y float64, facing int, speed float64) {
	proj := &Projectile{
		X:           x,
		Y:           y,
		VX:          speed * float64(facing),
		VY:          0,
		Facing:      facing,
		Active:      true,
		Damage:      10,
		Lifetime:    0,
		MaxLifetime: 60, // 1秒后消失（假设TPS=60）
	}
	ps.Projectiles = append(ps.Projectiles, proj)
}

// UpdateProjectiles 更新所有飞行物
func (ps *ProjectileSystem) UpdateProjectiles() {
	for _, proj := range ps.Projectiles {
		if proj.Active {
			proj.X += proj.VX
			proj.Y += proj.VY
			proj.Lifetime++

			// 检查是否超出生命周期
			if proj.Lifetime >= proj.MaxLifetime {
				proj.Active = false
			}
		}
	}

	// 清理已失效的飞行物
	ps.cleanupProjectiles()
}

// cleanupProjectiles 清理失效的飞行物
func (ps *ProjectileSystem) cleanupProjectiles() {
	active := make([]*Projectile, 0)
	for _, proj := range ps.Projectiles {
		if proj.Active {
			active = append(active, proj)
		}
	}
	ps.Projectiles = active
}

// GetActiveProjectiles 获取所有活跃的飞行物
func (ps *ProjectileSystem) GetActiveProjectiles() []*Projectile {
	active := make([]*Projectile, 0)
	for _, proj := range ps.Projectiles {
		if proj.Active {
			active = append(active, proj)
		}
	}
	return active
}

// ProjectileCharacterExtension 飞行物角色扩展
// 这是一个示例，展示如何为角色添加飞行物能力
type ProjectileCharacterExtension struct {
	ProjectileSystem *ProjectileSystem
	LaunchCooldown   int
	MaxCooldown      int
}

// NewProjectileCharacterExtension 创建飞行物角色扩展
func NewProjectileCharacterExtension(maxCooldown int) *ProjectileCharacterExtension {
	return &ProjectileCharacterExtension{
		ProjectileSystem: NewProjectileSystem(),
		LaunchCooldown:   0,
		MaxCooldown:      maxCooldown,
	}
}

// CanLaunchProjectile 检查是否可以发射飞行物
func (ext *ProjectileCharacterExtension) CanLaunchProjectile() bool {
	return ext.LaunchCooldown <= 0
}

// LaunchProjectile 发射飞行物（如果有冷却）
func (ext *ProjectileCharacterExtension) LaunchProjectile(body *models.PhysicsBody, facing int, speed float64) {
	if ext.CanLaunchProjectile() {
		ext.ProjectileSystem.LaunchProjectile(body.X, body.Y, facing, speed)
		ext.LaunchCooldown = ext.MaxCooldown
	}
}

// Update 更新飞行物系统和冷却
func (ext *ProjectileCharacterExtension) Update() {
	ext.ProjectileSystem.UpdateProjectiles()

	// 更新冷却
	if ext.LaunchCooldown > 0 {
		ext.LaunchCooldown--
	}
}

// ExampleProjectileCharacter 示例：有飞行物能力的角色
// 这个示例展示了如何重写 BaseCharacter 的方法来实现特殊能力
type ExampleProjectileCharacter struct {
	*BaseCharacter
	ProjectileExtension *ProjectileCharacterExtension
}

// NewExampleProjectileCharacter 创建示例角色
func NewExampleProjectileCharacter(id, name string, runtime *Runtime, data *Data) *ExampleProjectileCharacter {
	return &ExampleProjectileCharacter{
		BaseCharacter:      NewBaseCharacter(id, name, runtime, data),
		ProjectileExtension: NewProjectileCharacterExtension(30), // 0.5秒冷却
	}
}

// handleSpecialActions 重写：处理飞行物发射
// 这是钩子方法，在状态转换之前执行
func (c *ExampleProjectileCharacter) handleSpecialActions() {
	// 检查是否有攻击意图
	if c.Runtime.Intent.AttackPressed && c.ProjectileExtension.CanLaunchProjectile() {
		// 发射飞行物
		c.ProjectileExtension.LaunchProjectile(
			c.Runtime.Body,
			c.Runtime.Facing,
			10.0, // 飞行物速度
		)

		// 可以在这里添加特殊状态转换
		// 例如：c.Runtime.State = state.Attack
	}
}

// handlePostUpdate 重写：更新飞行物位置
// 这是钩子方法，在所有更新完成后执行
func (c *ExampleProjectileCharacter) handlePostUpdate() {
	// 更新飞行物系统
	c.ProjectileExtension.Update()

	// 可以在这里添加其他逻辑，比如：
	// - 检查飞行物碰撞
	// - 处理特殊技能冷却
	// - 更新其他自定义系统
}

// GetProjectiles 获取角色的所有飞行物（供渲染系统使用）
func (c *ExampleProjectileCharacter) GetProjectiles() []*Projectile {
	return c.ProjectileExtension.ProjectileSystem.GetActiveProjectiles()
}