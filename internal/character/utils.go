package character

import "github.com/xyy0411/ebiten_paractice/internal/animation"

func actionToAnim(a Action) animation.AnimType {
	switch a {
	case ActionJ1:
		return animation.AnimJ1
	case ActionJ2:
		return animation.AnimJ2
	case ActionJ3:
		return animation.AnimJ3
	default:
		return animation.AnimIdle
	}
}
