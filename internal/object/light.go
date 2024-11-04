package object

import (
	"NBodySim/internal/mathutils/vector"
	"image/color"
)

type Light interface {
	GetId() int64
	Intensity() color.Color
	CalculateLightContribution(point, view vector.Vector3d, color color.Color) color.Color
	Clone() Light
	Accept(visitor LightVisitor)
}
