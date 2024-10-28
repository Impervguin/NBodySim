package builder

import "NBodySim/internal/object"

type Director interface {
	Construct() (object.Object, error)
}

type Builder interface {
	getObject() (object.Object, error)
}
