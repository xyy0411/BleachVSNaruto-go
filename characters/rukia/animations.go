package rukia

import (
	"github.com/xyy0411/bleachVSnaruto/assets"
	"github.com/xyy0411/bleachVSnaruto/common/state"
	"github.com/xyy0411/bleachVSnaruto/core/animatable"
)

func buildAnimations() animatable.Set {
	return animatable.Set{
		ByState: map[state.State]*animatable.ActionAnimation{
			state.Idle: {
				FramesKeys: assets.LoadCharacterActionFrames(RoleID, state.Idle.String()),
				FPS:        6,
				Loop:       true,
			},
			state.Run: {
				FramesKeys: assets.LoadCharacterActionFrames(RoleID, state.Run.String()),
				FPS:        5,
				Loop:       true,
			},
			state.JumpStart: {
				FramesKeys: assets.LoadCharacterActionFrames(RoleID, state.JumpStart.String()),
				FPS:        10,
				Loop:       false,
			},
			state.Jump: {
				FramesKeys: assets.LoadCharacterActionFrames(RoleID, state.Jump.String()),
				FPS:        13,
				Loop:       true,
			},
			state.JustLanded: {
				FramesKeys: assets.LoadCharacterActionFrames(RoleID, state.JustLanded.String()),
				FPS:        10,
				Loop:       false,
			},
			state.Dash: {
				FramesKeys: assets.LoadCharacterActionFrames(RoleID, state.Dash.String()),
				FPS:        5,
				Loop:       false,
			},
		},
	}
}
