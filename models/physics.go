package models

type PhysicsBody struct {
	X, Y   float64
	VX, VY float64

	OnGround bool

	// Dashing 标记当前帧是否在冲刺，用于状态与动画切换
	Dashing bool
}
