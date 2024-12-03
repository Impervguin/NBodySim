package vector

import "testing"

func TestToVector3d(t *testing.T) {
	vec := NewHomoVector(1, 2, 3, 0.5)
	result := vec.ToVector3d()
	if result.X != 2 || result.Y != 4 || result.Z != 6 {
		t.Error("Expected (2, 4, 6), got", result)
	}
}

func TestCopy4d(t *testing.T) {
	vec := NewHomoVector(1, 2, 3, 0.5)
	copy := vec.Copy()
	if vec.X != copy.X || vec.Y != copy.Y || vec.Z != copy.Z || vec.W != copy.W {
		t.Error("Expected same vector, got different copy")
	}
	if vec == copy {
		t.Error("Expected copied vector to be different from the original, got same reference")
	}
}

func TestAdd4d(t *testing.T) {
	v1 := NewHomoVector(1, 2, 3, 0.5)
	v2 := NewHomoVector(4, 5, 6, 0.5)
	result := AddHomoVectors(v1, v2)
	if result.X != 5 || result.Y != 7 || result.Z != 9 || result.W != 0.5 {
		t.Error("Expected (5, 7, 9, .5), got", result)
	}
}

func TestAddMethod4d(t *testing.T) {
	v1 := NewHomoVector(1, 2, 3, 0.5)
	v2 := NewHomoVector(4, 5, 6, 0.5)
	v1.Add(v2)
	if v1.X != 5 || v1.Y != 7 || v1.Z != 9 || v1.W != 0.5 {
		t.Error("Expected (5, 7, 9, 0.5), got", v1)
	}
}

func TestSubtract4d(t *testing.T) {
	v1 := NewHomoVector(1, 2, 3, 0.5)
	v2 := NewHomoVector(4, 5, 6, 0.5)
	result := SubtractHomoVectors(v1, v2)
	if result.X != -3 || result.Y != -3 || result.Z != -3 || result.W != 0.5 {
		t.Error("Expected (-3, -3, -3, 0.5), got", result)
	}
}

func TestSubtractMethod4d(t *testing.T) {
	v1 := NewHomoVector(1, 2, 3, 0.5)
	v2 := NewHomoVector(4, 5, 6, 0.5)
	v1.Subtract(v2)
	if v1.X != -3 || v1.Y != -3 || v1.Z != -3 || v1.W != 0.5 {
		t.Error("Expected (-3, -3, -3, 0.5), got", v1)
	}
}

func TestMultiplyScalar4d(t *testing.T) {
	v := NewHomoVector(1, 2, 3, 0.5)
	v.MultiplyScalar(2)
	if v.X != 2 || v.Y != 4 || v.Z != 6 || v.W != 0.5 {
		t.Error("Expected (2, 4, 6, 0.5), got", v)
	}
}

func TestMultiplyScalarFunc(t *testing.T) {
	v := NewHomoVector(1, 2, 3, 0.5)
	result := MultiplyHomoVectorScalar(v, 2)
	if result.X != 2 || result.Y != 4 || result.Z != 6 || result.W != 0.5 {
		t.Error("Expected (2, 4, 6, 0.5), got", result)
	}
}

func TestDivideScalar4d(t *testing.T) {
	v := NewHomoVector(1, 2, 3, 0.5)
	v.DivideScalar(2)
	if v.X != 0.5 || v.Y != 1 || v.Z != 1.5 || v.W != 0.5 {
		t.Error("Expected (0.5, 1, 1.5, 0.5), got", v)
	}
}

func TestDivideScalar4dZero(t *testing.T) {
	v := NewHomoVector(1, 2, 3, 0.5)
	v.DivideScalar(0)
	if v.X != 1 || v.Y != 2 || v.Z != 3 || v.W != 0.5 {
		t.Error("Expected (1, 2, 3, 0.5), got", v)
	}
}

func TestDivideScalarFunc(t *testing.T) {
	v := NewHomoVector(1, 2, 3, 0.5)
	result := DivideHomoVectorScalar(v, 2)
	if result.X != 0.5 || result.Y != 1 || result.Z != 1.5 || result.W != 0.5 {
		t.Error("Expected (0.5, 1, 1.5, 0.5), got", result)
	}
}

func TestDivideScalarFuncZero(t *testing.T) {
	v := NewHomoVector(1, 2, 3, 0.5)
	result := DivideHomoVectorScalar(v, 0)
	if result != nil {
		t.Error("Expected nil, got", result)
	}
}
