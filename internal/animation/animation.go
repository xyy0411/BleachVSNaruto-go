package animation

import "github.com/hajimehoshi/ebiten/v2"

// Animation 用于管理单个动画序列
type Animation struct {
	Frames []*ebiten.Image // 动画帧集合
	Index  int             // 当前帧索引
	Timer  int             // 计时器累计值
	Speed  int             // 每帧切换所需的计数
	Loop   bool            // 是否循环播放
}

// Reset 将动画重置到初始帧
func (a *Animation) Reset() {
	a.Index = 0
	a.Timer = 0
}

// Update 推动动画计时并切换到下一帧
func (a *Animation) Update() {
	if len(a.Frames) == 0 {
		return
	}

	a.Timer++
	if a.Timer >= a.Speed {
		a.Timer = 0
		a.Index++

		if a.Index >= len(a.Frames) {
			if a.Loop {
				a.Index = 0
			} else {
				a.Index = len(a.Frames) - 1
			}
		}
	}
}

// CurrentFrame 返回当前帧
func (a *Animation) CurrentFrame() *ebiten.Image {
	if len(a.Frames) == 0 {
		return nil
	}
	return a.Frames[a.Index]
}

// IsFinished 判断动画是否播放到末尾（仅对非循环动画有效）
func (a *Animation) IsFinished() bool {
	return !a.Loop && a.Index == len(a.Frames)-1
}
