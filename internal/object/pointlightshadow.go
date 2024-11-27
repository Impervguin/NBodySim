package object

import (
	"NBodySim/internal/mathutils/normal"
	"NBodySim/internal/mathutils/vector"
	"image/color"
	"math"
)

type ShadowModel interface {
	PointInShadow(p vector.Vector3d) bool
	SurfacePointInShadow(p vector.Vector3d, normal vector.Vector3d) bool
}

type NoShadowModel struct{}

func (m *NoShadowModel) PointInShadow(p vector.Vector3d) bool {
	return false
}

func (m *NoShadowModel) SurfacePointInShadow(p vector.Vector3d, normal vector.Vector3d) bool {
	return false
}

type PointLightShadow struct {
	PointLight
	shadow ShadowModel
}

func NewPointLightShadow(intensity color.Color, position vector.Vector3d) *PointLightShadow {
	return &PointLightShadow{
		PointLight: *NewPointLight(intensity, position),
		shadow:     &NoShadowModel{},
	}
}

func (p *PointLightShadow) Accept(visitor LightVisitor) {
	visitor.VisitPointLightShadow(p)
}

func (p *PointLightShadow) SetShadowModel(m ShadowModel) {
	p.shadow = m
}

func (p *PointLightShadow) CalculateLightContribution(point, view vector.Vector3d, normal normal.Normal, c color.Color) color.RGBA64 {
	if _, ok := p.shadow.(*NoShadowModel); !ok {
		n := normal.ToVector()
		n.Normalize()
		lightVector := vector.SubtractVectors(&point, &p.position)
		distance := math.Sqrt(lightVector.Square())
		lightVector.Normalize()
		if (normal.NormalIsInner && n.Dot(lightVector) <= 0) || p.shadow.SurfacePointInShadow(point, n) {
			return p.applyLight(p.calculateAmbientPart(distance), c)
		}
	}

	return p.PointLight.CalculateLightContribution(point, view, normal, c)
}

func (p *PointLightShadow) Clone() Light {
	pcl := p.PointLight.Clone().(*PointLight)
	cl := &PointLightShadow{
		PointLight: *pcl,
		shadow:     p.shadow,
	}
	return cl
}
