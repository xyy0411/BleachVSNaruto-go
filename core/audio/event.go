package audio

import "path/filepath"

// Event 表示可触发的音效事件
type Event string

const (
	EventRunStep    Event = "run"
	EventJump       Event = "jump"
	EventJumpStart  Event = "jump_start"
	EventJustLanded Event = "just_landed"
	EventDash       Event = "dash"
)

// DefaultSFX 返回默认公共音效事件映射表
func DefaultSFX() map[Event][]string {
	return map[Event][]string{
		EventJumpStart:  {defaultPublicAudioPath(EventJumpStart)},
		EventJustLanded: {defaultPublicAudioPath(EventJustLanded)},
		EventDash:       {defaultPublicAudioPath(EventDash)},
		EventRunStep: {
			"assets/public_audio/run/0.wav",
			"assets/public_audio/run/1.wav",
			"assets/public_audio/run/2.wav",
		},
		EventJump: {defaultPublicAudioPath(EventJump)},
	}
}

func defaultPublicAudioPath(event Event) string {
	return filepath.Join("assets", "public_audio", string(event), "0.wav")
}
