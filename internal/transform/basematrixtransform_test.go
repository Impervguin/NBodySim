package transform

import (
	"NBodySim/internal/mathutils/vector"
	"testing"
)

func TestApplyMatrixTransform(t *testing.T) {
	v := vector.NewVector3d(1, 2, 3)
	matrix := vector.NewMatrix4d(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16)
	vc := v.ToHomoVector()

	btmatrix := NewBaseMatrixTransform()
	btmatrix.matrix = *matrix
	btmatrix.ApplyToVector(v)

	vres := matrix.RightMultiply(vc)
	vres3 := vres.ToVector3d()
	if vres3.X != v.X || vres3.Y != v.Y || vres3.Z != v.Z {
		t.Errorf("Expected %v, got %v", vres3, v)
	}
}

func TestApplyToHomoVector(t *testing.T) {
	v := vector.NewVector3d(1, 2, 3)
	matrix := vector.NewMatrix4d(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16)
	vc := v.ToHomoVector()

	btmatrix := NewBaseMatrixTransform()
	btmatrix.matrix = *matrix
	btmatrix.ApplyToHomoVector(vc)

	vres := matrix.RightMultiply(v.ToHomoVector())
	if vres.X != vc.X || vres.Y != vc.Y || vres.Z != vc.Z || vres.W != vc.W {
		t.Errorf("Expected %v, got %v", vres, vc)
	}
}
