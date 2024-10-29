package nbody

import (
	"NBodySim/internal/vectormath"
)

type Body interface {
	GetId() int64
	Clone() Body
	GetPosition() vectormath.Vector3d
	GetVelocity() vectormath.Vector3d
	GetMass() float64
	SetPosition(position vectormath.Vector3d)
	SetVelocity(velocity vectormath.Vector3d)
	SetMass(mass float64)
}
