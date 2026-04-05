package narutoS

import (
	"github.com/xyy0411/bleachVSnaruto/characters"
	"github.com/xyy0411/bleachVSnaruto/common/state"
	"github.com/xyy0411/bleachVSnaruto/core/animatable"
	"github.com/xyy0411/bleachVSnaruto/core/audio"
	"github.com/xyy0411/bleachVSnaruto/core/charactor"
	"github.com/xyy0411/bleachVSnaruto/models"
)

// RoleID 是 narutoS 角色在注册表中的唯一标识
const RoleID = "narutoS"

var (
	actionFPS = map[string]int{
		"idle":        12,
		"run":         8,
		"jump":        10,
		"jump_start":  10,
		"just_landed": 15,
		"dash":        12,
	}
	actionLoop = map[string]bool{
		"idle": true,
		"run":  true,
		"jump": true,
	}
)

// NarutoS 表示使用 atlas 动画配置驱动的鸣人角色
type NarutoS struct {
	id   string
	name string

	Runtime *charactor.Runtime
	Data    *charactor.Data
}

func init() {
	characters.AddChar(RoleID, New)
}

// New 创建并返回一个 narutoS 角色实例
func New() charactor.Character {
	body := &models.PhysicsBody{
		X:            100,
		Y:            500,
		OnGround:     true,
		DashDuration: 0.3,
		MaxJumps:     2,
	}

	animations := buildAnimations()
	data := &charactor.Data{
		MoveSpeed:  4,
		JumpPower:  12,
		Animations: animations,
		Audio:      charactor.DefaultAudioConfig(),
	}

	rt := &charactor.Runtime{
		BodyRect:      &charactor.Rect{Width: 40, Height: 50, Action: true},
		Body:          body,
		Facing:        1,
		PrevOnGround:  body.OnGround,
		PrevVY:        body.VY,
		PrevJumpsUsed: body.JumpsUsed,
		AnimPlayer:    animatable.Player{},
	}

	return &NarutoS{
		id:      RoleID,
		name:    "漩涡鸣人",
		Runtime: rt,
		Data:    data,
	}
}

// GetData 返回角色静态配置数据
func (n *NarutoS) GetData() *charactor.Data {
	return n.Data
}

// GetRuntime 返回角色运行时状态
func (n *NarutoS) GetRuntime() *charactor.Runtime {
	return n.Runtime
}

// GetID 返回角色唯一标识
func (n *NarutoS) GetID() string {
	return n.id
}

// GetName 返回角色显示名称
func (n *NarutoS) GetName() string {
	return n.name
}

// Update 根据角色状态推进动画和事件。
func (n *NarutoS) Update() {
	body := n.Runtime.Body
	events := charactor.Events{}
	if !n.Runtime.PrevOnGround && body.OnGround && n.Runtime.PrevVY > 0 {
		events.JustLanded = true
	}
	if body.JumpsUsed > n.Runtime.PrevJumpsUsed && body.VY < 0 {
		events.JumpStart = true
	}
	n.Runtime.Events = events

	jumpStartAnim := n.Data.Animations.ByState[state.JumpStart]
	justLandedAnim := n.Data.Animations.ByState[state.JustLanded]

	jumpStartLocked := jumpStartAnim != nil &&
		!jumpStartAnim.Loop &&
		n.Runtime.AnimPlayer.Current == jumpStartAnim &&
		n.Runtime.AnimPlayer.Frame < int64(len(jumpStartAnim.FramesKeys)-1)
	justLandedLocked := justLandedAnim != nil &&
		!justLandedAnim.Loop &&
		n.Runtime.AnimPlayer.Current == justLandedAnim &&
		n.Runtime.AnimPlayer.Frame < int64(len(justLandedAnim.FramesKeys)-1)

	switch {
	case justLandedLocked:
		n.Runtime.State = state.JustLanded
	case jumpStartLocked:
		n.Runtime.State = state.JumpStart
	case events.JustLanded:
		n.Runtime.State = state.JustLanded
		n.Runtime.AudioEvents = append(n.Runtime.AudioEvents, audio.EventJustLanded)
	case events.JumpStart:
		n.Runtime.State = state.JumpStart
		n.Runtime.AudioEvents = append(n.Runtime.AudioEvents, audio.EventJumpStart)
	case body.Dashing:
		n.Runtime.State = state.Dash
	case body.OnGround:
		if body.VX != 0 {
			n.Runtime.State = state.Run
		} else {
			n.Runtime.State = state.Idle
		}
	default:
		n.Runtime.State = state.Jump
	}

	if !n.Runtime.PrevDashed && body.Dashing {
		n.Runtime.AudioEvents = append(n.Runtime.AudioEvents, audio.EventDash)
	}

	// 处理跑步音效
	// 这里暂时拿不到delta,反正tps是固定的那就直接硬算得了1/6秒循环播放一次
	if body.OnGround && body.VX != 0 && n.Runtime.State == state.Run {
		n.Runtime.RunStepTimer++
		if n.Runtime.RunStepTimer >= 10 {
			n.Runtime.AudioEvents = append(n.Runtime.AudioEvents, audio.EventRunStep)
			n.Runtime.RunStepTimer = 0
		}
	} else {
		n.Runtime.RunStepTimer = 0
	}

	lockMoveX := n.Runtime.State == state.JustLanded || n.Runtime.State == state.JumpStart
	if lockMoveX {
		body.X -= body.VX
		body.VX = 0
	}

	if body.OnGround && !lockMoveX && !body.Dashing {
		if body.VX > 0 {
			n.Runtime.Facing = 1
		} else if body.VX < 0 {
			n.Runtime.Facing = -1
		}
	}

	n.Runtime.AnimPlayer.Play(n.Data.Animations.ByState[n.Runtime.State])
	n.Runtime.PrevOnGround = body.OnGround
	n.Runtime.PrevVY = body.VY
	n.Runtime.PrevJumpsUsed = body.JumpsUsed
	n.Runtime.PrevDashed = body.Dashing

}
