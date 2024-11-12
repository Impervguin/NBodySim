package buffers

import (
	"math"
)

type DepthBuffer interface {
	Reset()
	PutPoint(x, y int, z float64) (bool, error)
	GetDepth(x, y int) (float64, error)
}

type DepthBufferNull struct {
	buf    [][]float64
	width  int
	height int
}

type DepthBufferInf struct {
	buf    [][]float64
	width  int
	height int
}

func newDepthBuffer(width, height int) [][]float64 {
	buf := make([][]float64, height)
	for i := 0; i < width; i++ {
		buf[i] = make([]float64, width)
	}
	return buf
}

func NewInfDepthBuffer(width, height int) *DepthBufferInf {
	buf := newDepthBuffer(width, height)
	dbuf := &DepthBufferInf{
		buf:    buf,
		width:  width,
		height: height,
	}
	dbuf.Reset()
	return dbuf
}

func NewNullDepthBuffer(width, height int) *DepthBufferNull {
	buf := newDepthBuffer(width, height)
	dbuf := &DepthBufferNull{
		buf:    buf,
		width:  width,
		height: height,
	}
	dbuf.Reset()
	return dbuf
}

func (d *DepthBufferInf) Reset() {
	for i := range d.buf {
		for j := range d.buf[i] {
			d.buf[i][j] = math.Inf(1)
		}
	}
}

func (d *DepthBufferNull) Reset() {
	for _, el := range d.buf {
		for i := range el {
			el[i] = 0
		}
	}
}

func (d *DepthBufferInf) PutPoint(x, y int, z float64) (bool, error) {
	if (x < 0 || x >= d.width) || (y < 0 || y >= d.height) {
		return false, ErrOutOfBounds
	}
	if z < d.buf[y][x] {
		d.buf[y][x] = z
		return true, nil
	}
	return false, nil
}

func (d *DepthBufferNull) PutPoint(x, y int, z float64) (bool, error) {
	if (x < 0 || x >= d.width) || (y < 0 || y >= d.height) {
		return false, ErrOutOfBounds
	}
	if z > d.buf[y][x] {
		d.buf[y][x] = z
		return true, nil
	}
	return false, nil
}

func (d *DepthBufferInf) GetDepth(x, y int) (float64, error) {
	if (x < 0 || x >= d.width) || (y < 0 || y >= d.height) {
		return 0, ErrOutOfBounds
	}
	return d.buf[y][x], nil
}

func (d *DepthBufferNull) GetDepth(x, y int) (float64, error) {
	if (x < 0 || x >= d.width) || (y < 0 || y >= d.height) {
		return 0, ErrOutOfBounds
	}
	return d.buf[y][x], nil
}

type DepthBufferFabric interface {
	CreateDepthBuffer(width, height int) DepthBuffer
}

type DepthBufferInfFabric struct{}
type DepthBufferNullFabric struct{}

func (f *DepthBufferInfFabric) CreateDepthBuffer(width, height int) DepthBuffer {
	return NewInfDepthBuffer(width, height)
}

func (f *DepthBufferNullFabric) CreateDepthBuffer(width, height int) DepthBuffer {
	return NewNullDepthBuffer(width, height)
}
