package object

import (
	"NBodySim/internal/transform"
)

type CameraViewAction struct {
	transform.BaseMatrixTransform
}

func NewCameraViewAction(cam *Camera) *CameraViewAction {
	base := transform.NewBaseMatrixTransform()

	base.Modify(cam.GetViewMatrix())
	return &CameraViewAction{
		BaseMatrixTransform: *base,
	}
}
