package assets

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/xyy0411/bleachVSnaruto/core/animatable"
)

func init() {
	animatable.SetFrameImageResolver(func(key string) *ebiten.Image {
		return StdImagePool.GetImage(key)
	})
}
