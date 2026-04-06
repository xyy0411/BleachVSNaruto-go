package models

import "github.com/xyy0411/bleachVSnaruto/common/state"

// PhysicsBody 表示角色在战斗场景中的基础物理状态与动作判定数据。
type PhysicsBody struct {
	X, Y   float64
	VX, VY float64

	OnGround bool

	// Dashing 标记当前帧是否在冲刺，用于状态与动画切换
	Dashing bool

	// 冲刺参数与计时
	DashDuration  float64 // 冲刺持续时间（秒）
	DashTimer     float64 // 剩余冲刺时间
	DashDirection int     // 当前冲刺方向（-1 / 1）

	// 跳跃参数
	MaxJumps  int // 最大可连续跳跃次数（含地面起跳）
	JumpsUsed int // 已使用的跳跃次数

	State state.State //可以用来记录普攻段数等信息
	// 攻击参数
	AttackRequested bool // 是否请求开始普通攻击
	AttackQueued    bool // 当前攻击段结束后是否衔接下一段攻击

}

// RegisterAttackInput 记录一次普通攻击按键输入。
func (b *PhysicsBody) RegisterAttackInput() {
	if b.State.IsAttack() {
		b.AttackQueued = true
		return
	}

	b.AttackRequested = true
}

// ConsumeAttackRequest 消耗一次开始普通攻击的请求。
func (b *PhysicsBody) ConsumeAttackRequest() bool {
	if b == nil || !b.AttackRequested {
		return false
	}

	b.AttackRequested = false
	return true
}

// StartAttack 启动指定的普通攻击状态。
func (b *PhysicsBody) StartAttack(attackState state.State) {
	if b == nil {
		return
	}

	if !attackState.IsAttack() {
		attackState = state.J1
	}

	b.State = attackState
	b.AttackRequested = false
	b.AttackQueued = false
	b.Dashing = false
	b.DashTimer = 0
	b.VX = 0
}

// AdvanceAttackState 切换到下一段普通攻击状态。
func (b *PhysicsBody) AdvanceAttackState() state.State {
	nextState := state.NextAttackState(b.State)
	b.StartAttack(nextState)
	return nextState
}

// HasQueuedAttack 返回当前攻击段结束后是否需要衔接下一段攻击。
func (b *PhysicsBody) HasQueuedAttack() bool {
	return b.AttackQueued
}

// FinishAttack 结束当前普通攻击状态并清空缓冲。
func (b *PhysicsBody) FinishAttack() {
	b.State = state.Idle
	b.AttackRequested = false
	b.AttackQueued = false
}
