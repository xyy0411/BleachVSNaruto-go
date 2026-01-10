package effect

import "github.com/xyy0411/bleachVSnaruto/core/action"

// Set 特效集合，用于管理不同时机触发的特效
type Set struct {
	// OnActionStart 动作开始时触发的特效映射表
	// key: 动作类型, value: 该动作开始时触发的特效列表
	OnActionStart map[action.Action][]Def
	// OnHit 命中时触发的特效映射表
	// key: 动作类型, value: 该动作命中时触发的特效列表
	OnHit map[action.Action][]Def
}

// Def 特效定义，描述单个特效的基本属性
type Def struct {
	// Name 特效的名称标识
	Name string
	// AtFrame 特效触发的帧数（相对于动作开始帧）
	AtFrame int64 // ActionFrame
}
