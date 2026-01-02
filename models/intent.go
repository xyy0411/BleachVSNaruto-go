package models

type Intent struct {
	// 方向意图（连续态）
	MoveX int // -1, 0, 1
	MoveY int // -1, 0, 1

	// 动作意图（瞬时/持续）
	AttackPressed bool
	JumpPressed   bool
	DashHeld      bool
}
