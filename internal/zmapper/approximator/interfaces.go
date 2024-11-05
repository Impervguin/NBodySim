package approximator

import (
	"NBodySim/internal/mathutils/vector"
	"NBodySim/internal/object"
	"NBodySim/internal/zmapper/approximator/colorist"
	"image/color"
)

type DiscreteFlatPoint struct {
	X, Y  int
	Z     float64
	Color color.Color
}

type DiscreteApproximator interface {
	ApproximatePolygon(pol *object.Polygon, ch chan<- DiscreteFlatPoint) error
}

type DiscreteApproximatorFabric interface {
	CreateDiscreteApproximator(view vector.Vector3d) DiscreteApproximator
	GetColorist(view vector.Vector3d) colorist.Colorist
}
