package mapper

import (
	"NBodySim/internal/zmapper/approximator"
	"NBodySim/internal/zmapper/buffers"
	"image/color"
)

const DEPTH_NUM_WORKERS = 4

type ParallelDepthZmapper struct {
	DepthZmapper
	sync buffers.SyncBuffer
}

type ParallelDepthZmapperFabric struct {
	width, height int
	background    color.Color
}

func NewParallelDepthZmapperFabric(width, height int, background color.Color) *ParallelDepthZmapperFabric {
	return &ParallelDepthZmapperFabric{
		width:      width,
		height:     height,
		background: background,
	}
}

func (f *ParallelDepthZmapperFabric) CreateZmapper() Zmapper {
	return newParallelDepthZmapper(f.width, f.height, f.background, &buffers.DepthBufferInfFabric{})
}

func newParallelDepthZmapper(width, height int, background color.Color, df buffers.DepthBufferFabric) *ParallelDepthZmapper {
	par := ParallelDepthZmapper{}
	par.width = width
	par.height = height
	par.background = background
	par.sync = *buffers.NewSyncBuffer(width, height)
	par.dbuf = df.CreateDepthBuffer(width, height)
	return &par
}

func (zm *ParallelDepthZmapper) setPoint(x, y int, z float64, _ color.Color) {
	zm.sync.Lock(x, y)
	zm.dbuf.PutPoint(x, y, z)
	zm.sync.Unlock(x, y)
}

func (zm *ParallelDepthZmapper) DrawChannel(ch <-chan approximator.DiscreteFlatPoint) {
	for i := 0; i < DEPTH_NUM_WORKERS; i++ {
		go func() {
			for dp := range ch {
				zm.setPoint(dp.X, dp.Y, dp.Z, dp.Color)
				// fmt.Println(dp.X, dp.Y, dp.Z)
			}
		}()
	}
}

func (zm *ParallelDepthZmapper) SetPointDepth(p *approximator.DiscreteFlatPoint) {
	zm.sync.Lock(p.X, p.Y)
	p.Z, _ = zm.dbuf.GetDepth(p.X, p.Y)
	zm.sync.Unlock(p.X, p.Y)
}
