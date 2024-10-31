package transform

import "NBodySim/internal/mathutils/vector"

type MoveAction struct {
	BaseMatrixTransform
}

func NewMoveAction(translation *vector.Vector3d) *MoveAction {
	base := NewBaseMatrixTransform()

	moveMatrix := vector.NewMatrix4d(
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		translation.X, translation.Y, translation.Z, 1,
	)

	base.matrix = *base.matrix.Multiply(moveMatrix)

	return &MoveAction{BaseMatrixTransform: *base}
}
