package object

import (
	"NBodySim/internal/mathutils/vector"
)

// points must be in camera system
type PerspectiveTransform struct {
	camera *Camera
}

func NewPerspectiveTransform(camera *Camera) *PerspectiveTransform {
	return &PerspectiveTransform{
		camera: camera,
	}
}

func (act *PerspectiveTransform) ApplyToVector(vector *vector.Vector3d) {
	if vector.Z != 0 {
		vector.X = vector.X * act.camera.GetPerspectiveXYModifier() / vector.Z
		vector.Y = vector.Y * act.camera.GetPerspectiveXYModifier() / vector.Z
	} else {
		vector.X = 0
		vector.Y = 0
	}
}
func (act *PerspectiveTransform) ApplyToHomoVector(homoVector *vector.HomoVector) {
	if homoVector.Z != 0 {
		homoVector.X = homoVector.X * act.camera.GetPerspectiveXYModifier() / homoVector.Z * 2 * homoVector.W
		homoVector.Y = homoVector.Y * act.camera.GetPerspectiveXYModifier() / homoVector.Z * 2 * homoVector.W
	} else {
		homoVector.X = 0
		homoVector.Y = 0
	}
}
