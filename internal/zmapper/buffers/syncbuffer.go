package buffers

import "sync"

type SyncBuffer struct {
	buf    [][]sync.Mutex
	width  int
	height int
}

func NewSyncBuffer(width, height int) *SyncBuffer {
	buf := make([][]sync.Mutex, height)
	for i := 0; i < height; i++ {
		buf[i] = make([]sync.Mutex, width)
	}
	return &SyncBuffer{
		buf:    buf,
		width:  width,
		height: height,
	}
}

func (s *SyncBuffer) Lock(x, y int) {
	if (x < 0 || x >= s.width) || (y < 0 || y >= s.height) {
		return
	}
	s.buf[y][x].Lock()
}

func (s *SyncBuffer) Unlock(x, y int) {
	if (x < 0 || x >= s.width) || (y < 0 || y >= s.height) {
		return
	}
	s.buf[y][x].Unlock()
}
