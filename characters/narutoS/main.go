package narutoS

import (
	"github.com/xyy0411/bleachVSnaruto/characters"
	"github.com/xyy0411/bleachVSnaruto/core/animatable"
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
	*charactor.BaseCharacter
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
		Body:          body,
		Facing:        1,
		PrevOnGround:  body.OnGround,
		PrevVY:        body.VY,
		PrevJumpsUsed: body.JumpsUsed,
		AnimPlayer:    animatable.Player{},
	}

	base := charactor.NewBaseCharacter(RoleID, "漩涡鸣人", rt, data)

	return &NarutoS{
		BaseCharacter: base,
	}
}
