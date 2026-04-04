package animatable

import (
	"github.com/xyy0411/bleachVSnaruto/common/state"
)

// Set 动画集合，存储不同动作对应的动画
type Set struct {
	// ByState 状态到动画的映射表，用于根据角色状态获取对应的动画
	ByState map[state.State]*ActionAnimation
}
