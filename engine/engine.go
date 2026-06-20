package engine

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/xyy0411/bleachVSnaruto/core/animatable"
	"github.com/xyy0411/bleachVSnaruto/core/audio"
	"github.com/xyy0411/bleachVSnaruto/core/charactor"
	"github.com/xyy0411/bleachVSnaruto/core/event"
	"github.com/xyy0411/bleachVSnaruto/core/input"
	"github.com/xyy0411/bleachVSnaruto/core/physics"
	"github.com/xyy0411/bleachVSnaruto/core/time"
	"github.com/xyy0411/bleachVSnaruto/global"
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
	AnimationSystem *animatable.System
	AudioSystem     *audio.System
	EventBus        *event.Bus
}

// New 创建新的引擎实例
func New(TPS int) *Engine {
	ebiten.SetTPS(TPS)
	eventBus := event.NewBus()
	e := &Engine{
		Time:            new(time.Time).UpdataTPS(float64(TPS)),
		InputSystem:     []*input.System{},
		AnimationSystem: &animatable.System{},
		EventBus:        eventBus,
	}

	// 订阅音频事件
	e.subscribeAudioEvents()

	return e
}

// subscribeAudioEvents 订阅音频事件并处理
func (e *Engine) subscribeAudioEvents() {
	e.EventBus.Subscribe(event.Audio, func(ev event.Event) {
		if e.AudioSystem == nil {
			return
		}

		// 查找触发事件的角色
		var actor charactor.Character
		for _, a := range e.actors {
			if a.GetID() == ev.CharacterID {
				actor = a
				break
			}
		}

		if actor == nil {
			return
		}

		data := actor.GetData()
		runtime := actor.GetRuntime()

		paths := data.Audio.SFX[ev.AudioEvent]
		if len(paths) == 0 {
			return
		}

		if runtime.LastAudioVariant == nil {
			runtime.LastAudioVariant = make(map[audio.Event]int)
		}

		index := runtime.LastAudioVariant[ev.AudioEvent]
		path := paths[index%len(paths)]
		runtime.LastAudioVariant[ev.AudioEvent] = (index + 1) % len(paths)

		e.AudioSystem.Play(path, data.Audio.Volume)
	})
}

// ActorCount 返回当前已注册的角色数量
func (e *Engine) ActorCount() int {
	return len(e.actors)
}

// Actors 返回当前已注册的角色列表
func (e *Engine) Actors() []charactor.Character {
	return e.actors
}

func (e *Engine) RegisterSystem(s System) {
	global.Logger.Infoln("[Engine] Register System:", s.Name())
	e.system = append(e.system, s)
}

func (e *Engine) RegisterActor(a charactor.Character) {
	global.Logger.Infoln("[Engine] Register Actor:", a.GetName())
	a.SetEventBus(e.EventBus)
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

		leftX, frameWidth, _ := currentFrameLayout(rt)
		clampedLeftX := e.PhysicsSystem.World.ClampBodyRectX(leftX, frameWidth)
		clampedBodyX := rt.Body.X + (clampedLeftX - leftX)
		if clampedBodyX != rt.Body.X {
			rt.Body.X = clampedBodyX
			rt.Body.VX = 0
			rt.Body.Dashing = false
			rt.Body.DashTimer = 0
		}

		targets = append(targets, rt.Body.X)
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
		drawX, drawY := camera.WorldToScreen(frameDrawPosition(rt))
		op.GeoM.Translate(drawX, drawY)
		screen.DrawImage(frame, op)
	}
}

// currentFrameLayout 计算当前动画帧的布局信息
func currentFrameLayout(rt *charactor.Runtime) (leftX float64, width float64, centerX float64) {
	frame := rt.AnimPlayer.CurrentFrame()
	if frame == nil {
		return rt.Body.X, 0, rt.Body.X
	}

	meta := rt.AnimPlayer.CurrentFrameMeta()
	frameWidth := float64(frame.Bounds().Dx())
	originX := 0.0
	if meta != nil {
		originX = meta.Origin.X
	}

	// 根据角色朝向计算帧图像的左边界位置
	if rt.Facing == -1 {
		leftX = rt.Body.X - (frameWidth - originX)
	} else {
		leftX = rt.Body.X - originX
	}

	return leftX, frameWidth, leftX + frameWidth/2
}

func frameDrawPosition(rt *charactor.Runtime) (float64, float64) {
	frame := rt.AnimPlayer.CurrentFrame()
	if frame == nil {
		return rt.Body.X, rt.Body.Y
	}

	meta := rt.AnimPlayer.CurrentFrameMeta()
	originX := 0.0
	originY := float64(frame.Bounds().Dy())
	if meta != nil {
		originX = meta.Origin.X
		originY = meta.Origin.Y
	}

	drawX := rt.Body.X - originX
	if rt.Facing == -1 {
		drawX = rt.Body.X - (float64(frame.Bounds().Dx()) - originX)
	}

	return drawX, rt.Body.Y - originY
}
