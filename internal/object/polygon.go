package object

import (
	"NBodySim/internal/mathutils/vector"
	"image/color"
)

type PolygonColorModel interface{}

type Polygon struct {
	v1, v2, v3 *vector.Vector3d
	color      color.Color
	colorModel PolygonColorModel
}

func NewPolygon(v1, v2, v3 *vector.Vector3d, color color.Color) *Polygon {
	return &Polygon{
		v1:         v1,
		v2:         v2,
		v3:         v3,
		color:      color,
		colorModel: nil,
	}
}

func (p *Polygon) GetVertices() (*vector.Vector3d, *vector.Vector3d, *vector.Vector3d) {
	return p.v1, p.v2, p.v3
}

func (p *Polygon) SetVertices(v1, v2, v3 *vector.Vector3d) {
	p.v1, p.v2, p.v3 = v1, v2, v3
}

func (p *Polygon) GetColor() color.Color {
	return p.color
}

func (p *Polygon) GetNormal() *vector.Vector3d {
	v1 := vector.SubtractVectors(p.v2, p.v1)
	v2 := vector.SubtractVectors(p.v3, p.v1)
	return v1.Cross(v2)
}

func (p *Polygon) Transform(t TransformAction) {
	t.ApplyToVector(p.v1)
	t.ApplyToVector(p.v2)
	t.ApplyToVector(p.v3)
}

func (p *Polygon) Clone() *Polygon {
	return NewPolygon(p.v1.Copy(), p.v2.Copy(), p.v3.Copy(), p.color)
}

func (p *Polygon) SetColorModel(color PolygonColorModel) {
	p.colorModel = color
}

func (p *Polygon) GetColorModel() PolygonColorModel {
	return p.colorModel
}
