package vector

import "testing"

func TestFromSlice(t *testing.T) {
	vec, err := FromSlice3d([]float64{1, 2, 3})
	if err != nil {
		t.Error(err)
	}
	if vec.X != 1 || vec.Y != 2 || vec.Z != 3 {
		t.Error("Expected (1, 2, 3), got", vec)
	}
}

func TestFromSliceInvalid(t *testing.T) {
	_, err := FromSlice3d([]float64{1, 2})
	if err == nil {
		t.Error("Expected error, got nil")
	}
}

func TestAddVectors(t *testing.T) {
	v1 := NewVector3d(1, 2, 3)
	v2 := NewVector3d(4, 5, 6)
	result := AddVectors(v1, v2)
	if result.X != 5 || result.Y != 7 || result.Z != 9 {
		t.Error("Expected (5, 7, 9), got", result)
	}
}

func TestSubtractVectors(t *testing.T) {
	v1 := NewVector3d(1, 2, 3)
	v2 := NewVector3d(4, 5, 6)
	result := SubtractVectors(v1, v2)
	if result.X != -3 || result.Y != -3 || result.Z != -3 {
		t.Error("Expected (-3, -3, -3), got", result)
	}
}

func TestMultiplyVectorScalar(t *testing.T) {
	v := NewVector3d(1, 2, 3)
	scalar := 2.0
	result := MultiplyVectorScalar(v, scalar)
	if result.X != 2 || result.Y != 4 || result.Z != 6 {
		t.Error("Expected (2, 4, 6), got", result)
	}
}

func TestDotProduct(t *testing.T) {
	v1 := NewVector3d(1, 2, 3)
	v2 := NewVector3d(4, 5, 6)
	result := DotProduct(v1, v2)
	if result != 32 {
		t.Error("Expected 32, got", result)
	}
}

func TestCrossProduct(t *testing.T) {
	v1 := NewVector3d(1, 2, 3)
	v2 := NewVector3d(4, 5, 6)
	result := CrossProduct(v1, v2)
	if result.X != -3 || result.Y != 6 || result.Z != -3 {
		t.Error("Expected (-3, 6, 3), got", result)
	}
}

func TestNormalize(t *testing.T) {
	v := NewVector3d(1, 2, 3)
	result := NormalizeVector(v)
	length := result.X*result.X + result.Y*result.Y + result.Z*result.Z
	if length != 1 {
		t.Error("Expected normalized vector length to be 1, got", length)
	}
}

func TestLength(t *testing.T) {
	v := NewVector3d(3, 4, 0)
	result := Length(v)
	if result != 5 {
		t.Error("Expected length to be 5, got", result)
	}
}

func TestCopy(t *testing.T) {
	v := NewVector3d(1, 2, 3)
	result := v.Copy()
	if v.X != result.X || v.Y != result.Y || v.Z != result.Z {
		t.Error("Expected copied vector to be the same, got", result)
	}
	if result == v {
		t.Error("Expected copied vector to be different from the original, got", result)
	}
}

func TestAdd(t *testing.T) {
	v1 := NewVector3d(1, 2, 3)
	v2 := NewVector3d(4, 5, 6)
	v1.Add(v2)
	if v1.X != 5 || v1.Y != 7 || v1.Z != 9 {
		t.Error("Expected (5, 7, 9), got", v1)
	}
}

func TestSubtract(t *testing.T) {
	v1 := NewVector3d(1, 2, 3)
	v2 := NewVector3d(4, 5, 6)
	v1.Subtract(v2)
	if v1.X != -3 || v1.Y != -3 || v1.Z != -3 {
		t.Error("Expected (-3, -3, -3), got", v1)
	}
}

func TestMultiplyScalar(t *testing.T) {
	v := NewVector3d(1, 2, 3)
	scalar := 2.0
	v.MultiplyScalar(scalar)
	if v.X != 2 || v.Y != 4 || v.Z != 6 {
		t.Error("Expected (2, 4, 6), got", v)
	}
}

func TestDot(t *testing.T) {
	v1 := NewVector3d(1, 2, 3)
	v2 := NewVector3d(4, 5, 6)
	result := v1.Dot(v2)
	if result != 32 {
		t.Error("Expected 32, got", result)
	}
}

func TestCross(t *testing.T) {
	v1 := NewVector3d(1, 2, 3)
	v2 := NewVector3d(4, 5, 6)
	result := v1.Cross(v2)
	if result.X != -3 || result.Y != 6 || result.Z != -3 {
		t.Error("Expected (-3, 6, 3), got", result)
	}
}

func TestNormalizeMethod(t *testing.T) {
	v := NewVector3d(1, 2, 3)
	v.Normalize()
	length := v.X*v.X + v.Y*v.Y + v.Z*v.Z
	if length != 1 {
		t.Error("Expected normalized vector length to be 1, got", length)
	}
}

func TestToSlice(t *testing.T) {
	v := NewVector3d(1, 2, 3)
	result := v.ToSlice()
	if result[0] != 1 || result[1] != 2 || result[2] != 3 {
		t.Error("Expected [1, 2, 3], got", result)
	}
}

func TestSquare(t *testing.T) {
	v := NewVector3d(1, 2, 3)
	result := v.Square()
	if result != 14 {
		t.Error("Expected 14, got", result)
	}
}

func TestToHomoVector(t *testing.T) {
	v := NewVector3d(1, 2, 3)
	result := v.ToHomoVector()
	if result.X != 1 || result.Y != 2 || result.Z != 3 || result.W != 1 {
		t.Error("Expected (1, 2, 3, 1), got", result)
	}
}
