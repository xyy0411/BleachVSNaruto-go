package state

type State int

const (
	Idle       State = iota // 待机
	Run                     // 跑步
	JumpStart               // 刚开始跳
	Jump                    // 空中
	JustLanded              // 刚落地
	Dash                    // 冲刺
	NormalJ
	WJ
	SJ
	KJ
	NormalU
	WU
	SU
	KU
	NormalI
	WI
	SI
	KI
)

// String 获取状态对应的字符串
func String(s State) string {
	return [...]string{
		"idle",
		"run",
		"jump_start",
		"jump",
		"just_landed",
		"dash",
		"normal_j",
		"wj",
		"sj",
		"kj",
		"normal_u",
		"wu",
		"su",
		"ku",
		"normal_i",
		"wi",
		"si",
		"ki",
	}[s]
}

// ToString 将字符串转换为状态
func ToString(s string) State {
	stateMap := map[string]State{
		"idle":        Idle,
		"run":         Run,
		"jump_start":  JumpStart,
		"jump":        Jump,
		"just_landed": JustLanded,
		"dash":        Dash,
		"normal_j":    NormalJ,
		"wj":          WJ,
		"sj":          SJ,
		"kj":          KJ,
		"normal_u":    NormalU,
		"wu":          WU,
		"su":          SU,
		"ku":          KU,
		"normal_i":    NormalI,
		"wi":          WI,
		"si":          SI,
		"ki":          KI,
	}
	if state, ok := stateMap[s]; ok {
		return state
	}
	return Idle
}
