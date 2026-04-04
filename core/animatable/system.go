package animatable

type System struct{}

func (s *System) Update(chars []Animatable, delta float64) {
	for _, c := range chars {
		c.AnimationPlayer().Update(delta)
	}
}
