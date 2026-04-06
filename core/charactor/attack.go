package charactor

import (
	"github.com/xyy0411/bleachVSnaruto/common/state"
	"github.com/xyy0411/bleachVSnaruto/core/animatable"
)

// ResolveAttackState 根据当前动画进度推进普通攻击状态机，并返回本帧是否应保持攻击状态。
func ResolveAttackState(runtime *Runtime, animations animatable.Set) bool {
	if runtime == nil || runtime.Body == nil {
		return false
	}

	body := runtime.Body
	if !body.State.IsAttack() {
		if !body.ConsumeAttackRequest() {
			return false
		}

		if attackAnimationByState(animations, state.J1) == nil {
			return false
		}

		body.StartAttack(state.J1)
		return true
	}

	currentAnim := attackAnimationByState(animations, body.State)
	if currentAnim == nil || len(currentAnim.FramesKeys) == 0 {
		body.FinishAttack()
		return false
	}

	if runtime.AnimPlayer.Current != currentAnim {
		return true
	}

	lastFrame := int64(len(currentAnim.FramesKeys) - 1)
	if runtime.AnimPlayer.Frame < lastFrame {
		return true
	}

	if body.HasQueuedAttack() {
		nextState := state.NextAttackState(body.State)
		if attackAnimationByState(animations, nextState) != nil {
			body.AdvanceAttackState()
			return true
		}
	}

	body.FinishAttack()
	return false
}

func attackAnimationByState(animations animatable.Set, st state.State) *animatable.ActionAnimation {
	return animations.ByState[st]
}
