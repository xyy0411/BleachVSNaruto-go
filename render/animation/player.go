package animation

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/xyy0411/bleachVSnaruto/assets"
	"github.com/xyy0411/bleachVSnaruto/models"
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
	if p.Current == nil || p.Current.FPS <= 0 {
		return
	}
	if delta <= 0 || math.IsNaN(delta) || math.IsInf(delta, 0) {
		return
	}

	frameTime := 1.0 / float64(p.Current.FPS)
	if frameTime <= 0 || math.IsNaN(frameTime) || math.IsInf(frameTime, 0) {
		return
	}

	p.Timer += delta

	// 计算本帧需要推进的帧数，并做一个上限，防止极端值导致长循环
	steps := int(math.Floor(p.Timer / frameTime))
	const maxAdvance = 16 // 随意安全上限
	if steps > maxAdvance {
		steps = maxAdvance
		p.Timer = 0 // 异常情况清空累积时间
	} else {
		p.Timer -= float64(steps) * frameTime
	}

	for i := 0; i < steps; i++ {
		p.Frame++
		if p.Frame >= int64(len(p.Current.FramesKeys)) {
			if p.Current.Loop {
				p.Frame = 0
			} else {
				p.Frame = int64(len(p.Current.FramesKeys) - 1)
			}
		}
	}
}

// CurrentFrame 返回玩家当前动画帧的图像
func (p *Player) CurrentFrame() *ebiten.Image {
	if p.Current == nil || len(p.Current.FramesKeys) <= 0 {
		return nil
	}
	//return p.Current.FramesKeys[p.Frame]
	return assets.StdImagePool.GetImage(p.Current.FramesKeys[p.Frame])
}
