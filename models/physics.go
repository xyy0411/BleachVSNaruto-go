package models

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

}
