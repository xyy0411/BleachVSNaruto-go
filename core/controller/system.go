package controller

import (
	"github.com/xyy0411/ebiten_paractice/core/input"
	"github.com/xyy0411/ebiten_paractice/models"
)

type System struct {
	Input *input.System

	// 本帧意图（供后续系统读取）
	Current models.Intent

	// 用于检测 JustPressed
	prevAttack bool
	prevJump   bool
}

func (s *System) Name() string {
	return "controller"
}

func (s *System) Update() {
	in := s.Input.Current

	var intent models.Intent

	if in.Left {
		intent.MoveX = 1
	} else if in.Right {
		intent.MoveX = -1
	}

	if in.Up {
		intent.MoveY = 1
	} else if in.Down {
		intent.MoveY = -1
	}
	intent.AttackPressed = in.Attack && !s.prevAttack

	intent.JumpPressed = in.Jump && !s.prevJump

	intent.DashHeld = in.Dash

	s.Current = intent

	s.prevJump = in.Jump
	s.prevAttack = in.Attack
}
