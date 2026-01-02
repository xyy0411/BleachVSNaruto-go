package models

type PhysicsBody struct {
	X, Y   float64
	VX, VY float64

	OnGround bool
}
