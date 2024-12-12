package simulation

import (
	"NBodySim/internal/mathutils/vector"
	"NBodySim/internal/nbody"
	"NBodySim/internal/object"
	"NBodySim/internal/transform"
	"math"
)

type PhysicalBody struct {
	obj      object.Object
	velocity vector.Vector3d
	mass     float64
}

var ph nbody.Body = (*PhysicalBody)(nil)

func NewPhysicalBody(obj object.Object, velocity vector.Vector3d, mass float64) *PhysicalBody {
	b := &PhysicalBody{obj: obj, mass: mass}
	b.SetVelocity(velocity)
	return b
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
	oldVelocity := b.velocity.Copy()
	b.velocity = velocity
	newVelocity := b.velocity.Copy()

	if vector.IsEqual(oldVelocity, &vector.Vector3d{}) {
		oldVelocity = vector.NewVector3d(0, 0, 1)
	}
	if vector.IsEqual(newVelocity, &vector.Vector3d{}) {
		newVelocity = vector.NewVector3d(0, 0, 1)
	}
	oldVelocity.Normalize()
	newVelocity.Normalize()
	if vector.IsEqual(oldVelocity, newVelocity) {
		return
	}
	angleCos := oldVelocity.Dot(newVelocity)
	if angleCos < -1 || angleCos > 1 {
		angleCos = 1
	}
	angle := math.Acos(angleCos)
	axis := vector.CrossProduct(oldVelocity, newVelocity)
	axis.Normalize()
	center := b.GetPosition()
	rotate := transform.NewAxisRotateActionCenter(axis, angle, &center)
	b.obj.Transform(rotate)
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
