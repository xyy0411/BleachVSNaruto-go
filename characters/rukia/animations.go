package rukia

import (
	"github.com/xyy0411/bleachVSnaruto/assets"
	"github.com/xyy0411/bleachVSnaruto/common/state"
	"github.com/xyy0411/bleachVSnaruto/models"
	"github.com/xyy0411/bleachVSnaruto/render/animation"
)

func buildAnimations() animation.Set {
	return animation.Set{
		ByState: map[state.State]*models.ActionAnimation{
			state.Idle: {
				FramesKeys: assets.LoadCharacterActionFrames(RoleID, "idle"),
				FPS:        6,
				Loop:       true,
			},
			state.Run: {
				FramesKeys: assets.LoadCharacterActionFrames(RoleID, "run"),
				FPS:        5,
				Loop:       true,
			},
			state.JumpStart: {
				FramesKeys: assets.LoadCharacterActionFrames(RoleID, "jump_star"),
				FPS:        10,
				Loop:       false,
			},
			state.Jump: {
				FramesKeys: assets.LoadCharacterActionFrames(RoleID, "jump"),
				FPS:        13,
				Loop:       true,
			},
			state.JustLanded: {
				FramesKeys: assets.LoadCharacterActionFrames(RoleID, "just_landed"),
				FPS:        5,
				Loop:       false,
			},
			state.Dash: {
				FramesKeys: assets.LoadCharacterActionFrames(RoleID, "dash"),
				FPS:        5,
				Loop:       false,
			},
		},
	}
}