package models

type CharacterStats struct {
	MaxHP   int
	Speed   float64
	JumpPow float64
	Weight  float64
}

// InputFrame 是本帧发生的事件
type InputFrame struct {
	Frame int64

	// 方向
	Left  bool
	Right bool
	Up    bool
	Down  bool

	// 动作
	Attack     bool // J
	Dash       bool // L
	Jump       bool // K
	LangAtk    bool // U
	Outbreak   bool // I
	Assistance bool // O
}
