package character

type ActionPhase int

const (
	PhaseStartup  ActionPhase = iota // 前摇
	PhaseActive                      // 有效
	PhaseRecovery                    // 后摇
)

// Phase 方法用于根据给定的帧数判断当前动作所处的阶段
// 参数:
//
//	frame - 当前帧数，用于判断动作阶段
//
// 返回值:
//
//	ActionPhase - 根据帧数返回对应的动作阶段
func (a ActionDef) Phase(frame int64) ActionPhase {
	// 如果当前帧数小于命中开始帧，则处于动作启动阶段
	if frame < a.HitStart {
		return PhaseStartup
	}
	// 如果当前帧数在命中开始帧和结束帧之间（包含结束帧），则处于动作生效阶段
	if frame <= a.HitEnd {
		return PhaseActive
	}
	// 其他情况则处于动作恢复阶段
	return PhaseRecovery
}
