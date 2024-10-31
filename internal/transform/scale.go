package transform

import "NBodySim/internal/mathutils/vector"

type ScaleAction struct {
	BaseMatrixTransform
}

func NewScaleAction(scale *vector.Vector3d) *ScaleAction {
	base := NewBaseMatrixTransform()

	scaleMatrix := vector.NewMatrix4d(
		scale.X, 0, 0, 0,
		0, scale.Y, 0, 0,
		0, 0, scale.Z, 0,
		0, 0, 0, 1,
	)

	base.matrix = *base.matrix.Multiply(scaleMatrix)

	return &ScaleAction{BaseMatrixTransform: *base}
}

func NewScaleActionCenter(center *vector.Vector3d, scale *vector.Vector3d) *ScaleAction {
	toCenter := NewMoveAction(vector.MultiplyVectorScalar(center, -1))
	scaleAction := NewScaleAction(scale)
	toOrigin := NewMoveAction(center)

	base := NewBaseMatrixTransform()

	base.matrix = *((toCenter.GetMatrix().Multiply(&scaleAction.matrix)).Multiply(toOrigin.GetMatrix()))

	return &ScaleAction{BaseMatrixTransform: *base}
}
