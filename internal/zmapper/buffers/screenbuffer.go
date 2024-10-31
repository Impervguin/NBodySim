package buffers

import (
	"image/color"
)

type ScreenBuffer struct {
	buf        [][]color.Color
	width      int
	height     int
	background color.Color
}

func newScreenBuffer(width, height int) *ScreenBuffer {
	buf := make([][]color.Color, height)
	for i := 0; i < height; i++ {
		buf[i] = make([]color.Color, width)
	}
	return &ScreenBuffer{
		buf:    buf,
		width:  width,
		height: height,
	}
}

func NewScreenBuffer(width, height int, background color.Color) *ScreenBuffer {
	dbuf := newScreenBuffer(width, height)
	dbuf.background = background
	dbuf.Reset()
	return dbuf
}

func (s *ScreenBuffer) Reset() {
	for i := range s.buf {
		for j := range s.buf[i] {
			s.buf[i][j] = s.background
		}
	}
}
func (s *ScreenBuffer) PutPoint(x, y int, color color.Color) error {
	if (x < 0 || x >= s.width) || (y < 0 || y >= s.height) {
		return ErrOutOfBounds
	}
	s.buf[y][x] = color
	return nil
}

func (s *ScreenBuffer) GetPoint(x, y int) color.Color {
	// if (x < 0 || x >= s.width) || (y < 0 || y >= s.height) {
	// 	return s.background
	// }
	return s.buf[y][x]
}
