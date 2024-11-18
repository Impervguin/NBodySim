package mapper

import (
	"NBodySim/internal/mathutils/vector"
	"NBodySim/internal/object"
	"NBodySim/internal/transform"
	"NBodySim/internal/zmapper/approximator"
	"NBodySim/internal/zmapper/buffers"
	"image"
	"image/color"
	"sync"
)

const NUM_WORKERS_ParallelZmapperWithNormals = 4
const NUM_LIGHT_WORKERS = 4

type ParallelZmapperWithNormals struct {
	dbuf   buffers.DepthBuffer
	sbuf   buffers.ScreenBuffer
	nbuf   buffers.NormalBuffer
	sync   buffers.SyncBuffer
	width  int
	height int
}

type ParallelZmapperWithNormalsFabric struct {
	width, height int
	background    color.Color
	df            buffers.DepthBufferFabric
}

func NewParallelZmapperWithNormalsFabric(width, height int, background color.Color, df buffers.DepthBufferFabric) *ParallelZmapperWithNormalsFabric {
	return &ParallelZmapperWithNormalsFabric{
		width:      width,
		height:     height,
		background: background,
		df:         df,
	}
}

func (f *ParallelZmapperWithNormalsFabric) CreateNormalZmapper() NormalZmapper {
	return newParallelZmapperWithNormals(f.width, f.height, f.background, f.df)
}

func newParallelZmapperWithNormals(width, height int, background color.Color, df buffers.DepthBufferFabric) *ParallelZmapperWithNormals {
	par := ParallelZmapperWithNormals{
		width:  width,
		height: height,
		sbuf:   *buffers.NewScreenBuffer(width, height, background),
		nbuf:   *buffers.NewNormalBuffer(width, height, *vector.NewVector3d(0, 0, 0)),
		dbuf:   df.CreateDepthBuffer(width, height),
		sync:   *buffers.NewSyncBuffer(width, height),
	}
	return &par
}

func (pz *ParallelZmapperWithNormals) setPoint(p *approximator.DiscreteNormalPoint) {
	pz.sync.Lock(p.X, p.Y)
	put, _ := pz.dbuf.PutPoint(p.X, p.Y, p.Z)
	if put {
		pz.sbuf.PutPoint(p.X, p.Y, p.Color)
		pz.nbuf.PutPoint(p.X, p.Y, p.Normal)
	}
	pz.sync.Unlock(p.X, p.Y)
}

func (pz *ParallelZmapperWithNormals) updatePoint(p *approximator.DiscreteNormalPoint) {
	pz.sync.Lock(p.X, p.Y)
	pz.dbuf.PutPoint(p.X, p.Y, p.Z)
	pz.sbuf.PutPoint(p.X, p.Y, p.Color)
	pz.nbuf.PutPoint(p.X, p.Y, p.Normal)
	pz.sync.Unlock(p.X, p.Y)
}

func (pz *ParallelZmapperWithNormals) DrawChannel(ch <-chan approximator.DiscreteNormalPoint) {
	wg := sync.WaitGroup{}
	for i := 0; i < NUM_WORKERS_ParallelZmapperWithNormals; i++ {
		wg.Add(1)
		go func() {
			for dp := range ch {
				pz.setPoint(&dp)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}

func (pz *ParallelZmapperWithNormals) GetScreenFunction() buffers.ScreenFunction {
	return func(x, y, w, h int) color.Color {
		return pz.sbuf.GetPoint(x, y)
	}
}

func (pz *ParallelZmapperWithNormals) Reset() {
	pz.sbuf.Reset()
	pz.nbuf.Reset()
	pz.dbuf.Reset()
}

func (pz *ParallelZmapperWithNormals) Bounds() image.Rectangle {
	return image.Rect(0, 0, pz.width, pz.height)
}

func (pz *ParallelZmapperWithNormals) At(x, y int) color.Color {
	return pz.sbuf.GetPoint(x, y)
}

func (pz *ParallelZmapperWithNormals) ColorModel() color.Model {
	return color.RGBA64Model
}

func (pz *ParallelZmapperWithNormals) GetPoint(x, y int) *approximator.DiscreteNormalPoint {
	dp := approximator.DiscreteNormalPoint{X: x, Y: y}
	dp.Color = pz.sbuf.GetPoint(x, y)
	dp.Normal, _ = pz.nbuf.GetPoint(x, y)
	dp.Z, _ = pz.dbuf.GetDepth(x, y)
	return &dp
}

func (pz *ParallelZmapperWithNormals) ApplyLight(lp *object.LightPool, tolights transform.TransformAction) {
	yPerWorker := pz.height / NUM_LIGHT_WORKERS
	yRemainigs := pz.height % NUM_LIGHT_WORKERS
	// fmt.Println(yPerWorker)
	wg := sync.WaitGroup{}
	for i := 0; i < NUM_LIGHT_WORKERS; i++ {
		wg.Add(1)
		go func(workerID int) {
			yStart := workerID * yPerWorker
			yEnd := yStart + yPerWorker
			if workerID == NUM_LIGHT_WORKERS-1 {
				yEnd += yRemainigs
			}
			for y := yStart; y < yEnd; y++ {
				for x := 0; x < pz.width; x++ {
					if ok, _ := pz.dbuf.Empty(x, y); ok {
						continue
					}
					dp := pz.GetPoint(x, y)
					vec := vector.NewVector3d(float64(dp.X), float64(dp.Y), dp.Z)
					toVec := transform.NewMoveAction(vec)
					tolights.ApplyToVector(vec)
					fromVec := transform.NewMoveAction(vector.MultiplyVectorScalar(vec, -1))
					normal := dp.Normal
					toVec.ApplyToVector(&normal)
					tolights.ApplyToVector(&normal)
					fromVec.ApplyToVector(&normal)
					normal.Normalize()
					view := vector.NewVector3d(float64(pz.width)/2, float64(pz.height)/2, 0)
					tolights.ApplyToVector(view)
					dp.Color = lp.CalculateLight(*vec, *view, normal, dp.Color)
					// r, _, _, _ := dp.Color.RGBA()
					// if r < 1000 {
					// 	fmt.Println(vec, normal)
					// }
					pz.updatePoint(dp)
				}
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
}
