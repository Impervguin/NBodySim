package simulation

import (
	"NBodySim/internal/nbody"
	"NBodySim/internal/object"
)

type SimulationObject struct {
	obj  object.Object
	body *nbody.Body
}

func NewSimulationObject(obj object.Object, body *nbody.Body) *SimulationObject {
	return &SimulationObject{obj: obj, body: body}
}
