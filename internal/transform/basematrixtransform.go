package transform

import "NBodySim/internal/mathutils/vector"

type BaseMatrixTransform struct {
	matrix vector.Matrix4d
}

func NewBaseMatrixTransform() *BaseMatrixTransform {
	return &BaseMatrixTransform{matrix: *vector.NewEyeMatrix4d()}
}

func (t *BaseMatrixTransform) ApplyToVector(vector *vector.Vector3d) {
	homo := vector.ToHomoVector()
	res := t.matrix.RightMultiply(homo)
	*vector = *(res.ToVector3d())
}

func (t *BaseMatrixTransform) ApplyToHomoVector(homoPoint *vector.HomoVector) {
	res := t.matrix.RightMultiply(homoPoint)
	*homoPoint = *res
}

func (t *BaseMatrixTransform) GetMatrix() *vector.Matrix4d {
	return &t.matrix
}

func (t *BaseMatrixTransform) Modify(matrix *vector.Matrix4d) {
	t.matrix = *t.matrix.Multiply(matrix)
}
