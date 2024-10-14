package nbody

import (
	"NBodySim/internal/vectormath"
	"fmt"
)

type Body struct {
	Id       int64
	Position *vectormath.Vector3d
	Velocity *vectormath.Vector3d
	Mass     float64
}

var nextBodyId int64 = 0

func getNextBodyId() int64 {
	nextBodyId++
	return nextBodyId
}

func NewBody(position *vectormath.Vector3d, velocity *vectormath.Vector3d, mass float64) *Body {
	return &Body{
		Id:       getNextBodyId(),
		Position: position,
		Velocity: velocity,
		Mass:     mass,
	}
}

func (b *Body) Copy() *Body {
	return &Body{
		b.Id,
		b.Position.Copy(),
		b.Velocity.Copy(),
		b.Mass,
	}
}

func (b *Body) Update(other *Body) {
	b.Id = other.Id
	b.Position = other.Position.Copy()
	b.Velocity = other.Velocity.Copy()
	b.Mass = other.Mass
}

func (b *Body) ToString() string {
	return fmt.Sprintf("Position: %v, Velocity: %v, Mass: %f",
		b.Position.ToSlice(), b.Velocity.ToSlice(), b.Mass)
}
