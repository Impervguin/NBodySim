package vector

import "testing"

func TestNewZeroMatrix(t *testing.T) {
	matrix := NewZeroMatrix4d()
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if matrix.Elements[i][j] != 0 {
				t.Error("Expected all elements to be zero, got", matrix.Elements[i][j])
			}
		}
	}
}

func TestNewEyeMatrix(t *testing.T) {
	matrix := NewEyeMatrix4d()
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if i == j && matrix.Elements[i][j] != 1 {
				t.Error("Expected diagonal elements to be 1, got", matrix.Elements[i][j])
			} else if i != j && matrix.Elements[i][j] != 0 {
				t.Error("Expected non-diagonal elements to be 0, got", matrix.Elements[i][j])
			}
		}
	}
}

func TestMultiplyScalarMatrix4d(t *testing.T) {
	matrix := NewMatrix4d(
		1, 2, 3, 4,
		5, 6, 7, 8,
		9, 10, 11, 12,
		13, 14, 15, 16)
	scalar := 2.0
	matrix.MultiplyScalar(scalar)
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if matrix.Elements[i][j] != 2*(float64(i)*4+float64(j+1)) {
				t.Error("Expected matrix elements to be multiplied by scalar, got", matrix.Elements[i][j])
			}
		}
	}
}

func TestDivideScalarMatrix4d(t *testing.T) {
	matrix := NewMatrix4d(
		1, 2, 3, 4,
		5, 6, 7, 8,
		9, 10, 11, 12,
		13, 14, 15, 16)
	scalar := 2.0
	matrix.DivideScalar(scalar)
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if matrix.Elements[i][j] != (float64(i)*4+float64(j+1))/scalar {
				t.Error("Expected matrix elements to be divided by scalar, got", matrix.Elements[i][j])
			}
		}
	}
}

func TestAddScalarMatrix4d(t *testing.T) {
	matrix := NewMatrix4d(
		1, 2, 3, 4,
		5, 6, 7, 8,
		9, 10, 11, 12,
		13, 14, 15, 16)
	scalar := 2.0
	matrix.AddScalar(scalar)
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if matrix.Elements[i][j] != (float64(i)*4+float64(j+1))+scalar {
				t.Error("Expected matrix elements to be added scalar, got", matrix.Elements[i][j])
			}
		}
	}
}

func TestSubtractScalarMatrix4d(t *testing.T) {
	matrix := NewMatrix4d(
		1, 2, 3, 4,
		5, 6, 7, 8,
		9, 10, 11, 12,
		13, 14, 15, 16)
	scalar := 2.0
	matrix.SubtractScalar(scalar)
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if matrix.Elements[i][j] != (float64(i)*4+float64(j+1))-scalar {
				t.Error("Expected matrix elements to be subtracted scalar, got", matrix.Elements[i][j])
			}
		}
	}
}

func TestTransposeMatrix4d(t *testing.T) {
	matrix := NewMatrix4d(
		1, 2, 3, 4,
		5, 6, 7, 8,
		9, 10, 11, 12,
		13, 14, 15, 16)
	expected := NewMatrix4d(
		1, 5, 9, 13,
		2, 6, 10, 14,
		3, 7, 11, 15,
		4, 8, 12, 16)
	matrix.Transpose()
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if matrix.Elements[i][j] != expected.Elements[i][j] {
				t.Error("Expected transposed matrix, got", matrix.Elements[i][j])
			}
		}
	}
}

func TestLeftMultiplyMatrix4d(t *testing.T) {
	matrix := NewMatrix4d(
		1, 2, 3, 4,
		5, 6, 7, 8,
		9, 10, 11, 12,
		13, 14, 15, 16)
	v := NewHomoVector(1, 2, 3, 2)
	expected := NewHomoVector(22, 54, 86, 118)
	result := matrix.LeftMultiply(v)
	if result.X != expected.X || result.Y != expected.Y || result.Z != expected.Z || result.W != expected.W {
		t.Error("Expected left-multiplied vector, got", result)
	}
}

func TestRightMultiplyMatrix4d(t *testing.T) {
	matrix := NewMatrix4d(
		1, 2, 3, 4,
		5, 6, 7, 8,
		9, 10, 11, 12,
		13, 14, 15, 16)
	v := NewHomoVector(1, 2, 3, 2)
	expected := NewHomoVector(64, 72, 80, 88)
	result := matrix.RightMultiply(v)
	if result.X != expected.X || result.Y != expected.Y || result.Z != expected.Z || result.W != expected.W {
		t.Error("Expected right-multiplied vector, got", result)
	}
}

func TestMultiplyMatrix4d(t *testing.T) {
	matrix1 := NewMatrix4d(
		1, 2, 3, 4,
		5, 6, 7, 8,
		9, 10, 11, 12,
		13, 14, 15, 16)
	matrix2 := NewMatrix4d(
		17, 18, 19, 20,
		21, 22, 23, 24,
		25, 26, 27, 28,
		29, 30, 31, 32)
	expected := NewMatrix4d(
		250, 260, 270, 280,
		618, 644, 670, 696,
		986, 1028, 1070, 1112,
		1354, 1412, 1470, 1528)
	result := matrix1.Multiply(matrix2)
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if result.Elements[i][j] != expected.Elements[i][j] {
				t.Error("Expected multiplied matrix, got", result.Elements[i][j])
			}
		}
	}
}
