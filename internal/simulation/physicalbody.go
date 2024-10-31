package simulation

import (
	"NBodySim/internal/mathutils/vector"
	"NBodySim/internal/nbody"
	"NBodySim/internal/object"
	"NBodySim/internal/transform"
)

type PhysicalBody struct {
	obj      object.Object
	velocity vector.Vector3d
	mass     float64
}

var ph nbody.Body = (*PhysicalBody)(nil)

func NewPhysicalBody(obj object.Object, velocity vector.Vector3d, mass float64) *PhysicalBody {
	return &PhysicalBody{obj: obj, velocity: velocity, mass: mass}
}

func (b *PhysicalBody) GetId() int64 {
	return b.obj.GetId()
}

func (b *PhysicalBody) Clone() nbody.Body {
	return NewPhysicalBody(b.obj.Clone(), b.velocity, b.mass)
}

func (b *PhysicalBody) GetPosition() vector.Vector3d {
	return b.obj.GetCenter()
}

func (b *PhysicalBody) GetVelocity() vector.Vector3d {
	return b.velocity
}

func (b *PhysicalBody) GetMass() float64 {
	return b.mass
}

func (b *PhysicalBody) SetVelocity(velocity vector.Vector3d) {
	b.velocity = velocity
}
func (b *PhysicalBody) SetPosition(position vector.Vector3d) {
	currentPos := b.GetPosition()
	delta := vector.SubtractVectors(&position, &currentPos)
	move := transform.NewMoveAction(delta)
	b.obj.Transform(move)
}

func (b *PhysicalBody) SetMass(mass float64) {
	b.mass = mass
}
