package state

type State int

const (
	Idle       State = iota // 待机
	Run                     // 跑步
	JumpStart               // 刚开始跳
	Jump                    // 空中
	JustLanded              // 刚落地
	Dash                    // 冲刺

	//占位:倒地状态
	Fall
)

var StateNameList = [...]string{"idle", "run", "jump_start", "jump", "just_landed", "dash"}

// String 获取状态对应的字符串
func (s State) String() string {
	return StateNameList[s]
}
