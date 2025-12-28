package models

type Character struct {
	ID int

	// 运行时状态（所有角色共用）
	X, Y   float64
	Vx, Vy float64
	Facing int

	Action      Action
	ActionFrame int64

	Attack *AttackController

	// 指向角色定义
	Data *CharacterData
}

type CharacterData struct {
	Name string

	Actions map[Action]ActionDef
	Attacks map[Action]AttackDef

	Animations map[Action]AnimID

	Stats CharacterStats
}

type CharacterStats struct {
	MaxHP   int
	Speed   float64
	JumpPow float64
	Weight  float64
}

type DamageEvent struct {
	SourceID string
	Damage   int
	KnockX   float64
	KnockY   float64
	HitType  HitType
}

type CharacterSnapshot struct {
	X, Y        float64
	Facing      int
	Action      Action
	ActionFrame int64
	HP          int
}
