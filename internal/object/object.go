package object

import (
	"NBodySim/internal/mathutils/vector"
	"NBodySim/internal/transform"
)

type TransformAction interface {
	ApplyToVector(vector *vector.Vector3d)
	ApplyToHomoVector(homoPoint *vector.HomoVector)
}

type Object interface {
	GetId() int64
	GetCenter() vector.Vector3d
	Clone() Object
	Transform(action transform.TransformAction)
	IsVisible() bool
	Accept(visitor ObjectVisitor)
}
