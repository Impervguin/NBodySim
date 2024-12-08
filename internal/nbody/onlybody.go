package nbody

import (
	"NBodySim/internal/mathutils/vector"
)

type OnlyBody struct {
	id       int64
	position vector.Vector3d
	velocity vector.Vector3d
	mass     float64
}

var onlyBodyId int64 = 1

func getOnlyBodyId() int64 {
	onlyBodyId++
	return onlyBodyId
}

func NewOnlyBody(position, velocity vector.Vector3d, mass float64) *OnlyBody {
	return &OnlyBody{id: getOnlyBodyId(), position: position, velocity: velocity, mass: mass}
}

func (b *OnlyBody) GetId() int64 {
	return b.id
}

func (b *OnlyBody) GetPosition() vector.Vector3d {
	return b.position
}

func (b *OnlyBody) SetVelocity(velocity vector.Vector3d) {
	b.velocity = velocity
}

func (b *OnlyBody) SetPosition(position vector.Vector3d) {
	b.position = position
}

func (b *OnlyBody) GetVelocity() vector.Vector3d {
	return b.velocity
}

func (b *OnlyBody) GetMass() float64 {
	return b.mass
}

func (b *OnlyBody) Clone() Body {
	return &OnlyBody{
		id:       b.id,
		position: *b.position.Copy(),
		velocity: *b.velocity.Copy(),
		mass:     b.mass,
	}
}

func (b *OnlyBody) SetMass(mass float64) {
	b.mass = mass
}
