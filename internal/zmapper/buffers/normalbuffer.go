package buffers

import "NBodySim/internal/mathutils/vector"

type NormalBuffer struct {
	buf        [][]vector.Vector3d
	width      int
	height     int
	background vector.Vector3d
}

func NewNormalBuffer(width, height int, background vector.Vector3d) *NormalBuffer {
	buf := make([][]vector.Vector3d, height)
	for i := 0; i < height; i++ {
		buf[i] = make([]vector.Vector3d, width)
		for j := 0; j < width; j++ {
			buf[i][j] = background
		}
	}
	return &NormalBuffer{buf: buf, width: width, height: height, background: background}
}

func (n *NormalBuffer) Reset() {
	for i := range n.buf {
		for j := range n.buf[i] {
			n.buf[i][j] = n.background
		}
	}
}

func (n *NormalBuffer) PutPoint(x, y int, v vector.Vector3d) (bool, error) {
	if (x < 0 || x >= n.width) || (y < 0 || y >= n.height) {
		return false, ErrOutOfBounds
	}
	n.buf[y][x] = v
	return true, nil
}

func (n *NormalBuffer) GetPoint(x, y int) (vector.Vector3d, error) {
	if (x < 0 || x >= n.width) || (y < 0 || y >= n.height) {
		return vector.Vector3d{}, ErrOutOfBounds
	}
	return n.buf[y][x], nil
}
