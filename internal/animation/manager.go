package animation

import "github.com/hajimehoshi/ebiten/v2"

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
