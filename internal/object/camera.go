package object

import (
	"NBodySim/internal/mathutils/vector"
	"NBodySim/internal/transform"
)

type Camera struct {
	ObjectWithId
	InvisibleObject
	position            vector.Vector3d
	forward             vector.Vector3d
	up                  vector.Vector3d
	px                  float64 // Полуразмер окна просмотра по x(в пространстве камеры)
	py                  float64 // Полуразмер окна просмотра по y(в пространстве камеры)
	d                   float64 // Дальность от позиции камеры, до окна просмотра
	view                CameraViewAction
	modifiedView        bool
	perspective         PerspectiveTransform
	modifiedPerspective bool
}

func NewCamera(position, forward, up vector.Vector3d, px, py, d float64) *Camera {
	forward.Normalize()
	up.Normalize()
	cam := &Camera{
		position: position,
		forward:  forward,
		up:       up,
		px:       px,
		py:       py,
		d:        d,
	}
	cam.view = *NewCameraViewAction(cam)
	cam.modifiedView = false
	cam.perspective = *NewPerspectiveTransform(cam)
	cam.modifiedPerspective = false
	cam.id = getNextId()
	return cam
}

func (c *Camera) GetCenter() vector.Vector3d {
	return c.position
}

func (c *Camera) Clone() Object {
	return NewCamera(c.position, c.forward, c.up, c.px, c.py, c.d)
}

func (c *Camera) Accept(visitor ObjectVisitor) {
	visitor.VisitCamera(c)
}

func (c *Camera) Transform(action transform.TransformAction) {
	beforePos := c.position.Copy()
	action.ApplyToVector(&c.position)

	c.forward.Add(beforePos)
	action.ApplyToVector(&c.forward)
	c.forward.Subtract(&c.position)
	c.forward.Normalize()

	c.up.Add(beforePos)
	action.ApplyToVector(&c.up)
	c.up.Subtract(&c.position)
	c.up.Normalize()

	c.modifiedPerspective = true
	c.modifiedView = true
}

func (c *Camera) GetViewMatrix() *vector.Matrix4d {
	right := c.GetRight()

	viewMatrix := vector.NewMatrix4d(
		-right.X, c.up.X, c.forward.X, 0,
		-right.Y, c.up.Y, c.forward.Y, 0,
		-right.Z, c.up.Z, c.forward.Z, 0,
		0, 0, 0, 1,
	)

	toCameraCenter := transform.NewMoveAction(vector.MultiplyVectorScalar(&c.position, -1))
	toCameraCenter.Modify(viewMatrix)

	return toCameraCenter.GetMatrix()
}

func (c *Camera) GetPerspectiveXYModifier() float64 {
	return c.d
}

func (c *Camera) GetWidth() float64 {
	return c.px
}

func (c *Camera) GetHeight() float64 {
	return c.py
}

func (c *Camera) GetDistance() float64 {
	return c.d
}

func (c *Camera) GetViewAction() *CameraViewAction {
	if !c.modifiedView {
		return &c.view
	}
	c.modifiedView = false
	c.view = *NewCameraViewAction(c)
	return &c.view
}

func (c *Camera) GetPerspectiveTransform() *PerspectiveTransform {
	if !c.modifiedPerspective {
		return &c.perspective
	}
	c.modifiedPerspective = false
	c.perspective = *NewPerspectiveTransform(c)
	return &c.perspective
}

func (c *Camera) GetUp() vector.Vector3d {
	return c.up
}

func (c *Camera) GetForward() vector.Vector3d {
	return c.forward
}

func (c *Camera) GetRight() vector.Vector3d {
	right := vector.CrossProduct(&c.forward, &c.up)
	right.Normalize()
	return *right
}
