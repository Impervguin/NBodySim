package gui

import (
	"NBodySim/internal/object"
)

type CameraManager interface {
	SetCamera(camera *object.Camera)
	GetCamera() *object.Camera
	MoveCamera(dx, dy, dz float64)
	RotateRight(angle float64) // in degrees
	RotateUp(angle float64)    // in degrees
}
