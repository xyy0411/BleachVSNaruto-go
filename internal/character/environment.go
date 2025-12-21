package character

type Environment interface {
	GroundY(x float64) float64
	CheckPlatform(x, y, vy float64) (onGround bool, newY float64)
}
