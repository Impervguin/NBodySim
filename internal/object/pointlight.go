package object

import (
	"NBodySim/internal/mathutils"
	"NBodySim/internal/mathutils/normal"
	"NBodySim/internal/mathutils/vector"
	"NBodySim/internal/transform"
	"image/color"
	"math"
)

const diffuseCoefficient = 1
const minimalDistance = 100
const diffuseDistanceFactor float64 = 1. / minimalDistance
const ambientCoefficient = 0.15

type PointLight struct {
	ObjectWithId
	InvisibleObject
	intensity color.RGBA64
	position  vector.Vector3d
}

func NewPointLight(intensity color.Color, position vector.Vector3d) *PointLight {
	p := &PointLight{
		intensity: mathutils.ToRGBA64(intensity),
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
	p.intensity = mathutils.ToRGBA64(intensity)
}

func (p *PointLight) GetId() int64 {
	return p.id
}

func (p *PointLight) GetCenter() vector.Vector3d {
	return p.position
}

func (p *PointLight) Clone() Light {
	np := &PointLight{
		position:  p.position,
		intensity: p.intensity,
	}
	np.id = p.GetId()
	return np
}

func (p *PointLight) Transform(action transform.TransformAction) {
	action.ApplyToVector(&p.position)
}

func (p *PointLight) calculateAmbientPart(distance float64) color.RGBA64 {
	ambient := ambientCoefficient / ((minimalDistance + distance) * diffuseDistanceFactor)
	return mathutils.MultRGBA64(p.intensity, ambient)
}

func (p *PointLight) applyLight(light color.RGBA64, objColor color.Color) color.RGBA64 {
	r, g, b, a := objColor.RGBA()
	// t := color.RGBA64{
	// 	R: uint16(float64(r) * (float64(light.R) / 65535)),
	// 	G: uint16(float64(g) * (float64(light.G) / 65535)),
	// 	B: uint16(float64(b) * (float64(light.B) / 65535)),
	// 	A: uint16(a),
	// }
	alt_t := color.RGBA64{
		R: uint16(r * uint32(light.R) / 65535),
		G: uint16(g * uint32(light.G) / 65535),
		B: uint16(b * uint32(light.B) / 65535),
		A: uint16(a),
	}
	return alt_t

}
func (p *PointLight) CalculateLightContribution(point, view vector.Vector3d, normal normal.Normal, c color.Color) color.RGBA64 {
	lightVector := vector.SubtractVectors(&p.position, &point)
	distance := math.Sqrt(lightVector.Square())
	lightVector.Normalize()
	n := normal.ToVector()
	n.Normalize()
	diffuse := math.Abs(lightVector.Dot(&n)*diffuseCoefficient) / ((minimalDistance + distance) * diffuseDistanceFactor)
	diff := mathutils.MultRGBA64(p.intensity, diffuse)

	amb := p.calculateAmbientPart(distance)

	res := diff
	res = mathutils.AddRGBA64(res, amb)
	return p.applyLight(res, c)
}
