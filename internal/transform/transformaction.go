package transform

import (
	"NBodySim/internal/vectormath"
)

type TransformAction interface {
	ApplyToVector(vector *vectormath.Vector3d) 
	ApplyToHomoVector(homoPoint *vectormath.HomoVector)
}


