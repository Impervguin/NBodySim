package approximator

import (
	"NBodySim/internal/mathutils"
	"NBodySim/internal/object"
	"NBodySim/internal/zmapper/approximator/colorist"
	"fmt"
	"image/color"
)

type GuroApproximator struct {
}

type GuroApproximatorFabric struct {
}

func NewGuroApproximatorFabric() *GuroApproximatorFabric {
	return &GuroApproximatorFabric{}
}

func (ga *GuroApproximatorFabric) CreateDiscreteApproximator() DiscreteApproximator {
	return newGuroApproximator()
}

func (ga *GuroApproximatorFabric) GetColorist() colorist.Colorist {
	return colorist.NewGuroColorist()
}

func (ga *GuroApproximator) ApproximatePolygon(p *object.Polygon, ch chan<- DiscreteFlatPoint) error {
	model := p.GetColorModel()
	guro, ok := model.(*colorist.GuroColorModel)
	if !ok {
		return fmt.Errorf("polygon color model is not a guro color model")
	}

	p1, p2, p3 := p.GetVertices()
	c1, c2, c3 := make(map[int64]color.RGBA, len(guro.C)), make(map[int64]color.RGBA, len(guro.C)), make(map[int64]color.RGBA, len(guro.C))

	for id, color := range guro.C {
		c1[id], c2[id], c3[id] = color.C1, color.C2, color.C3
	}

	if p1.Y < p2.Y {
		p1, p2 = p2, p1
		c1, c2 = c2, c1
	}

	if p1.Y < p3.Y {
		p1, p3 = p3, p1
		c1, c3 = c3, c1
	}

	if p2.Y < p3.Y {
		p2, p3 = p3, p2
		c2, c3 = c3, c2
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

	c12 := make(map[int64][]color.RGBA64, len(c1))
	c13 := make(map[int64][]color.RGBA64, len(c1))
	c23 := make(map[int64][]color.RGBA64, len(c1))
	c123 := make(map[int64][]color.RGBA64, len(c1))
	for id, _ := range c1 {
		c12[id] = mathutils.LinearColorInterpolation(mathutils.ToInt(p1.Y), mathutils.ToInt(p2.Y), c1[id], c2[id])
		c13[id] = mathutils.LinearColorInterpolation(mathutils.ToInt(p1.Y), mathutils.ToInt(p3.Y), c1[id], c3[id])
		c23[id] = mathutils.LinearColorInterpolation(mathutils.ToInt(p2.Y), mathutils.ToInt(p3.Y), c2[id], c3[id])
		c123[id] = append(c23[id], c12[id]...)
	}

	med := len(p123) / 2
	if p123[med][0] < p13[med][0] {
		for i := range p123 {
			z := z123[i]
			c := make(map[int64][]color.RGBA64, len(c1))
			for id, _ := range c1 {
				c[id] = mathutils.LinearColorInterpolation(p123[i][0], p13[i][0], c123[id][i], c13[id][i])
			}
			ci := 0
			dz := 0.
			if p13[i][0] != p123[i][0] {
				dz = (z13[i] - z123[i]) / float64(p13[i][0]-p123[i][0])
			}
			for x := p123[i][0]; x <= p13[i][0]; x++ {
				curcol := color.RGBA64{}
				for id, _ := range c {
					curcol = mathutils.AddRGBA64(curcol, c[id][ci])
				}
				ch <- DiscreteFlatPoint{X: x, Y: p123[i][1], Z: z, Color: curcol}
				ci++
				z += dz
			}
		}
	} else {
		for i := range p123 {
			z := z13[i]
			c := make(map[int64][]color.RGBA64, len(c1))
			for id, _ := range c1 {
				c[id] = mathutils.LinearColorInterpolation(p13[i][0], p123[i][0], c13[id][i], c123[id][i])
			}
			ci := 0
			dz := 0.
			if p13[i][0] != p123[i][0] {
				dz = (z123[i] - z13[i]) / float64(p123[i][0]-p13[i][0])
			}
			for x := p13[i][0]; x <= p123[i][0]; x++ {
				curcol := color.RGBA64{}
				for id, _ := range c {
					curcol = mathutils.AddRGBA64(curcol, c[id][ci])
				}
				ch <- DiscreteFlatPoint{X: x, Y: p13[i][1], Z: z, Color: curcol}
				ci++
				z += dz
			}
		}
	}
	return nil
}

func newGuroApproximator() *GuroApproximator {
	return &GuroApproximator{}
}
