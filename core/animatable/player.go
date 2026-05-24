package animatable

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

var frameImageResolver = func(string) *ebiten.Image {
	return nil
}

// SetFrameImageResolver 设置动画帧 key 到实际图像的解析函数
func SetFrameImageResolver(resolver func(string) *ebiten.Image) {
	if resolver == nil {
		frameImageResolver = func(string) *ebiten.Image {
			return nil
		}
		return
	}
	frameImageResolver = resolver
}

// Player 负责维护当前动作、帧索引与播放计时
type Player struct {
	// Current 表示当前正在播放的动作动画
	Current *ActionAnimation
	// Frame 表示当前帧索引
	Frame int64
	// Timer 表示累计未消费的时间，单位为秒
	Timer float64
}

// Play 切换到新的动作动画
func (p *Player) Play(anim *ActionAnimation) {
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

	steps := int(math.Floor(p.Timer / frameTime))
	const maxAdvance = 16
	if steps > maxAdvance {
		steps = maxAdvance
		p.Timer = 0
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

// CurrentFrame 返回当前帧的图像
func (p *Player) CurrentFrame() *ebiten.Image {
	if p.Current == nil || len(p.Current.FramesKeys) == 0 {
		return nil
	}
	return frameImageResolver(p.Current.FramesKeys[p.Frame])
}

// CurrentFrameMeta 返回当前帧的元数据
func (p *Player) CurrentFrameMeta() *AnimationFrame {
	if p.Current == nil || len(p.Current.Frames) == 0 {
		return nil
	}

	index := int(p.Frame)
	if index < 0 || index >= len(p.Current.Frames) {
		return nil
	}

	return &p.Current.Frames[index]
}
