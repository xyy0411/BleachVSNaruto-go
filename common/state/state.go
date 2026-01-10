package state

type State int

const (
	Idle State = iota
	Run
	JumpStart
	Jump // 空中
	JustLanded
	Dash
)
