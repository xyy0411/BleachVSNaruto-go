package models

// ActionAnimation 动作动画，包含一个动作的完整动画信息
type ActionAnimation struct {
	// Frames 动画帧序列，存储该动作的所有帧图像Key
	FramesKeys []string
	// FPS 动画播放的帧率，每秒播放的帧数
	FPS int
	// Loop 是否循环播放该动画
	Loop bool
}
