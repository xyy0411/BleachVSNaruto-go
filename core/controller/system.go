package controller

import (
	"github.com/xyy0411/bleachVSnaruto/core/charactor"
	"github.com/xyy0411/bleachVSnaruto/core/input"
)

// System 负责将输入帧转换为角色运行时可消费的行为意图
type System struct {
	Input *input.System
	// Runtime 指向当前控制器绑定的角色运行时数据
	Runtime *charactor.Runtime

	// prevAttack prevJump prevDash 用于检测 JustPressed 输入边沿
	prevAttack bool
	prevJump   bool
	prevDash   bool
}

// Name 返回系统名称
func (s *System) Name() string {
	return "controller"
}

// Update 根据输入状态更新当前帧的角色行为意图
func (s *System) Update() {
	if s.Input == nil || s.Runtime == nil {
		return
	}

	in := s.Input.Current
	intent := s.Runtime.Intent

	switch {
	case in.Left:
		intent.MoveX = -1
	case in.Right:
		intent.MoveX = 1
	default:
		intent.MoveX = 0
	}

	switch {
	case in.Up:
		intent.MoveY = 1
	case in.Down:
		intent.MoveY = -1
	default:
		intent.MoveY = 0
	}

	intent.AttackPressed = in.Attack && !s.prevAttack
	intent.JumpPressed = in.Jump && !s.prevJump
	intent.DashPressed = in.Dash && !s.prevDash
	intent.DashHeld = in.Dash

	s.Runtime.Intent = intent
	s.prevAttack = in.Attack
	s.prevJump = in.Jump
	s.prevDash = in.Dash
}
