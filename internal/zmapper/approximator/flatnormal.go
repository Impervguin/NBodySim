package approximator

import (
	"NBodySim/internal/mathutils"
	"NBodySim/internal/object"
)

type FlatNormalApproximator struct{}

type FlatNormalApproximatorFabric struct{}

func NewFlatNormalApproximatorFabric() FlatNormalApproximatorFabric {
	return FlatNormalApproximatorFabric{}
}

func (f FlatNormalApproximatorFabric) CreateDiscreteApproximator() DiscreteNormalApproximator {
	return &FlatNormalApproximator{}
}

func (a FlatNormalApproximator) ApproximatePolygon(pol *object.Polygon, ch chan<- DiscreteNormalPoint) error {
	p1, p2, p3 := pol.GetVertices()

	color := pol.GetColor()

	if p1.Y < p2.Y {
		p1, p2 = p2, p1
	}
	if p1.Y < p3.Y {
		p1, p3 = p3, p1
	}
	if p2.Y < p3.Y {
		p2, p3 = p3, p2
	}

	p12 := mathutils.LinearYIntInterpolation(mathutils.ToInt(p1.X), mathutils.ToInt(p1.Y), mathutils.ToInt(p2.X), mathutils.ToInt(p2.Y))
	p13 := mathutils.LinearYIntInterpolation(mathutils.ToInt(p1.X), mathutils.ToInt(p1.Y), mathutils.ToInt(p3.X), mathutils.ToInt(p3.Y))
	p23 := mathutils.LinearYIntInterpolation(mathutils.ToInt(p2.X), mathutils.ToInt(p2.Y), mathutils.ToInt(p3.X), mathutils.ToInt(p3.Y))

	p23 = p23[:len(p23)-1]
	p123 := append(p23, p12...)

	_, z12 := mathutils.LinearXInterpolation(mathutils.ToInt(p1.Y), p1.Z, mathutils.ToInt(p2.Y), p2.Z)
	_, z13 := mathutils.LinearXInterpolation(mathutils.ToInt(p1.Y), p1.Z, mathutils.ToInt(p3.Y), p3.Z)
	_, z23 := mathutils.LinearXInterpolation(mathutils.ToInt(p2.Y), p2.Z, mathutils.ToInt(p3.Y), p3.Z)

	z23 = z23[:len(z23)-1]
	z123 := append(z23, z12...)

	med := len(p123) / 2
	if p123[med][0] < p13[med][0] {
		for i := range p123 {
			z := z123[i]
			dz := (z13[i] - z123[i]) / float64(p13[i][0]-p123[i][0])
			for x := p123[i][0]; x <= p13[i][0]; x++ {
				ch <- DiscreteNormalPoint{X: x, Y: p123[i][1], Z: z, Color: color, Normal: *pol.GetNormal()}
				z += dz
			}
		}
	} else {
		for i := range p123 {
			z := z13[i]
			dz := (z123[i] - z13[i]) / float64(p123[i][0]-p13[i][0])
			for x := p13[i][0]; x <= p123[i][0]; x++ {
				ch <- DiscreteNormalPoint{X: x, Y: p13[i][1], Z: z, Color: color, Normal: *pol.GetNormal()}
				z += dz
			}
		}
	}
	return nil
}
