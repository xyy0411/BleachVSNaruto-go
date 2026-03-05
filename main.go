package main

import (
	"log"

	"github.com/xyy0411/bleachVSnaruto/characters"
	"github.com/xyy0411/bleachVSnaruto/config"
	coreaudio "github.com/xyy0411/bleachVSnaruto/core/audio"
	"github.com/xyy0411/bleachVSnaruto/core/controller"
	"github.com/xyy0411/bleachVSnaruto/core/input"
	"github.com/xyy0411/bleachVSnaruto/core/physics"
	"github.com/xyy0411/bleachVSnaruto/core/world"
	"github.com/xyy0411/bleachVSnaruto/engine"
	"github.com/xyy0411/bleachVSnaruto/game"
	"github.com/xyy0411/bleachVSnaruto/global"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"

	//初始化角色
	_ "github.com/xyy0411/bleachVSnaruto/characters/rukia"
)

func main() {
	config.InitLog()

	e := engine.New(60)
	audioCtx := audio.NewContext(11000)
	coreaudio.Init(audioCtx)

	inputSys := &input.System{
		Time:   e.Time,
		Source: &input.KeyboardSource{},
	}

	controllerSys := &controller.System{
		Input: inputSys,
	}

	w := world.World{
		GroundY: 500,
	}

	physicsSys := &physics.System{
		Controller: controllerSys,
		World:      &w,
		Time:       e.Time,
		Gravity:    0.8,
		MoveSpeed:  5,
		JumpSpeed:  17,
		DashSpeed:  8,
	}

	e.InputSystem = inputSys
	e.PhysicsSystem = physicsSys
	e.RegisterSystem(controllerSys)
	e.RegisterSystem(physicsSys)

	//以后逻辑修改为用户选择角色
	player := characters.SelectChar("rukia")()
	rt := player.GetRuntime()
	rt.Body.Y = w.GroundY
	rt.Body.OnGround = true

	physicsSys.Bodies = append(physicsSys.Bodies, rt.Body)

	e.RegisterActor(player)
	g := game.Game{Engine: e}
	ebiten.SetWindowSize(800, 600)
	ebiten.SetWindowTitle("死神VS火影 demo")
	global.Logger.Infoln("开始")
	if err := ebiten.RunGame(&g); err != nil {
		log.Fatal(err)
	}
}
