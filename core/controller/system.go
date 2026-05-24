package controller

import (
	"github.com/xyy0411/bleachVSnaruto/core/action"
	"github.com/xyy0411/bleachVSnaruto/core/charactor"
	"github.com/xyy0411/bleachVSnaruto/core/input"
	"github.com/xyy0411/bleachVSnaruto/models"
)

// System 负责将输入帧转换为角色运行时可消费的行为意图
type System struct {
	Input *input.System
	// Runtime 指向当前控制器绑定的角色运行时数据
	Runtime *charactor.Runtime

	// prevAttack prevJump prevDash 用于检测 JustPressed 输入边沿
	prevAttack   bool
	prevJump     bool
	prevDash     bool
	prevLangAtk  bool
	prevOutbreak bool
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
	intent.ComboCommand = ""
	intent.ComboStage = 0

	s.resolveComboWithNewSystem(&intent, in)

	s.Runtime.Intent = intent
	s.prevAttack = in.Attack
	s.prevJump = in.Jump
	s.prevDash = in.Dash
	s.prevLangAtk = in.LangAtk
	s.prevOutbreak = in.Outbreak
}

// resolveComboWithNewSystem 使用新连招系统解析输入
func (s *System) resolveComboWithNewSystem(intent *models.Intent, in models.InputFrame) {
	if s.Runtime.ComboSystem == nil || s.Runtime.ComboSystem.Table == nil {
		return
	}

	// 获取当前动作上下文
	ctx := s.Runtime.ComboSystem.Context

	// 查询连招表
	trans := s.Runtime.ComboSystem.Table.Resolve(ctx, s.Input.Buffer, in.Frame)

	if trans == nil {
		return
	}

	intent.ComboCommand = trans.ToAction
	intent.ComboStage = trans.ToStage
	// 更新连招系统状态
	s.Runtime.ComboSystem.Context = action.Context{
		Action: trans.ToAction,
		Stage:  trans.ToStage,
		Frame:  0,
	}

	// 没有找到连招，检查是否有新的输入触发
	s.checkNewInput(intent, in)
}

// checkNewInput 检查是否有新的输入触发
func (s *System) checkNewInput(intent *models.Intent, in models.InputFrame) {
	// 检测按键输入
	attackPressed := in.Attack && !s.prevAttack
	langAtkPressed := in.LangAtk && !s.prevLangAtk
	outbreakPressed := in.Outbreak && !s.prevOutbreak

	// 列举所有二键起手连招
	switch {
	case in.Up && attackPressed:
		intent.ComboCommand = "wj"
		intent.ComboStage = 1
	case in.Down && attackPressed:
		intent.ComboCommand = "sj"
		intent.ComboStage = 1
	case in.Jump && attackPressed:
		intent.ComboCommand = "kj"
		intent.ComboStage = 1
	case attackPressed:
		intent.ComboCommand = "J1"
		intent.ComboStage = 1
	case in.Up && langAtkPressed:
		intent.ComboCommand = "wu"
		intent.ComboStage = 1
	case in.Down && langAtkPressed:
		intent.ComboCommand = "su"
		intent.ComboStage = 1
	case in.Jump && langAtkPressed:
		intent.ComboCommand = "ku"
		intent.ComboStage = 1
	case langAtkPressed:
		intent.ComboCommand = "u"
		intent.ComboStage = 1
	case in.Up && outbreakPressed:
		intent.ComboCommand = "wi"
		intent.ComboStage = 1
	case in.Down && outbreakPressed:
		intent.ComboCommand = "si"
		intent.ComboStage = 1
	case in.Jump && outbreakPressed:
		intent.ComboCommand = "ki"
		intent.ComboStage = 1
	case outbreakPressed:
		intent.ComboCommand = "i"
		intent.ComboStage = 1
	}

	// 如果触发了新动作，重置连招系统上下文
	if intent.ComboCommand != "" {
		s.Runtime.ComboSystem.Context = action.Context{
			Action: intent.ComboCommand,
			Stage:  intent.ComboStage,
			Frame:  0,
		}
	}
}
