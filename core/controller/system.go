package controller

import (
	"github.com/xyy0411/bleachVSnaruto/common/state"
	"github.com/xyy0411/bleachVSnaruto/core/input"
	"github.com/xyy0411/bleachVSnaruto/models"
)

type System struct {
	Input *input.System
	// 本帧意图（供后续系统读取）
	Current models.Intent

	// 用于检测 JustPressed
	prevAttack   bool
	prevJump     bool
	prevDash     bool
	prevOnGround bool
	prevVy       float64
	prevState    state.State
}

func (s *System) Name() string {
	return "controller"
}

// Update 字段s.Current在此方法里已经代表为上一帧发生的事件了
func (s *System) Update() {
	// 获取本帧发生的事件
	in := s.Input.Current
	var intent models.Intent

	switch {
	case in.Left:
		intent.MoveX = -1
	case in.Right:
		intent.MoveX = 1
	}

	switch {
	case in.Up:
		intent.MoveY = 1
	case in.Down:
		intent.MoveY = -1
	}

	intent.AttackPressed = in.Attack && !s.prevAttack
	intent.JumpPressed = in.Jump && !s.prevJump
	intent.DashPressed = in.Dash && !s.prevDash

	// 刚落地
	if !s.prevOnGround && s.prevVy == 0 {
		intent.StatePressed = state.JustLanded
	}

	// 刚开始跳
	if s.prevOnGround && s.prevVy < 0 {
		intent.StatePressed = state.JumpStart
	}

	if intent.JumpPressed && s.prevState != state.Jump {
		intent.StatePressed = state.Jump
	}

	if intent.DashPressed && s.prevState != state.Dash {
		intent.StatePressed = state.Dash
	}

	intent.DashHeld = in.Dash

	s.Current = intent

	s.prevAttack = in.Attack
	s.prevJump = in.Jump
	s.prevDash = in.Dash
}
