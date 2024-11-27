package object

import (
	"NBodySim/internal/mathutils/normal"
	"NBodySim/internal/mathutils/vector"
	"NBodySim/internal/transform"
	"fmt"
	"image/color"
	"math"
)

type PolygonColorModel interface{}

type Polygon struct {
	v1, v2, v3  *vector.Vector3d
	color       color.Color
	colorModel  PolygonColorModel
	normal      normal.Normal
	normalInner bool
	normalOuter bool
}

func NewPolygon(v1, v2, v3 *vector.Vector3d, color color.Color) *Polygon {
	p := &Polygon{
		v1:          v1,
		v2:          v2,
		v3:          v3,
		color:       color,
		colorModel:  nil,
		normalInner: false,
		normalOuter: false,
	}
	p.normal = *p.calculateNormal()
	return p
}

func NewPolygonInnerNormal(v1, v2, v3 *vector.Vector3d, normal *normal.Normal, color color.Color) (*Polygon, error) {
	p := NewPolygon(v1, v2, v3, color)
	p.normalInner = true
	p.normal = *normal
	p.normal.NormalIsInner = true
	n := p.calculateNormal().ToVector()
	nv := p.normal.ToVector()
	nv.Normalize()
	if math.Abs(n.Dot(&nv)) < 1.-1e-6 {
		return nil, fmt.Errorf("incorrect normal for  given polygon")
	}
	return p, nil
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

func (p *Polygon) calculateNormal() *normal.Normal {
	v1 := vector.SubtractVectors(p.v2, p.v1)
	v2 := vector.SubtractVectors(p.v3, p.v1)
	n := v1.Cross(v2)
	n.Normalize()

	return normal.NewNormal(*v1, *vector.AddVectors(v1, n))
}

func (p *Polygon) GetNormal() *normal.Normal {
	return &p.normal
}

func (p *Polygon) Transform(t transform.TransformAction) {
	t.ApplyToVector(p.v1)
	t.ApplyToVector(p.v2)
	t.ApplyToVector(p.v3)
	p.normal.Transform(t)
}

func (p *Polygon) TransformNormal(t transform.TransformAction) {
	p.normal.Transform(t)
}

func (p *Polygon) Clone() *Polygon {
	return &Polygon{
		v1:         p.v1.Copy(),
		v2:         p.v2.Copy(),
		v3:         p.v3.Copy(),
		color:      p.color,
		colorModel: p.colorModel,
		normal:     *p.normal.Copy(),
	}
}

func (p *Polygon) SetColorModel(color PolygonColorModel) {
	p.colorModel = color
}

func (p *Polygon) GetColorModel() PolygonColorModel {
	return p.colorModel
}

func (p *Polygon) NormalIsInner() bool {
	return p.normalInner
}

func (p *Polygon) SetNormalInner() {
	p.normalInner = true
}

func (p *Polygon) ResetNormalInner() {
	p.normalInner = false
}

func (p *Polygon) NormalIsOuter() bool {
	return p.normalOuter
}

func (p *Polygon) SetNormalOuter() {
	p.normalOuter = true
}

func (p *Polygon) ResetNormalOuter() {
	p.normalOuter = false
}
