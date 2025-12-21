package character

// Action 描述角色的动作定义
type Action int

const (
	ActionIdle Action = iota
	ActionRun
	ActionJump
	ActionBlink
	ActionJ1
	ActionJ2
	ActionJ3
	ActionAttackJ
	ActionAttackADJ
	ActionAttackSJ
	ActionAttackWJ
	ActionAttackU
	ActionAttackADU
	ActionAttackWU
)

// ActionDef 控制动作帧与取消窗口
type ActionDef struct {
	TotalFrames   int64          // 总帧数
	CancelWindows []CancelWindow // 允许取消的区间
	CancelStart   int64
	CancelEnd     int64
	NextOnJ       Action
}

var actionTable = map[Action]ActionDef{
	ActionJ1: {
		TotalFrames: 24,
		CancelStart: 8,
		CancelEnd:   14,
		NextOnJ:     ActionJ2,
	},
	ActionJ2: {
		TotalFrames: 26,
		CancelStart: 10,
		CancelEnd:   16,
		NextOnJ:     ActionJ3,
	},
	ActionJ3: {
		TotalFrames: 30,
		CancelStart: -1, // 不可取消
		CancelEnd:   -1,
	},
}

// CancelWindow 定义动作取消的可行区间
type CancelWindow struct {
	Start int      // 取消窗口起始帧
	End   int      // 取消窗口结束帧
	To    []Action // 可取消到的动作列表
}

// AttackKey 普通攻击按键
type AttackKey int

const (
	AttackJ AttackKey = iota
	AttackU
	AttackI
)

const (
	AttackDuration = 20 // 攻击持续帧数
)

// SkillKey 用于匹配招式表
type SkillKey struct {
	Dir DirInput  // 方向输入
	Key AttackKey // 具体按键
}

// Skill 预留的技能定义结构
type Skill struct {
	Frame int // 触发帧数
}
