package models

type CharacterStats struct {
	MaxHP   int
	Speed   float64
	JumpPow float64
	Weight  float64
}

type InputFrame struct {
	Frame int64

	// 方向
	Left  bool
	Right bool
	Up    bool
	Down  bool

	// 动作
	Attack bool
	Dash   bool
	Jump   bool
}
