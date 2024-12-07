package vector

import (
	"fmt"
	"math"
)

type Vector3d struct {
	X, Y, Z float64
}

func NewVector3d(x, y, z float64) *Vector3d {
	return &Vector3d{x, y, z}
}

func FromSlice3d(vec []float64) (*Vector3d, error) {
	if len(vec) != 3 {
		return nil, fmt.Errorf("slice len must be 3")
	}
	return NewVector3d(vec[0], vec[1], vec[2]), nil
}

func AddVectors(v1, v2 *Vector3d) *Vector3d {
	return NewVector3d(v1.X+v2.X, v1.Y+v2.Y, v1.Z+v2.Z)
}

func SubtractVectors(v1, v2 *Vector3d) *Vector3d {
	return NewVector3d(v1.X-v2.X, v1.Y-v2.Y, v1.Z-v2.Z)
}

func MultiplyVectorScalar(v *Vector3d, scalar float64) *Vector3d {
	return NewVector3d(v.X*scalar, v.Y*scalar, v.Z*scalar)
}

func SquareVector(v *Vector3d) float64 {
	return v.X*v.X + v.Y*v.Y + v.Z*v.Z
}

func DotProduct(v1, v2 *Vector3d) float64 {
	return v1.X*v2.X + v1.Y*v2.Y + v1.Z*v2.Z
}

func CrossProduct(v1, v2 *Vector3d) *Vector3d {
	return NewVector3d(
		v1.Y*v2.Z-v1.Z*v2.Y,
		v1.Z*v2.X-v1.X*v2.Z,
		v1.X*v2.Y-v1.Y*v2.X,
	)
}

func Length(v *Vector3d) float64 {
	return math.Sqrt(DotProduct(v, v))
}

func NormalizeVector(v *Vector3d) *Vector3d {
	length := Length(v)
	if length != 0 {
		v.X /= length
		v.Y /= length
		v.Z /= length
	}
	return v
}

func (v *Vector3d) Copy() *Vector3d {
	return NewVector3d(v.X, v.Y, v.Z)
}

func (v *Vector3d) Add(v1 *Vector3d) {
	v.X += v1.X
	v.Y += v1.Y
	v.Z += v1.Z
}

func (v *Vector3d) Subtract(v1 *Vector3d) {
	v.X -= v1.X
	v.Y -= v1.Y
	v.Z -= v1.Z
}

func (v *Vector3d) MultiplyScalar(scalar float64) {
	v.X *= scalar
	v.Y *= scalar
	v.Z *= scalar
}

func (v *Vector3d) Normalize() {
	length := Length(v)
	if length != 0 {
		v.X /= length
		v.Y /= length
		v.Z /= length
	}
}

func (v *Vector3d) Dot(v1 *Vector3d) float64 {
	return v.X*v1.X + v.Y*v1.Y + v.Z*v1.Z
}

func (v *Vector3d) Cross(v1 *Vector3d) *Vector3d {
	return NewVector3d(
		v.Y*v1.Z-v.Z*v1.Y,
		v.Z*v1.X-v.X*v1.Z,
		v.X*v1.Y-v.Y*v1.X,
	)
}

func (v *Vector3d) ToSlice() []float64 {
	return []float64{v.X, v.Y, v.Z}
}

func (v *Vector3d) Square() float64 {
	return v.X*v.X + v.Y*v.Y + v.Z*v.Z
}

func (v *Vector3d) ToHomoVector() *HomoVector {
	return NewHomoVector(v.X, v.Y, v.Z, 1)
}
