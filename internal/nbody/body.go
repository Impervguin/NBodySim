package nbody

import (
	"NBodySim/internal/mathutils/vector"
)

type Body interface {
	GetId() int64
	Clone() Body
	GetPosition() vector.Vector3d
	GetVelocity() vector.Vector3d
	GetMass() float64
	SetPosition(position vector.Vector3d)
	SetVelocity(velocity vector.Vector3d)
	SetMass(mass float64)
}
