package event

import "github.com/xyy0411/bleachVSnaruto/core/audio"

// Type 表示事件类型
type Type int

const (
	// None 无事件
	None Type = iota
	// Audio 音频事件
	Audio
	// StateChange 状态变更事件
	StateChange
	// Movement 移动事件
	Movement
)

// Event 表示一个事件
type Event struct {
	Type Type
	// CharacterID 触发事件的角色ID
	CharacterID string
	// AudioEvent 音频事件数据（当 Type == Audio 时使用）
	AudioEvent audio.Event
	// Data 通用事件数据
	Data interface{}
}

// Handler 事件处理函数
type Handler func(Event)