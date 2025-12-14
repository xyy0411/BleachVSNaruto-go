package animation

import "github.com/hajimehoshi/ebiten/v2"

// AnimType 定义角色可播放的动画类型
type AnimType int

const (
	AnimIdle AnimType = iota
	AnimRun
	AnimJump
	AnimAttack
)

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

// Manager 管理角色的多个动画序列
type Manager struct {
	Anims   map[AnimType]*Animation // 已注册的动画集合
	Current AnimType                // 当前播放的动画类型
}

// Add 注册新的动画
func (m *Manager) Add(t AnimType, anim *Animation) {
	m.Anims[t] = anim
}

// Play 切换到指定的动画类型
func (m *Manager) Play(t AnimType) {
	if m.Current == t {
		return
	}

	m.Current = t
	if anim, ok := m.Anims[t]; ok {
		anim.Reset()
	}
}

// Update 更新当前动画的播放进度
func (m *Manager) Update() {
	if anim, ok := m.Anims[m.Current]; ok {
		anim.Update()
	}
}

// CurrentFrame 返回当前正在播放的帧
func (m *Manager) CurrentFrame() *ebiten.Image {
	if anim, ok := m.Anims[m.Current]; ok {
		return anim.CurrentFrame()
	}
	return nil
}

// IsFinished 判断当前动画是否播放完毕
func (m *Manager) IsFinished() bool {
	if anim, ok := m.Anims[m.Current]; ok {
		return anim.IsFinished()
	}
	return true
}

// NewManager 创建动画管理器
func NewManager() *Manager {
	return &Manager{Anims: make(map[AnimType]*Animation)}
}
