package input

import (
	"github.com/xyy0411/ebiten_paractice/core/time"
	"github.com/xyy0411/ebiten_paractice/models"
)

// System 输入系统，用于管理游戏输入相关的状态和数据
type System struct {
	// Time 用于获取时间相关的信息
	Time *time.Time
	// Source 输入源，用于获取输入数据
	Source Source

	// Current 当前输入帧数据，存储当前帧的输入状态
	Current models.InputFrame
}

// Collect 收集数据帧
func (s *System) Collect() {
	frame := s.Source.Read()
	frame.Frame = s.Time.GlobalFrame
	s.Current = frame
}
