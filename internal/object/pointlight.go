package object

import (
	"NBodySim/internal/mathutils/vector"
	"NBodySim/internal/transform"
	"image/color"
)

type PointLight struct {
	ObjectWithId
	InvisibleObject
	intensity color.Color
	position  vector.Vector3d
}

func NewPointLight(intensity color.Color, position vector.Vector3d) *PointLight {
	p := &PointLight{
		intensity: intensity,
		position:  position,
	}
	p.id = getNextId()
	return p
}

func (p *PointLight) Accept(visitor LightVisitor) {
	visitor.VisitPointLight(p)
}

func (p *PointLight) Intensity() color.Color {
	return p.intensity
}

func (p *PointLight) GetPosition() vector.Vector3d {
	return p.position
}

func (p *PointLight) SetPosition(position vector.Vector3d) {
	p.position = position
}

func (p *PointLight) SetIntensity(intensity color.Color) {
	p.intensity = intensity
}

func (p *PointLight) GetId() int64 {
	return p.id
}

func (p *PointLight) GetCenter() vector.Vector3d {
	return p.position
}

func (p *PointLight) Clone() Light {
	return NewPointLight(p.intensity, *p.position.Copy())
}

func (p *PointLight) Transform(action transform.TransformAction) {
	action.ApplyToVector(&p.position)
}

func (p *PointLight) CalculateLightContribution(point, view vector.Vector3d, color color.Color) color.Color {
	// lightVector := vector.SubtractVectors(&point, &p.position)
	// lightVector.Normalize()

	// distance := vector.Distance(point, p.position)
	// attenuation := 1 / (p.attenuationConstant + distance*distance*p.attenuationLinear + distance*distance*distance*p.attenuationQuadratic)

	// diffuse := vector.DotProduct(lightVector, &view) * attenuation
	// if diffuse < 0 {
	// 	diffuse = 0
	// }

	return p.intensity
}
