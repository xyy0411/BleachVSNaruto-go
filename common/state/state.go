package state

type State int

const (
	Idle         State = iota // 待机
	Run                       // 跑步
	JumpStart                 // 刚开始跳
	Jump                      // 空中
	JustLanded                // 刚落地
	Dash                      // 冲刺
	Animation                 // animation
	Hit                       // hit
	IAfter                    // I_after
	IBefore                   // I_before
	IMiss                     // I_miss
	J1                        // J1
	J2                        // J2
	J3                        // J3
	J4                        // J4
	Fall                      // fall
	KI                        // KI
	Kick                      // kick
	KJ                        // KJ
	KU                        // KU
	KUDown                    // KU_down
	S                         // S
	SIBefore                  // SI_before
	SIMiss                    // SI_miss
	SIAfter                   // SI_after
	SJ                        // SJ
	SU                        // SU
	SUU                       // SUU
	U                         // U
	UAfter                    // U_after
	Unknown                   // Unknown
	WIBefore                  // WI_before
	WIMiss                    // WI_miss
	WIAfter                   // WI_after
	AnimationWin              // animation_win
	WJ                        // WJ
	WU                        // WU
	WUU                       // WUU

	NumOfState //State
)

var StateNameList = [...]string{
	"idle", "run", "jump_start", "jump", "just_landed", "dash",
	"animation", "hit", "I_after", "I_before", "I_miss",
	"J1", "J2", "J3", "J4", "fall",
	"KI", "kick", "KJ", "KU", "KU_down",
	"S", "SI_before", "SI_miss", "SI_after", "SJ", "SU", "SUU", "U", "U_after",
	"Unknown", "WI_before",
	"WI_miss", "WI_after", "animation_win", "WJ", "WU", "WUU",
}

// String 获取状态对应的字符串
func (s State) String() string {
	return StateNameList[s]
}
