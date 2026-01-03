package animation

import (
	"github.com/xyy0411/bleachVSnaruto/common/state"
	"github.com/xyy0411/bleachVSnaruto/core/action"
	"github.com/xyy0411/bleachVSnaruto/models"
)

// Set 动画集合，存储不同动作对应的动画
type Set struct {
	// ByState 状态到动画的映射表，用于根据角色状态获取对应的动画
	ByState map[state.State]*models.ActionAnimation
	// ByAction 动作到动画的映射表，用于根据动作类型获取对应的动画
	ByAction map[action.Action]models.ActionAnimation
}
