package animation

import "github.com/xyy0411/ebiten_paractice/core/animatable"

type System struct{}

func (s *System) Update(chars []animatable.Animatable, delta float64) {
	for _, c := range chars {
		c.AnimationPlayer().Update(delta)
	}
}
