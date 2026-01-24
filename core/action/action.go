package action

// Action 描述角色的动作定义
type Action int

const (
	ActionBlink Action = iota
	ActionAttackJ
	ActionAttackADJ
	ActionAttackSJ
	ActionAttackWJ
	ActionAttackU
	ActionAttackADU
	ActionAttackWU
)
