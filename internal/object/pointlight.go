package object

import (
	"NBodySim/internal/mathutils"
	"NBodySim/internal/mathutils/vector"
	"NBodySim/internal/transform"
	"image/color"
	"math"
)

const diffuseCoefficient = 1
const minimalDistance = 100
const diffuseDistanceFactor float64 = 1. / minimalDistance

type PointLight struct {
	ObjectWithId
	InvisibleObject
	intensity color.RGBA
	position  vector.Vector3d
}

// func MirrorVector3d(normal, v vector.Vector3d) vector.Vector3d {
// 	normal.Normalize()

// }

func NewPointLight(intensity color.Color, position vector.Vector3d) *PointLight {
	p := &PointLight{
		intensity: mathutils.ToRGBA(intensity),
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
	p.intensity = mathutils.ToRGBA(intensity)
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

func (p *PointLight) CalculateLightContribution(point, view, normal vector.Vector3d, c color.Color) color.Color {
	// dirToCam := vector.SubtractVectors(&view, &point)
	lightVector := vector.SubtractVectors(&p.position, &point)
	distance := math.Sqrt(lightVector.Square())
	lightVector.Normalize()
	diffuse := math.Abs(lightVector.Dot(&normal)*diffuseCoefficient) / ((minimalDistance + distance) * diffuseDistanceFactor)
	diff := color.RGBA{
		R: uint8(float64(p.intensity.R) * diffuse),
		G: uint8(float64(p.intensity.G) * diffuse),
		B: uint8(float64(p.intensity.B) * diffuse),
		A: p.intensity.A,
	}

	res := diff

	r, g, b, a := c.RGBA()
	t := color.RGBA{
		R: uint8(float64(r>>8) * (float64(res.R) / 255)),
		G: uint8(float64(g>>8) * (float64(res.G) / 255)),
		B: uint8(float64(b>>8) * (float64(res.B) / 255)),
		A: uint8(a >> 8),
	}

	return t
}
