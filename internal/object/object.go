package object

import (
	"NBodySim/internal/transform"
	"NBodySim/internal/vectormath"
)

type TransformAction interface {
	ApplyToVector(vector *vectormath.Vector3d)
	ApplyToHomoVector(homoPoint *vectormath.HomoVector)
}

type Object interface {
	GetId() int64
	GetCenter() vectormath.Vector3d
	Clone() Object
	Transform(action transform.TransformAction)
	IsVisible() bool
	Accept(visitor ObjectVisitor)
}
