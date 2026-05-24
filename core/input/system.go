package input

import (
	"github.com/xyy0411/bleachVSnaruto/core/time"
	"github.com/xyy0411/bleachVSnaruto/models"
)

const DefaultBufferMaxAge int64 = 12

// System 输入系统，用于管理游戏输入相关的状态和数据
type System struct {
	// Time 用于获取时间相关的信息
	Time *time.Time
	// Source 输入源，用于获取输入数据
	Source Source

	// Current 当前输入帧数据，存储当前帧的输入状态
	Current models.InputFrame
	// Previous 上一帧的输入状态，用于检测边沿
	Previous models.InputFrame
	// Buffer 输入缓冲区
	Buffer *Buffer
}

// NewSystem ...
func NewSystem(time *time.Time, source Source, bufferMaxAge int64) *System {
	if bufferMaxAge <= 0 {
		bufferMaxAge = DefaultBufferMaxAge
	}
	return &System{
		Time:   time,
		Source: source,
		Buffer: NewInputBuffer(bufferMaxAge),
	}
}

// Collect 收集数据帧
func (s *System) Collect() {
	s.Previous = s.Current
	frame := s.Source.Read()
	frame.Frame = s.Time.GlobalFrame
	s.Current = frame

	// 检测按键边沿并推入缓冲区
	s.pushEdgeEvents(frame.Frame)

	// 清理过期事件
	s.Buffer.Clean(frame.Frame)
}

// pushEdgeEvents 检测按键边沿并推入缓冲区
func (s *System) pushEdgeEvents(currentFrame int64) {
	// Attack (j)
	if s.Current.Attack && !s.Previous.Attack {
		s.Buffer.Push("j", currentFrame, true)
	} else if !s.Current.Attack && s.Previous.Attack {
		s.Buffer.Push("j", currentFrame, false)
	}

	// LangAtk (u)
	if s.Current.LangAtk && !s.Previous.LangAtk {
		s.Buffer.Push("u", currentFrame, true)
	} else if !s.Current.LangAtk && s.Previous.LangAtk {
		s.Buffer.Push("u", currentFrame, false)
	}

	// Outbreak (i)
	if s.Current.Outbreak && !s.Previous.Outbreak {
		s.Buffer.Push("i", currentFrame, true)
	} else if !s.Current.Outbreak && s.Previous.Outbreak {
		s.Buffer.Push("i", currentFrame, false)
	}

	// Dash
	if s.Current.Dash && !s.Previous.Dash {
		s.Buffer.Push("dash", currentFrame, true)
	} else if !s.Current.Dash && s.Previous.Dash {
		s.Buffer.Push("dash", currentFrame, false)
	}

	// Jump (k)
	if s.Current.Jump && !s.Previous.Jump {
		s.Buffer.Push("k", currentFrame, true)
	} else if !s.Current.Jump && s.Previous.Jump {
		s.Buffer.Push("k", currentFrame, false)
	}

	// Up (w)
	if s.Current.Up && !s.Previous.Up {
		s.Buffer.Push("w", currentFrame, true)
	} else if !s.Current.Up && s.Previous.Up {
		s.Buffer.Push("w", currentFrame, false)
	}

	// Down (s)
	if s.Current.Down && !s.Previous.Down {
		s.Buffer.Push("s", currentFrame, true)
	} else if !s.Current.Down && s.Previous.Down {
		s.Buffer.Push("s", currentFrame, false)
	}
}
