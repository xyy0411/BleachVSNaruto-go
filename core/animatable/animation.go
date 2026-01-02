package animatable

type Animatable interface {
	AnimationPlayer() AnimationPlayer
}

type AnimationPlayer interface {
	Update(delta float64)
}
