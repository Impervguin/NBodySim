package transform

import "NBodySim/internal/mathutils/vector"

type ViewportToCanvas struct {
	BaseMatrixTransform
}

func NewViewportToCanvas(cx, cy float64) *ViewportToCanvas {
	scaleMatrix := vector.NewMatrix4d(
		cx, 0, 0, 0,
		0, cy, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1,
	)

	action := NewBaseMatrixTransform()
	action.Modify(scaleMatrix)

	return &ViewportToCanvas{BaseMatrixTransform: *action}
}
