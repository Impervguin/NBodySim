package mapper

import (
	"NBodySim/internal/zmapper/approximator"
	"NBodySim/internal/zmapper/buffers"
	"image/color"
)

const NUM_WORKERS = 4

type ParallelZmapper struct {
	SimpleZmapper
	sync buffers.SyncBuffer
}

type ParallelZmapperFabric struct {
	width, height int
	background    color.Color
}

func NewParallelZmapperFabric(width, height int, background color.Color) *ParallelZmapperFabric {
	return &ParallelZmapperFabric{
		width:      width,
		height:     height,
		background: background,
	}
}

func (f *ParallelZmapperFabric) CreateZmapper() Zmapper {
	return newParallelZmapper(f.width, f.height, f.background, &buffers.DepthBufferInfFabric{})
}

func newParallelZmapper(width, height int, background color.Color, df buffers.DepthBufferFabric) *ParallelZmapper {
	par := ParallelZmapper{}
	par.width = width
	par.height = height
	par.sbuf = *buffers.NewScreenBuffer(width, height, background)
	par.sync = *buffers.NewSyncBuffer(width, height)
	par.dbuf = df.CreateDepthBuffer(width, height)
	return &par
}

func (zm *ParallelZmapper) setPoint(x, y int, z float64, color color.Color) {
	zm.sync.Lock(x, y)
	ok, _ := zm.dbuf.PutPoint(x, y, z)
	if ok {
		zm.sbuf.PutPoint(x, y, color)
	}
	zm.sync.Unlock(x, y)
}

func (zm *ParallelZmapper) DrawChannel(ch <-chan approximator.DiscreteFlatPoint) {
	for i := 0; i < NUM_WORKERS; i++ {
		go func() {
			for dp := range ch {
				zm.setPoint(dp.X, dp.Y, dp.Z, dp.Color)
			}
		}()
	}
}
