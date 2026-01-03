package main

import (
	"github.com/xyy0411/bleachVSnaruto/characters/rukia"
	"github.com/xyy0411/bleachVSnaruto/config"
	"github.com/xyy0411/bleachVSnaruto/core/controller"
	"github.com/xyy0411/bleachVSnaruto/core/input"
	"github.com/xyy0411/bleachVSnaruto/core/physics"
	"github.com/xyy0411/bleachVSnaruto/core/world"
	"github.com/xyy0411/bleachVSnaruto/engine"
	"github.com/xyy0411/bleachVSnaruto/game"
	"github.com/xyy0411/bleachVSnaruto/global"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	config.InitLog()

	e := engine.New()

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
		JumpSpeed:  10,
		DashSpeed:  10,
	}

	e.InputSystem = inputSys
	e.PhysicsSystem = physicsSys
	e.RegisterSystem(controllerSys)
	e.RegisterSystem(physicsSys)

	player := rukia.New()
	rt := player.GetRuntime()
	rt.Body.Y = w.GroundY
	rt.Body.OnGround = true

	physicsSys.Bodies = append(physicsSys.Bodies, rt.Body)

	e.RegisterActor(player)

	g := game.Game{Engine: e}
	ebiten.SetWindowSize(800, 600)
	ebiten.SetWindowTitle("死神VS火影 demo")
	ebiten.SetTPS(60)
	global.Logger.Infoln("开始")
	if err := ebiten.RunGame(&g); err != nil {
		log.Fatal(err)
	}
}
