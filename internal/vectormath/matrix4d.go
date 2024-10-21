package vectormath

type Matrix4d struct {
	Elements [4][4]float64
}

func NewMatrix4d(
	m11, m12, m13, m14 float64,
	m21, m22, m23, m24 float64,
	m31, m32, m33, m34 float64,
	m41, m42, m43, m44 float64) *Matrix4d {
	return &Matrix4d{
		Elements: [4][4]float64{
			{m11, m12, m13, m14},
			{m21, m22, m23, m24},
			{m31, m32, m33, m34},
			{m41, m42, m43, m44},
		},
	}
}

func NewZeroMatrix4d() *Matrix4d {
	return &Matrix4d{
		Elements: [4][4]float64{},
	}
}

func NewEyeMatrix4d() *Matrix4d {
	return &Matrix4d{
		Elements: [4][4]float64{
			{1, 0, 0, 0},
			{0, 1, 0, 0},
			{0, 0, 1, 0},
			{0, 0, 0, 1},
		},
	}
}

func (m *Matrix4d) MultiplyScalar(scalar float64) {
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			m.Elements[i][j] *= scalar
		}
	}
}

func (m *Matrix4d) DivideScalar(scalar float64) {
	if scalar != 0 {
		for i := 0; i < 4; i++ {
			for j := 0; j < 4; j++ {
				m.Elements[i][j] /= scalar
			}
		}
	}
}

func (m *Matrix4d) AddScalar(scalar float64) {
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			m.Elements[i][j] += scalar
		}
	}
}

func (m *Matrix4d) SubtractScalar(scalar float64) {
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			m.Elements[i][j] -= scalar
		}
	}
}

func (m *Matrix4d) Multiply(other *Matrix4d) *Matrix4d {
	res := NewZeroMatrix4d()
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			for k := 0; k < 4; k++ {
				res.Elements[i][j] += m.Elements[i][k] * other.Elements[k][j]
			}
		}
	}
	return res
}

func (m *Matrix4d) Transpose() {
	for i := 0; i < 4; i++ {
		for j := i + 1; j < 4; j++ {
			m.Elements[i][j], m.Elements[j][i] = m.Elements[j][i], m.Elements[i][j]
		}
	}
}

// Multiplication when matrix is on the left
func (m *Matrix4d) LeftMultiply(v *HomoVector) *HomoVector {
	return NewHomoVector(
		m.Elements[0][0]*v.X+m.Elements[0][1]*v.Y+m.Elements[0][2]*v.Z+m.Elements[0][3]*v.W,
		m.Elements[1][0]*v.X+m.Elements[1][1]*v.Y+m.Elements[1][2]*v.Z+m.Elements[1][3]*v.W,
		m.Elements[2][0]*v.X+m.Elements[2][1]*v.Y+m.Elements[2][2]*v.Z+m.Elements[2][3]*v.W,
		m.Elements[3][0]*v.X+m.Elements[3][1]*v.Y+m.Elements[3][2]*v.Z+m.Elements[3][3]*v.W,
	)
}

// Multiplication when matrix is on the right
func (m *Matrix4d) RightMultiply(v *HomoVector) *HomoVector {
	return NewHomoVector(
		v.X*m.Elements[0][0]+v.Y*m.Elements[1][0]+v.Z*m.Elements[2][0]+v.W*m.Elements[3][0],
		v.X*m.Elements[0][1]+v.Y*m.Elements[1][1]+v.Z*m.Elements[2][1]+v.W*m.Elements[3][1],
		v.X*m.Elements[0][2]+v.Y*m.Elements[1][2]+v.Z*m.Elements[2][2]+v.W*m.Elements[3][2],
		v.X*m.Elements[0][3]+v.Y*m.Elements[1][3]+v.Z*m.Elements[2][3]+v.W*m.Elements[3][3],
	)
}
