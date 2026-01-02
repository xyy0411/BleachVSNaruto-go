package state

type State int

const (
	StateIdle State = iota
	StateRun
	StateJump
)
