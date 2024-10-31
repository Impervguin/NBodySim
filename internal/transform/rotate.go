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
