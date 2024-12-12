package transform

import (
	"NBodySim/internal/mathutils/vector"
	"math"
)

type RotateAction struct {
	BaseMatrixTransform
}

func NewRotateAction(ov *vector.Vector3d) *RotateAction {
	base := NewBaseMatrixTransform()

	x := vector.NewMatrix4d(
		1, 0, 0, 0,
		0, math.Cos(ov.X), math.Sin(ov.X), 0,
		0, -math.Sin(ov.X), math.Cos(ov.X), 0,
		0, 0, 0, 1,
	)

	y := vector.NewMatrix4d(
		math.Cos(ov.Y), 0, -math.Sin(ov.Y), 0,
		0, 1, 0, 0,
		math.Sin(ov.Y), 0, math.Cos(ov.Y), 0,
		0, 0, 0, 1,
	)
	z := vector.NewMatrix4d(
		math.Cos(ov.Z), math.Sin(ov.Z), 0, 0,
		-math.Sin(ov.Z), math.Cos(ov.Z), 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1,
	)

	base.matrix = *base.matrix.Multiply(x).Multiply(y).Multiply(z)

	return &RotateAction{BaseMatrixTransform: *base}
}

func NewRotateActionCenter(center *vector.Vector3d, rotate *vector.Vector3d) *RotateAction {
	toCenter := NewMoveAction(vector.MultiplyVectorScalar(center, -1))
	rotateAction := NewRotateAction(rotate)
	toOrigin := NewMoveAction(center)

	base := NewBaseMatrixTransform()

	base.matrix = *((toCenter.GetMatrix().Multiply(&rotateAction.matrix)).Multiply(toOrigin.GetMatrix()))

	return &RotateAction{BaseMatrixTransform: *base}
}

type AxisRotateAction struct {
	BaseMatrixTransform
}

func NewAxisRotateAction(axis *vector.Vector3d, angle float64) *AxisRotateAction {
	base := NewBaseMatrixTransform()

	cosAngle := math.Cos(angle)
	sinAngle := math.Sin(angle)
	cos1 := 1 - cosAngle
	ux, uy, uz := axis.X, axis.Y, axis.Z

	rot := vector.NewMatrix4d(
		ux*ux*cos1+cosAngle, ux*uy*cos1+uz*sinAngle, ux*uz*cos1-uy*sinAngle, 0,
		ux*uy*cos1-uz*sinAngle, uy*uy*cos1+cosAngle, uy*uz*cos1+ux*sinAngle, 0,
		ux*uz*cos1+uy*sinAngle, uy*uz*cos1-ux*sinAngle, uz*uz*cos1+cosAngle, 0,
		0, 0, 0, 1,
	)
	base.matrix = *rot
	return &AxisRotateAction{BaseMatrixTransform: *base}
}

func NewAxisRotateActionCenter(axis *vector.Vector3d, angle float64, center *vector.Vector3d) *AxisRotateAction {
	toCenter := NewMoveAction(vector.MultiplyVectorScalar(center, -1))
	rotateAction := NewAxisRotateAction(axis, angle)
	toOrigin := NewMoveAction(center)
	base := NewBaseMatrixTransform()
	base.matrix = *((toCenter.GetMatrix().Multiply(&rotateAction.matrix)).Multiply(toOrigin.GetMatrix()))
	return &AxisRotateAction{BaseMatrixTransform: *base}
}
