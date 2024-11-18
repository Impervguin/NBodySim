package object

import (
	"NBodySim/internal/mathutils/vector"
	"NBodySim/internal/transform"
	"image/color"
)

type Light interface {
	GetId() int64
	Intensity() color.Color
	CalculateLightContribution(point, view, normal vector.Vector3d, color color.Color) color.RGBA64
	Clone() Light
	Accept(visitor LightVisitor)
	Transform(action transform.TransformAction)
	GetCenter() vector.Vector3d
}
