package character

// CharState 角色状态枚举
type CharState int

const (
	StateIdle CharState = iota
	StateRun
	StateJump
	State
	StateAttack
)

// DirInput 表示方向键输入
type DirInput int

const (
	DirNone DirInput = iota
	DirUp
	DirDown
	DirForward
)
