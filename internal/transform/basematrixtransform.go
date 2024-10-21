package transform

import "NBodySim/internal/vectormath"

type baseMatrixTransform struct {
	matrix vectormath.Matrix4d
}

func NewBaseMatrixTransform() *baseMatrixTransform {
	return &baseMatrixTransform{matrix: *vectormath.NewEyeMatrix4d()}
}

func (t *baseMatrixTransform) ApplyToVector(vector *vectormath.Vector3d) {
	homo := vector.ToHomoVector()
	res := t.matrix.RightMultiply(homo)
	*vector = *(res.ToVector3d())
}

func (t *baseMatrixTransform) ApplyToHomoVector(homoPoint *vectormath.HomoVector) {
	res := t.matrix.RightMultiply(homoPoint)
	*homoPoint = *res
}

func (t *baseMatrixTransform) GetMatrix() *vectormath.Matrix4d {
	return &t.matrix
}

func (t *baseMatrixTransform) Modify(matrix *vectormath.Matrix4d) {
	t.matrix.Multiply(matrix)
}
