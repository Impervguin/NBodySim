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
	normalv1   *vector.Vector3d
}

func NewPolygon(v1, v2, v3 *vector.Vector3d, color color.Color) *Polygon {
	p := &Polygon{
		v1:         v1,
		v2:         v2,
		v3:         v3,
		color:      color,
		colorModel: nil,
	}
	p.normalv1 = p.calculateNormal()
	p.normalv1.Add(v1)
	return p
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

func (p *Polygon) SetColor(color color.Color) {
	p.color = color
}

func (p *Polygon) calculateNormal() *vector.Vector3d {
	v1 := vector.SubtractVectors(p.v2, p.v1)
	v2 := vector.SubtractVectors(p.v3, p.v1)
	n := v1.Cross(v2)
	n.Normalize()
	return n
}

func (p *Polygon) GetNormal() *vector.Vector3d {
	n := vector.SubtractVectors(p.normalv1, p.v1)
	n.Normalize()
	return n
}

func (p *Polygon) Transform(t TransformAction) {
	t.ApplyToVector(p.v1)
	t.ApplyToVector(p.v2)
	t.ApplyToVector(p.v3)
	t.ApplyToVector(p.normalv1)
}

func (p *Polygon) TransformNormal(t TransformAction) {
	t.ApplyToVector(p.normalv1)
}

func (p *Polygon) Clone() *Polygon {
	return &Polygon{
		v1:         p.v1.Copy(),
		v2:         p.v2.Copy(),
		v3:         p.v3.Copy(),
		color:      p.color,
		colorModel: p.colorModel,
		normalv1:   p.normalv1.Copy(),
	}
}

func (p *Polygon) SetColorModel(color PolygonColorModel) {
	p.colorModel = color
}

func (p *Polygon) GetColorModel() PolygonColorModel {
	return p.colorModel
}
