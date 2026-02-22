package charactor

import "github.com/xyy0411/bleachVSnaruto/core/audio"

func DefaultAudioConfig() AudioConfig {
	return AudioConfig{
		SFX:    audio.DefaultSFX(),
		Volume: 1.0,
	}
}
