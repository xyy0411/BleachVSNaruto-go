package rukia

import (
	"github.com/xyy0411/bleachVSnaruto/common/state"
	"github.com/xyy0411/bleachVSnaruto/models"
	"github.com/xyy0411/bleachVSnaruto/render/animation"
)

func buildAnimations() animation.Set {
	return animation.Set{
		ByState: map[state.State]*models.ActionAnimation{
			state.Idle: {
				FramesKeys: loadIdleFrames(),
				FPS:    6,
				Loop:   true,
			},
			state.Run: {
			FramesKeys: loadRunFrames(),
				FPS:    5,
				Loop:   true,
			},
			state.JumpStart: {
			FramesKeys: loadJumpStartFrames(),
				FPS:    10,
				Loop:   false,
			},
			state.Jump: {
			FramesKeys: loadJumpFrames(),
				FPS:    13,
				Loop:   true,
			},
			state.JustLanded: {
			FramesKeys: loadJustLandedFrames(),
				FPS:    5,
				Loop:   false,
			},
			state.Dash: {
			FramesKeys: loadDashFrames(),
				FPS:    5,
				Loop:   false,
			},
		},
	}
}

func loadIdleFrames() []string {
	url := "assets/characters/rukia/animation/idle/"
	return []string{
		url + "0.png", url + "1.png", url + "2.png", url + "3.png",
	}
}

func loadRunFrames() []string {
	url := "assets/characters/rukia/animation/run/"
	return []string{
		url + "0.png", url + "1.png", url + "2.png", url + "3.png",
		url + "4.png", url + "5.png", url + "6.png",
	}
}

func loadJumpFrames() []string {
	url := "assets/characters/rukia/animation/jump/"
	return []string{
		url + "0.png",
	}
}

func loadJustLandedFrames() []string {
	url := "assets/characters/rukia/animation/just_landed/"
	return []string{
		url + "0.png", url + "1.png", url + "2.png", url + "3.png",
	}
}

func loadJumpStartFrames() []string {
	url := "assets/characters/rukia/animation/jump_start/"
	return []string{
		url + "0.png",
	}
}

func loadDashFrames() []string {
	url := "assets/characters/rukia/animation/dash/"
	return []string{
		url + "0.png", url + "1.png", url + "2.png",
	}
}
