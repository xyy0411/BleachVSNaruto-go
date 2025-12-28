package models

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
