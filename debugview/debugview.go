package debugview

import (
	"fmt"
	"image/color"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	textv2 "github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/xyy0411/bleachVSnaruto/assets"
	"github.com/xyy0411/bleachVSnaruto/common/state"
	"github.com/xyy0411/bleachVSnaruto/engine"
)

// Panel 显示调试信息
type Panel struct {
	Visible bool

	face    *textv2.GoTextFace
	faceErr error
}

// Update ...
func (p *Panel) Update() {
	if inpututil.IsKeyJustPressed(ebiten.KeyF3) {
		p.Visible = !p.Visible
	}
}

// Draw ...
func (p *Panel) Draw(screen *ebiten.Image, e *engine.Engine) {
	if !p.Visible || e == nil || e.PhysicsSystem == nil || e.PhysicsSystem.World == nil {
		return
	}
	if err := p.ensureFace(); err != nil {
		return
	}

	var lines []string
	lines = append(lines, fmt.Sprintf("FPS: %.2f", ebiten.ActualFPS()))
	lines = append(lines, fmt.Sprintf("TPS: %.2f", ebiten.ActualTPS()))
	lines = append(lines, fmt.Sprintf("当前角色数: %d", e.ActorCount()))
	lines = append(lines, fmt.Sprintf("总推进帧数: %d", e.Time.GlobalFrame))

	camera := e.PhysicsSystem.World.Camera
	if camera != nil {
		lines = append(lines, fmt.Sprintf(
			"相机X: %.2f 相机缩放: %.2f 可视宽度: %.2f",
			camera.X,
			camera.Scale(),
			camera.VisibleWidth(),
		))
	}

	for i, actor := range e.Actors() {
		rt := actor.GetRuntime()
		if rt == nil || rt.Body == nil {
			continue
		}

		lines = append(lines, fmt.Sprintf(
			"角色%d %s 状态=%s 位置=(%.1f, %.1f) 速度=(%.1f, %.1f) 动画帧=%d",
			i+1,
			actor.GetName(),
			state.String(rt.State),
			rt.Body.X,
			rt.Body.Y,
			rt.Body.VX,
			rt.Body.VY,
			rt.AnimPlayer.Frame,
		))
	}

	op := &textv2.DrawOptions{}
	op.GeoM.Translate(8, 8)
	op.ColorScale.ScaleWithColor(color.White)
	op.LineSpacing = 20
	textv2.Draw(screen, strings.Join(lines, "\n"), p.face, op)
}

// ensureFace 确保调试面板字体已加载
func (p *Panel) ensureFace() error {
	if p.face != nil || p.faceErr != nil {
		return p.faceErr
	}

	face, err := assets.DebugTextFace(16)
	if err != nil {
		p.faceErr = err
		return err
	}

	p.face = face
	return nil
}
