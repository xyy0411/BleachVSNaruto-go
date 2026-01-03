package animation

import "github.com/xyy0411/bleachVSnaruto/core/animatable"

type System struct{}

func (s *System) Update(chars []animatable.Animatable, delta float64) {
	for _, c := range chars {
		c.AnimationPlayer().Update(delta)
	}
}
