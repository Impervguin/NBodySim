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
	intensity color.RGBA64
	position  vector.Vector3d
}

// func MirrorVector3d(normal, v vector.Vector3d) vector.Vector3d {
// 	normal.Normalize()

// }

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

func (p *PointLight) CalculateLightContribution(point, view, normal vector.Vector3d, c color.Color) color.RGBA64 {
	// dirToCam := vector.SubtractVectors(&view, &point)
	lightVector := vector.SubtractVectors(&p.position, &point)
	distance := math.Sqrt(lightVector.Square())
	lightVector.Normalize()
	diffuse := math.Abs(lightVector.Dot(&normal)*diffuseCoefficient) / ((minimalDistance + distance) * diffuseDistanceFactor)
	diff := mathutils.MultRGBA64(p.intensity, diffuse)

	res := diff

	r, g, b, a := c.RGBA()
	// if r < 1000 {
	// fmt.Println(r)
	// }
	t := color.RGBA64{
		R: uint16(float64(r) * (float64(res.R) / 65535)),
		G: uint16(float64(g) * (float64(res.G) / 65535)),
		B: uint16(float64(b) * (float64(res.B) / 65535)),
		A: uint16(a),
	}
	// if t.R < 1000 {
	// 	fmt.Println(lightVector.Dot(&normal))
	// }

	return t
}
