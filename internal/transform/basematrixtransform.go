package transform

import "NBodySim/internal/vectormath"

type BaseMatrixTransform struct {
	matrix vectormath.Matrix4d
}

func NewBaseMatrixTransform() *BaseMatrixTransform {
	return &BaseMatrixTransform{matrix: *vectormath.NewEyeMatrix4d()}
}

func (t *BaseMatrixTransform) ApplyToVector(vector *vectormath.Vector3d) {
	homo := vector.ToHomoVector()
	res := t.matrix.RightMultiply(homo)
	*vector = *(res.ToVector3d())
}

func (t *BaseMatrixTransform) ApplyToHomoVector(homoPoint *vectormath.HomoVector) {
	res := t.matrix.RightMultiply(homoPoint)
	*homoPoint = *res
}

func (t *BaseMatrixTransform) GetMatrix() *vectormath.Matrix4d {
	return &t.matrix
}

func (t *BaseMatrixTransform) Modify(matrix *vectormath.Matrix4d) {
	t.matrix = *t.matrix.Multiply(matrix)
}
