package buffers

import (
	"NBodySim/internal/object"
)

type NormalBuffer struct {
	buf        [][]object.PolygonNormal
	width      int
	height     int
	background object.PolygonNormal
}

func NewNormalBuffer(width, height int, background object.PolygonNormal) *NormalBuffer {
	buf := make([][]object.PolygonNormal, height)
	for i := 0; i < height; i++ {
		buf[i] = make([]object.PolygonNormal, width)
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

func (n *NormalBuffer) PutPoint(x, y int, v object.PolygonNormal) (bool, error) {
	if (x < 0 || x >= n.width) || (y < 0 || y >= n.height) {
		return false, ErrOutOfBounds
	}
	n.buf[y][x] = v
	return true, nil
}

func (n *NormalBuffer) GetPoint(x, y int) (object.PolygonNormal, error) {
	if (x < 0 || x >= n.width) || (y < 0 || y >= n.height) {
		return object.PolygonNormal{}, ErrOutOfBounds
	}
	return n.buf[y][x], nil
}
