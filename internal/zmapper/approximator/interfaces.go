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

type DiscreteNormalPoint struct {
	X, Y   int
	Z      float64
	Normal vector.Vector3d
	Color  color.Color
}

type DiscreteApproximator interface {
	ApproximatePolygon(pol *object.Polygon, ch chan<- DiscreteFlatPoint) error
}

type DiscreteApproximatorFabric interface {
	CreateDiscreteApproximator() DiscreteApproximator
	GetColorist() colorist.Colorist
}

type DiscreteNormalApproximator interface {
	ApproximatePolygon(pol *object.Polygon, ch chan<- DiscreteNormalPoint) error
}

type DiscreteNormalApproximatorFabric interface {
	CreateDiscreteApproximator() DiscreteNormalApproximator
}
