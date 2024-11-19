package object

import (
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

func (p *PointLightShadow) CalculateLightContribution(point, view, normal vector.Vector3d, c color.Color) color.RGBA64 {
	if p.shadow.SurfacePointInShadow(point, normal) {
		lightVector := vector.SubtractVectors(&p.position, &point)
		distance := math.Sqrt(lightVector.Square())
		return p.applyLight(p.calculateAmbientPart(distance), c)
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
