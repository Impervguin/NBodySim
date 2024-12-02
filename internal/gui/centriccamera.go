package gui

import (
	"NBodySim/internal/mathutils"
	"NBodySim/internal/mathutils/vector"
	"NBodySim/internal/object"
	"NBodySim/internal/transform"
)

type CentricCameraManager struct {
	camera             *object.Camera
	distanceFromCenter float64
}

func NewCentricCameraManager(cam *object.Camera) *CentricCameraManager {
	return &CentricCameraManager{
		camera: cam,
	}
}

func (m *CentricCameraManager) SetCamera(cam *object.Camera) {
	m.camera = cam
}

func (m *CentricCameraManager) GetCamera() *object.Camera {
	return m.camera
}

func (m *CentricCameraManager) MoveCamera(_, _, dz float64) {
	moveDistance := dz
	if m.distanceFromCenter-dz < 0.01 {
		moveDistance = m.distanceFromCenter - 0.01
	}
	m.distanceFromCenter -= moveDistance
	forward := m.camera.GetForward()
	m.camera.Transform(transform.NewMoveAction(vector.MultiplyVectorScalar(&forward, moveDistance)))
}

func (m *CentricCameraManager) RotateUp(angle float64) {
	angle = mathutils.ToRadians(angle)
	right := m.camera.GetRight()
	m.camera.Transform(transform.NewAxisRotateAction(&right, -angle))
}

func (m *CentricCameraManager) RotateRight(angle float64) {
	angle = mathutils.ToRadians(angle)
	up := vector.NewVector3d(0, 1, 0)
	m.camera.Transform(transform.NewAxisRotateAction(up, angle))
}
