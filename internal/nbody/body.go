package nbody

import (
	"NBodySim/internal/vectormath"
	"fmt"
)

type Body struct {
	Position *vectormath.Vector3d
	Velocity *vectormath.Vector3d
	Mass     float64
}

func NewBody(position *vectormath.Vector3d, velocity *vectormath.Vector3d, mass float64) *Body {
	return &Body{
		Position: position,
		Velocity: velocity,
		Mass:     mass,
	}
}

func (b *Body) Copy() *Body {
	return NewBody(
		b.Position.Copy(),
		b.Velocity.Copy(),
		b.Mass,
	)
}

func (b *Body) ToString() string {
	return fmt.Sprintf("Position: %v, Velocity: %v, Mass: %f",
		b.Position.ToSlice(), b.Velocity.ToSlice(), b.Mass)
}
