package transform

import "NBodySim/internal/vectormath"

type MoveAction struct {
	baseMatrixTransform
}

func NewMoveAction(translation *vectormath.Vector3d) *MoveAction {
	base := NewBaseMatrixTransform()

	moveMatrix := vectormath.NewMatrix4d(
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		translation.X, translation.Y, translation.Z, 1,
	)

	base.matrix = *base.matrix.Multiply(moveMatrix)

	return &MoveAction{baseMatrixTransform: *base}
}
