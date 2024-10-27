package transform

import "NBodySim/internal/vectormath"

type ViewportToCanvas struct {
	BaseMatrixTransform
}

func NewViewportToCanvas(cx, cy float64) *ViewportToCanvas {
	scaleMatrix := vectormath.NewMatrix4d(
		cx, 0, 0, 0,
		0, cy, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1,
	)

	action := NewBaseMatrixTransform()
	action.Modify(scaleMatrix)

	return &ViewportToCanvas{BaseMatrixTransform: *action}
}
