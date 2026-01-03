package state

type State int

const (
	Idle State = iota
	Run
	Jump
	Dash
)
