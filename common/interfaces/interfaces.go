package interfaces

import (
	"github.com/xyy0411/ebiten_paractice/models"
)

type Character interface {
	// GetID 获取角色ID
	GetID() string

	// GetName 获取角色名称
	GetName() string

	// Update 更新角色状态
	Update() error

	// ApplyInput 统一处理输入
	ApplyInput(input engine.Input) error

	// ApplyDamage 受到攻击事件（纯结果）
	ApplyDamage(dmg models.DamageEvent)

	// Snapshot 只读状态快照（用于渲染 / 联机 / Debug）
	Snapshot() models.CharacterSnapshot
}
