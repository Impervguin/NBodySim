package transform

import (
	"NBodySim/internal/mathutils/vector"
)

type TransformAction interface {
	ApplyToVector(vector *vector.Vector3d)
	ApplyToHomoVector(homoPoint *vector.HomoVector)
	ApplyAfter(tr TransformAction)
}
