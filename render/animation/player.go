package animation

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/xyy0411/ebiten_paractice/models"
)

// Player 动画播放器，用于控制和管理动画的播放状态
type Player struct {
	// Current 当前正在播放的动作动画
	Current *models.ActionAnimation

	// Frame 当前已经进行的动画帧数
	Frame int64
	// Timer 时间计时器用于判断是否切换动画
	// 单位(s)
	Timer float64
}

func (p *Player) Play(anim *models.ActionAnimation) {
	if p.Current == anim {
		return
	}

	p.Current = anim
	p.Frame = 0
	p.Timer = 0
}

// Update 方法用于更新玩家的动画状态
// 参数 delta 表示自上一帧以来经过的时间（秒）
func (p *Player) Update(delta float64) {
	// 检查当前动画是否存在或FPS是否有效
	if p.Current == nil || p.Current.FPS <= 0 {
		return
	}

	// 计算每帧应该持续的时间（秒）
	frameTime := float64(1.0 / p.Current.FPS)
	p.Timer += delta

	// 当累计时间超过单帧时间时，更新动画帧
	for p.Timer >= frameTime {
		// 减去已消耗的帧时间
		p.Timer -= frameTime
		// 切换到下一帧
		p.Frame++

		// 检查是否到达动画末尾
		if p.Frame >= int64(len(p.Current.Frames)) {
			// 根据Loop属性决定是否循环播放
			if p.Current.Loop {
				p.Frame = 0
			} else {
				p.Frame = int64(len(p.Current.Frames) - 1)
			}
		}
	}
}

// CurrentFrame 返回玩家当前动画帧的图像
func (p *Player) CurrentFrame() *ebiten.Image {
	if p.Current == nil || len(p.Current.Frames) <= 0 {
		return nil
	}
	return p.Current.Frames[p.Frame]
}
