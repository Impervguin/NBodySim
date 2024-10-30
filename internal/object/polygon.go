package object

import (
	"NBodySim/internal/vectormath"
	"image/color"
)

type Polygon struct {
	v1, v2, v3 *vectormath.Vector3d
	color      color.Color
}

func NewPolygon(v1, v2, v3 *vectormath.Vector3d, color color.Color) *Polygon {
	return &Polygon{
		v1:    v1,
		v2:    v2,
		v3:    v3,
		color: color,
	}
}

func (p *Polygon) GetVertices() (*vectormath.Vector3d, *vectormath.Vector3d, *vectormath.Vector3d) {
	return p.v1, p.v2, p.v3
}

func (p *Polygon) SetVertices(v1, v2, v3 *vectormath.Vector3d) {
	p.v1, p.v2, p.v3 = v1, v2, v3
}

func (p *Polygon) GetColor() color.Color {
	return p.color
}
