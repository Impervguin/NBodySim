package mapper

import (
	"NBodySim/internal/object"
	"NBodySim/internal/zmapper/buffers"
	"image/color"
	"sync"
)

type ParallelPerPolygonZmapper struct {
	SimpleZmapper
	syncbuf buffers.SyncBuffer
}

type ParallelPerPolygonZmapperFabric struct{}

func NewParallelPerPolygonZmapperFabric() *ParallelPerPolygonZmapperFabric {
	return &ParallelPerPolygonZmapperFabric{}
}

func (f *ParallelPerPolygonZmapperFabric) CreateZmapper(width, height int, background color.Color) Zmapper {
	return newParallelPerPolygonZmapper(width, height, background, &buffers.DepthBufferInfFabric{})
}

func newParallelPerPolygonZmapper(width, height int, background color.Color, df buffers.DepthBufferFabric) *ParallelPerPolygonZmapper {
	return &ParallelPerPolygonZmapper{
		SimpleZmapper: *newSimpleZmapper(width, height, background, df),
		syncbuf:       *buffers.NewSyncBuffer(width, height),
	}
}

func (p *ParallelPerPolygonZmapper) setPoint(x, y int, z float64, color color.Color) {
	if (x < 0 || x >= p.width) || (y < 0 || y >= p.height) {
		return
	}
	p.syncbuf.Lock(x, y)
	p.SimpleZmapper.setPoint(x, y, z, color)
	p.syncbuf.Unlock(x, y)
}

func (p *ParallelPerPolygonZmapper) VisitPolygonObject(po *object.PolygonObject) {
	w := sync.WaitGroup{}
	for _, polygon := range po.GetPolygons() {
		w.Add(1)
		go func() {
			defer w.Done()
			p.processPolygon(polygon)
		}()
	}
	w.Wait()
}

func (zm *ParallelPerPolygonZmapper) processPolygon(p *object.Polygon) {
	ch := zm.polygonGenerator(p)
	for zmp := range ch {
		zm.setPoint(zmp.X, zmp.Y, zmp.Z, zmp.Color)
	}
}
