package engine

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/xyy0411/bleachVSnaruto/core/animatable"
	"github.com/xyy0411/bleachVSnaruto/core/charactor"
	"github.com/xyy0411/bleachVSnaruto/core/input"
	"github.com/xyy0411/bleachVSnaruto/core/physics"
	"github.com/xyy0411/bleachVSnaruto/core/time"
	"github.com/xyy0411/bleachVSnaruto/global"
	"github.com/xyy0411/bleachVSnaruto/render/animation"
)

type System interface {
	Update()
	Name() string
}

type Engine struct {
	system          []System
	actors          []charactor.Character
	Time            *time.Time
	InputSystem     []*input.System
	PhysicsSystem   *physics.System
	AnimationSystem *animation.System
}

func New(TPS int) *Engine {
	ebiten.SetTPS(TPS)
	return &Engine{
		Time:            new(time.Time).UpdataTPS(float64(TPS)),
		InputSystem:     []*input.System{},
		AnimationSystem: &animation.System{},
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
	for _, v := range e.InputSystem {
		v.Collect()
	}

	for _, system := range e.system {
		system.Update()
	}

	var animatables []animatable.Animatable
	for _, actor := range e.actors {
		actor.Update()
		animatables = append(animatables, actor.GetRuntime())
	}

	e.AnimationSystem.Update(animatables, e.Time.Delta)

	if e.PhysicsSystem == nil || e.PhysicsSystem.World == nil {
		return
	}

	var targets []float64
	for _, actor := range e.actors {
		rt := actor.GetRuntime()
		if rt == nil || rt.Body == nil {
			continue
		}

		frame := rt.AnimPlayer.CurrentFrame()
		frameWidth := 0.0
		if frame != nil {
			frameWidth = float64(frame.Bounds().Dx())
		}

		clampedX := e.PhysicsSystem.World.ClampBodyRectX(rt.Body.X, frameWidth)
		if clampedX != rt.Body.X {
			rt.Body.X = clampedX
			rt.Body.VX = 0
			rt.Body.Dashing = false
			rt.Body.DashTimer = 0
		}

		// 摄像头按角色中心点跟随
		// 这样左右朝向变化或动画宽度变化时更稳定
		targets = append(targets, rt.Body.X+frameWidth/2)
	}
	e.PhysicsSystem.World.FollowTargetsX(targets...)
}

func (e *Engine) Draw(screen *ebiten.Image) {
	camera := e.PhysicsSystem.World.Camera
	cameraX := 0.0
	cameraZoom := 1.0
	if camera != nil {
		cameraX = camera.X
		cameraZoom = camera.Scale()
	}

	e.PhysicsSystem.World.GroundPainter.Draw(screen, cameraX, cameraZoom, e.PhysicsSystem.World.GroundY)
	for _, actor := range e.actors {
		rt := actor.GetRuntime()
		frame := rt.AnimPlayer.CurrentFrame()
		if frame == nil {
			continue
		}

		op := &ebiten.DrawImageOptions{}

		if rt.Facing == -1 {
			op.GeoM.Scale(-1, 1)
			op.GeoM.Translate(float64(frame.Bounds().Dx()), 0)
		}

		op.GeoM.Scale(cameraZoom, cameraZoom)
		drawX, drawY := camera.WorldToScreen(rt.Body.X, rt.Body.Y-float64(frame.Bounds().Dy()))
		op.GeoM.Translate(drawX, drawY)
		screen.DrawImage(frame, op)
	}
}
