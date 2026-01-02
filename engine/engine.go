package engine

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/xyy0411/ebiten_paractice/core/animatable"
	"github.com/xyy0411/ebiten_paractice/core/charactor"
	"github.com/xyy0411/ebiten_paractice/core/input"
	"github.com/xyy0411/ebiten_paractice/core/physics"
	"github.com/xyy0411/ebiten_paractice/core/time"
	"github.com/xyy0411/ebiten_paractice/global"
	"github.com/xyy0411/ebiten_paractice/render/animation"
	"github.com/xyy0411/ebiten_paractice/render/draw"
)

type System interface {
	Update()
	Name() string
}

type Engine struct {
	system          []System
	actors          []charactor.Character
	Time            *time.Time
	InputSystem     *input.System
	PhysicsSystem   *physics.System
	AnimationSystem *animation.System
}

func New() *Engine {
	return &Engine{
		Time: &time.Time{
			TPS: 30,
		},
		InputSystem: &input.System{},
	}
}

func (e *Engine) RegisterSystem(s System) {
	global.Logger.Infoln("[Engine] Register System:", s.Name())
	e.system = append(e.system, s)
}

func (e *Engine) RegisterActor(a charactor.Character) {
	global.Logger.Infoln("[Engine] Register Actor:", a.GetName())
	e.actors = append(e.actors, a)
}

func (e *Engine) Update() {
	e.Time.Tick()
	e.InputSystem.Collect()

	for _, system := range e.system {
		system.Update()
	}

	var animatables []animatable.Animatable
	for _, actor := range e.actors {
		actor.Update()
		animatables = append(animatables, actor.GetRuntime())
	}

	e.AnimationSystem.Update(animatables, e.Time.Delta)
}

func (e *Engine) Draw(screen *ebiten.Image) {
	draw.Ground(screen, e.PhysicsSystem.World.GroundY)
	for _, actor := range e.actors {
		rt := actor.GetRuntime()
		frame := rt.AnimPlayer.CurrentFrame()
		if frame == nil {
			continue
		}

		op := &ebiten.DrawImageOptions{}

		// 反转图像
		if rt.Facing == -1 {
			op.GeoM.Scale(-1, 1)
			op.GeoM.Translate(float64(frame.Bounds().Dx()), 0)
		}

		op.GeoM.Translate(rt.Body.X, rt.Body.Y)
		screen.DrawImage(frame, op)
	}
}
