package animation

// AnimType 定义角色可播放的动画类型
type AnimType int

const (
	AnimIdle AnimType = iota
	AnimRun
	AnimJump
	AnimAttack
)
