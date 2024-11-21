package gui

import (
	"NBodySim/internal/mathutils"
	"NBodySim/internal/mathutils/vector"
	"NBodySim/internal/object"
	"NBodySim/internal/transform"
)

type FreeCameraManager struct {
	camera *object.Camera
}

func NewFreeCameraManager(cam *object.Camera) *FreeCameraManager {
	return &FreeCameraManager{
		camera: cam,
	}
}

func (m *FreeCameraManager) SetCamera(cam *object.Camera) {
	m.camera = cam
}

func (m *FreeCameraManager) GetCamera() *object.Camera {
	return m.camera
}

func (m *FreeCameraManager) MoveCamera(dx, dy, dz float64) {
	m.camera.Transform(transform.NewMoveAction(vector.NewVector3d(dx, dy, dz)))
}

func (m *FreeCameraManager) RotateRight(angle float64) {
	angle = mathutils.ToRadians(angle)
	c := m.camera.GetCenter()
	m.camera.Transform(transform.NewMoveAction(vector.MultiplyVectorScalar(&c, -1)))
	up := m.camera.GetUp()
	m.camera.Transform(transform.NewAxisRotateAction(&up, -angle))
	m.camera.Transform(transform.NewMoveAction(&c))
}

func (m *FreeCameraManager) RotateUp(angle float64) {
	angle = mathutils.ToRadians(angle)
	c := m.camera.GetCenter()
	m.camera.Transform(transform.NewMoveAction(vector.MultiplyVectorScalar(&c, -1)))
	right := m.camera.GetRight()
	m.camera.Transform(transform.NewAxisRotateAction(&right, -angle))
	m.camera.Transform(transform.NewMoveAction(&c))
}

func (m *FreeCameraManager) RotateYaw(angle float64) {
	angle = mathutils.ToRadians(angle)
	c := m.camera.GetCenter()
	m.camera.Transform(transform.NewMoveAction(vector.MultiplyVectorScalar(&c, -1)))
	forward := m.camera.GetForward()
	m.camera.Transform(transform.NewAxisRotateAction(&forward, -angle))
	m.camera.Transform(transform.NewMoveAction(&c))
}
