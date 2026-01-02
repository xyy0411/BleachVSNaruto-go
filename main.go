package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/xyy0411/ebiten_paractice/characters/rukia"
	"github.com/xyy0411/ebiten_paractice/config"
	"github.com/xyy0411/ebiten_paractice/core/controller"
	"github.com/xyy0411/ebiten_paractice/core/input"
	"github.com/xyy0411/ebiten_paractice/core/physics"
	"github.com/xyy0411/ebiten_paractice/core/world"
	"github.com/xyy0411/ebiten_paractice/engine"
	"github.com/xyy0411/ebiten_paractice/game"
	"github.com/xyy0411/ebiten_paractice/models"
	"log"
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
		Gravity:    0.8,
		MoveSpeed:  4,
		JumpSpeed:  12,
	}

	playerBody := &models.PhysicsBody{
		X:        100,
		Y:        w.GroundY,
		OnGround: true,
	}

	physicsSys.Bodies = append(physicsSys.Bodies, playerBody)

	e.PhysicsSystem = physicsSys
	e.RegisterSystem(controllerSys)
	e.RegisterSystem(physicsSys)

	e.RegisterActor(rukia.New())

	g := game.Game{Engine: e}
	ebiten.SetWindowSize(800, 600)
	ebiten.SetWindowTitle("死神VS火影 demo")
	if err := ebiten.RunGame(&g); err != nil {
		log.Fatal(err)
	}
}
