package object

import (
	"NBodySim/internal/mathutils/vector"
	"NBodySim/internal/transform"
)

// points must be in camera system
type PerspectiveTransform struct {
	camera *Camera
	after  TransformAction
}

func NewPerspectiveTransform(camera *Camera) *PerspectiveTransform {
	return &PerspectiveTransform{
		camera: camera,
		after:  nil,
	}
}

func (act *PerspectiveTransform) ApplyToVector(vector *vector.Vector3d) {
	if vector.Z != 0 {
		vector.X = vector.X * act.camera.GetDistance() / (vector.Z * act.camera.GetWidth())
		vector.Y = vector.Y * act.camera.GetDistance() / (vector.Z * act.camera.GetHeight())
		vector.Z = 1 + 2*act.camera.GetDistance()/vector.Z
	}
	if act.after != nil {
		act.after.ApplyToVector(vector)
	}
}
func (act *PerspectiveTransform) ApplyToHomoVector(homoVector *vector.HomoVector) {
	if homoVector.Z != 0 {
		homoVector.X = homoVector.X * act.camera.GetPerspectiveXYModifier() / homoVector.Z * 2 * homoVector.W
		homoVector.Y = homoVector.Y * act.camera.GetPerspectiveXYModifier() / homoVector.Z * 2 * homoVector.W
	}
	if act.after != nil {
		act.after.ApplyToHomoVector(homoVector)
	}

}

func (act *PerspectiveTransform) ApplyAfter(tr transform.TransformAction) {
	act.after = tr
}

type ReversePerspectiveTransform struct {
	camera *Camera
	after  TransformAction
}

func NewReversePerspectiveTransform(camera *Camera) *ReversePerspectiveTransform {
	return &ReversePerspectiveTransform{
		camera: camera,
		after:  nil,
	}
}

func (act *ReversePerspectiveTransform) ApplyToVector(vector *vector.Vector3d) {
	if vector.Z != 1 {
		vector.Z = 2 * act.camera.GetDistance() / (vector.Z - 1)
		vector.X = vector.X * (vector.Z * act.camera.GetWidth()) / (act.camera.GetPerspectiveXYModifier())
		vector.Y = vector.Y * (vector.Z * act.camera.GetHeight()) / (act.camera.GetPerspectiveXYModifier())
	}
	if act.after != nil {
		act.after.ApplyToVector(vector)
	}
}

func (act *ReversePerspectiveTransform) ApplyToHomoVector(homoVector *vector.HomoVector) {
	if homoVector.Z != 0 {
		homoVector.X = homoVector.X * homoVector.Z / (act.camera.GetPerspectiveXYModifier() * 2 * homoVector.W)
		homoVector.Y = homoVector.Y * homoVector.Z / (act.camera.GetPerspectiveXYModifier() * 2 * homoVector.W)
	}
	if act.after != nil {
		act.after.ApplyToHomoVector(homoVector)
	}
}

func (act *ReversePerspectiveTransform) ApplyAfter(tr transform.TransformAction) {
	act.after = tr
}
