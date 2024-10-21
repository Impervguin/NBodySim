package transform

import (
	"NBodySim/internal/vectormath"
	"math"
)

type RotateAction struct {
	baseMatrixTransform
}

func NewRotateAction(ov *vectormath.Vector3d) *RotateAction {
	base := NewBaseMatrixTransform()

	x := vectormath.NewMatrix4d(
		1, 0, 0, 0,
		0, math.Cos(ov.X), math.Sin(ov.X), 0,
		0, -math.Sin(ov.X), math.Cos(ov.X), 0,
		0, 0, 0, 1,
	)

	y := vectormath.NewMatrix4d(
		math.Cos(ov.Y), 0, -math.Sin(ov.Y), 0,
		0, 1, 0, 0,
		math.Sin(ov.Y), 0, math.Cos(ov.Y), 0,
		0, 0, 0, 1,
	)
	z := vectormath.NewMatrix4d(
		math.Cos(ov.Z), math.Sin(ov.Z), 0, 0,
		-math.Sin(ov.Z), math.Cos(ov.Z), 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1,
	)

	base.matrix = *base.matrix.Multiply(x).Multiply(y).Multiply(z)

	return &RotateAction{baseMatrixTransform: *base}
}

func NewRotateActionCenter(center *vectormath.Vector3d, rotate *vectormath.Vector3d) *RotateAction {
	toCenter := NewMoveAction(vectormath.MultiplyVectorScalar(center, -1))
	rotateAction := NewScaleAction(rotate)
	toOrigin := NewMoveAction(center)

	base := NewBaseMatrixTransform()

	base.matrix = *((toCenter.GetMatrix().Multiply(&rotateAction.matrix)).Multiply(toOrigin.GetMatrix()))

	return &RotateAction{baseMatrixTransform: *base}
}
