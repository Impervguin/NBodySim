package approximator

import (
	"NBodySim/internal/object"
	"image/color"
)

type DiscreteFlatPoint struct {
	X, Y  int
	Z     float64
	Color color.Color
}

type DiscreteApproximator interface {
	ApproximatePolygon(pol *object.Polygon, ch chan<- DiscreteFlatPoint)
}

type DiscreteApproximatorFabric interface {
	CreateDiscreteApproximator() DiscreteApproximator
}
