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
	// 整个动作持续的逻辑帧数
	TotalFrames int64

	// —— 攻击有效帧（命中帧）——
	HitStart int64
	HitEnd   int64

	// —— 取消窗口（能接下一个动作）——
	CancelStart int64
	CancelEnd   int64
	NextOnJ     Action

	// —— 阻尼（按阶段）——
	DampingStartup  float64 // 前摇
	DampingActive   float64 // 有效帧
	DampingRecovery float64 // 后摇
}

var actionTable = map[Action]ActionDef{
	ActionJ1: {
		TotalFrames: 24,

		HitStart: 6,
		HitEnd:   8,

		CancelStart: 10,
		CancelEnd:   14,
		NextOnJ:     ActionJ2,

		DampingStartup:  0.2,  // 前摇基本站住
		DampingActive:   0.5,  // 出拳微微前冲
		DampingRecovery: 0.85, // 后摇滑一下
	},

	ActionJ2: {
		TotalFrames: 26,

		HitStart: 7,
		HitEnd:   9,

		CancelStart: 11,
		CancelEnd:   16,
		NextOnJ:     ActionJ3,

		DampingStartup:  0.25,
		DampingActive:   0.6,
		DampingRecovery: 0.8,
	},

	ActionJ3: {
		TotalFrames: 30,

		HitStart: 8,
		HitEnd:   12,

		CancelStart: -1, // 不能再接 J
		CancelEnd:   -1,

		DampingStartup:  0.3,
		DampingActive:   0.7,  // 最后一击冲得更明显
		DampingRecovery: 0.75, // 后摇重
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
