package main

import "github.com/hajimehoshi/ebiten/v2"

type AnimationManager struct {
	Anims   map[AnimType]*Animation
	Current AnimType
}

// AnimType 定义要切换的动画类型
type AnimType int

const (
	AnimIdle AnimType = iota
	AnimRun
	AnimJump
	AnimAttack
)

// Animation 用于管理动画帧
type Animation struct {
	Frames []*ebiten.Image
	Index  int // 当前的帧的索引
	Timer  int
	Speed  int  //
	Loop   bool //是否循环播放
}

func (a *Animation) Reset() {
	a.Index = 0
	a.Timer = 0
}

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

func (a *Animation) CurrentFrame() *ebiten.Image {
	if len(a.Frames) == 0 {
		return nil
	}
	return a.Frames[a.Index]
}

func (a *Animation) IsFinished() bool {
	return !a.Loop && a.Index == len(a.Frames)-1
}

func (am *AnimationManager) Add(t AnimType, anim *Animation) {
	am.Anims[t] = anim
}

// Play 方法用于播放指定类型的动画
// 参数 t 表示要播放的动画类型
func (am *AnimationManager) Play(t AnimType) {
	// 如果当前正在播放的动画类型与请求的动画类型相同，则直接返回
	if am.Current == t {
		return
	}

	// 更新当前动画类型为请求的动画类型
	am.Current = t
	// 检查请求的动画类型是否存在于动画管理器中
	if anim, ok := am.Anims[t]; ok {
		// 如果存在，则重置动画到初始状态
		anim.Reset()
	}
}

func (am *AnimationManager) Update() {
	if anim, ok := am.Anims[am.Current]; ok {
		anim.Update()
	}
}

func (am *AnimationManager) CurrentFrame() *ebiten.Image {
	if anim, ok := am.Anims[am.Current]; ok {
		return anim.CurrentFrame()
	}
	return nil
}

func (am *AnimationManager) IsFinished() bool {
	if anim, ok := am.Anims[am.Current]; ok {
		return anim.IsFinished()
	}
	return true
}

func NewAnimationManager() *AnimationManager {
	return &AnimationManager{
		Anims: make(map[AnimType]*Animation),
	}
}
