package vectormath

type HomoVector struct {
	X, Y, Z, W float64
}

func NewHomoVector(x, y, z, w float64) *HomoVector {
	return &HomoVector{x, y, z, w}
}

func (v *HomoVector) Copy() *HomoVector {
	return &HomoVector{v.X, v.Y, v.Z, v.W}
}

func (v *HomoVector) ToVector3d() *Vector3d {
	return NewVector3d(v.X/v.W, v.Y/v.W, v.Z/v.W)
}

func (v *HomoVector) Add(v1 *HomoVector) {
	v.X += v1.X / v1.W * v.W
	v.Y += v1.Y / v1.W * v.W
	v.Z += v1.Z / v1.W * v.W
}

func (v *HomoVector) Subtract(v1 *HomoVector) {
	v.X -= v1.X / v1.W * v.W
	v.Y -= v1.Y / v1.W * v.W
	v.Z -= v1.Z / v1.W * v.W
}

func (v *HomoVector) MultiplyScalar(scalar float64) {
	v.X *= scalar
	v.Y *= scalar
	v.Z *= scalar
}

func (v *HomoVector) DivideScalar(scalar float64) {
	if scalar != 0 {
		v.X /= scalar
		v.Y /= scalar
		v.Z /= scalar
	}
}

func AddHomoVectors(v1 *HomoVector, v2 *HomoVector) *HomoVector {
	return &HomoVector{
		X: v1.X + v2.X/v2.W*v1.W,
		Y: v1.Y + v2.Y/v2.W*v1.W,
		Z: v1.Z + v2.Z/v2.W*v1.W,
		W: v1.W,
	}
}

func SubtractHomoVectors(v1 *HomoVector, v2 *HomoVector) *HomoVector {
	return &HomoVector{
		X: v1.X - v2.X/v2.W*v1.W,
		Y: v1.Y - v2.Y/v2.W*v1.W,
		Z: v1.Z - v2.Z/v2.W*v1.W,
		W: v1.W,
	}
}

func MultiplyHomoVectorScalar(v *HomoVector, scalar float64) *HomoVector {
	return &HomoVector{
		X: v.X * scalar,
		Y: v.Y * scalar,
		Z: v.Z * scalar,
		W: v.W,
	}
}

func DivideHomoVectorScalar(v *HomoVector, scalar float64) *HomoVector {
	if scalar != 0 {
		return &HomoVector{
			X: v.X / scalar,
			Y: v.Y / scalar,
			Z: v.Z / scalar,
			W: v.W,
		}
	}
	return nil
}
