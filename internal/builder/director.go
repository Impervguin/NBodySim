package builder

import (
	"NBodySim/internal/object"
	"image/color"
)

var DefaultObjectColor = color.White

type Director interface {
	Construct() (object.Object, error)
}

type Builder interface {
	getObject() (object.Object, error)
}
