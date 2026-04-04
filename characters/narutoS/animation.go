package narutoS

import (
	"github.com/xyy0411/bleachVSnaruto/assets"
	"github.com/xyy0411/bleachVSnaruto/common/state"
	"github.com/xyy0411/bleachVSnaruto/core/animatable"
	"github.com/xyy0411/bleachVSnaruto/global"
)

func buildAnimations() animatable.Set {
	cfg, err := assets.LoadCharacterAnimationConfig(RoleID)
	if err != nil {
		global.Logger.WithError(err).Error("加载动画配置文件失败")
		return animatable.Set{ByState: map[state.State]*animatable.ActionAnimation{}}
	}

	set, err := assets.BuildAnimationSetFromAtlas(RoleID, cfg, actionFPS, actionLoop)
	if err != nil {
		global.Logger.WithError(err).Errorf("构建鸣人动画帧失败")
		return animatable.Set{ByState: map[state.State]*animatable.ActionAnimation{}}
	}

	if set.ByState == nil {
		set.ByState = map[state.State]*animatable.ActionAnimation{}
	}
	if set.ByState[state.JumpStart] == nil {
		set.ByState[state.JumpStart] = set.ByState[state.Jump]
	}

	return set
}
