package buffers

import "NBodySim/internal/mathutils/normal"

type NormalBuffer struct {
	buf        [][]normal.Normal
	width      int
	height     int
	background normal.Normal
}

func NewNormalBuffer(width, height int, background normal.Normal) *NormalBuffer {
	buf := make([][]normal.Normal, height)
	for i := 0; i < height; i++ {
		buf[i] = make([]normal.Normal, width)
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

func (n *NormalBuffer) PutPoint(x, y int, v normal.Normal) (bool, error) {
	if (x < 0 || x >= n.width) || (y < 0 || y >= n.height) {
		return false, ErrOutOfBounds
	}
	n.buf[y][x] = v
	return true, nil
}

func (n *NormalBuffer) GetPoint(x, y int) (normal.Normal, error) {
	if (x < 0 || x >= n.width) || (y < 0 || y >= n.height) {
		return normal.Normal{}, ErrOutOfBounds
	}
	return n.buf[y][x], nil
}
