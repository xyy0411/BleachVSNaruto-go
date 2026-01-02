package rukia

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/xyy0411/ebiten_paractice/assets"
	"github.com/xyy0411/ebiten_paractice/common/state"
	"github.com/xyy0411/ebiten_paractice/models"
	"github.com/xyy0411/ebiten_paractice/render/animation"
)

func buildAnimations() animation.Set {
	return animation.Set{
		ByState: map[state.State]*models.ActionAnimation{
			state.StateIdle: {
				Frames: loadIdleFrames(),
				FPS:    6,
				Loop:   true,
			},
			state.StateRun: {
				Frames: loadRunFrames(),
				FPS:    10,
				Loop:   true,
			},
			state.StateJump: {
				Frames: loadJumpFrames(),
				FPS:    8,
				Loop:   false,
			},
		},
	}
}

func loadIdleFrames() []*ebiten.Image {
	url := "assets/characters/rukia/idle/"
	return []*ebiten.Image{
		assets.LoadImage(url + "0.png"),
		assets.LoadImage(url + "1.png"),
		assets.LoadImage(url + "2.png"),
		assets.LoadImage(url + "3.png"),
	}
}

func loadRunFrames() []*ebiten.Image {
	url := "assets/characters/rukia/run/"
	return []*ebiten.Image{
		assets.LoadImage(url + "0.png"),
		assets.LoadImage(url + "1.png"),
		assets.LoadImage(url + "2.png"),
		assets.LoadImage(url + "3.png"),
		assets.LoadImage(url + "4.png"),
		assets.LoadImage(url + "5.png"),
		assets.LoadImage(url + "6.png"),
	}
}

func loadJumpFrames() []*ebiten.Image {
	url := "assets/characters/rukia/jump/"
	return []*ebiten.Image{
		assets.LoadImage(url + "0.png"),
		assets.LoadImage(url + "1.png"),
		assets.LoadImage(url + "2.png"),
	}
}
