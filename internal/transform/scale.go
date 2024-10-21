package transform

import "NBodySim/internal/vectormath"

type ScaleAction struct {
	baseMatrixTransform
}

func NewScaleAction(scale *vectormath.Vector3d) *ScaleAction {
	base := NewBaseMatrixTransform()

	scaleMatrix := vectormath.NewMatrix4d(
		scale.X, 0, 0, 0,
		0, scale.Y, 0, 0,
		0, 0, scale.Z, 0,
		0, 0, 0, 1,
	)

	base.matrix = *base.matrix.Multiply(scaleMatrix)

	return &ScaleAction{baseMatrixTransform: *base}
}

func NewScaleActionCenter(center *vectormath.Vector3d, scale *vectormath.Vector3d) *ScaleAction {
	toCenter := NewMoveAction(vectormath.MultiplyVectorScalar(center, -1))
	scaleAction := NewScaleAction(scale)
	toOrigin := NewMoveAction(center)

	base := NewBaseMatrixTransform()

	base.matrix = *((toCenter.GetMatrix().Multiply(&scaleAction.matrix)).Multiply(toOrigin.GetMatrix()))

	return &ScaleAction{baseMatrixTransform: *base}
}
