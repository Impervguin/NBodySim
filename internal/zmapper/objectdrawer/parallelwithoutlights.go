package objectdrawer

import (
	"NBodySim/internal/mathutils/vector"
	"NBodySim/internal/object"
	"NBodySim/internal/zmapper/approximator"
	"NBodySim/internal/zmapper/mapper"
	"image"
	"sync"
)

type ParallelWithoutLightsDrawer struct {
	zmapper      mapper.NormalZmapper
	approxf      approximator.DiscreteNormalApproximatorFabric
	approx       approximator.DiscreteNormalApproximator
	cam          *object.Camera
	mut          sync.Mutex
	visitCounter int
	pchan        chan approximator.DiscreteNormalPoint
}

type ParallelWithoutLightsDrawerFabric struct {
	zf mapper.NormalZmapperFabric
	af approximator.DiscreteNormalApproximatorFabric
}

func NewParallelWithoutLightsDrawerFabric(zf mapper.NormalZmapperFabric, af approximator.DiscreteNormalApproximatorFabric) *ParallelWithoutLightsDrawerFabric {
	return &ParallelWithoutLightsDrawerFabric{
		zf: zf,
		af: af,
	}
}

func (pd *ParallelWithoutLightsDrawerFabric) CreateObjectDrawerWithoutLights() ObjectDrawerWithoutLights {
	return newParallelWithoutLightsDrawer(pd.zf, pd.af)
}

func newParallelWithoutLightsDrawer(zf mapper.NormalZmapperFabric, af approximator.DiscreteNormalApproximatorFabric) *ParallelWithoutLightsDrawer {
	return &ParallelWithoutLightsDrawer{
		zmapper:      zf.CreateNormalZmapper(),
		approxf:      af,
		approx:       nil,
		cam:          object.NewCamera(*vector.NewVector3d(0, 0, 0), *vector.NewVector3d(0, 0, 1), *vector.NewVector3d(0, 1, 0), 1, 1, 1),
		pchan:        nil,
		mut:          sync.Mutex{},
		visitCounter: 0,
	}
}

func (pd *ParallelWithoutLightsDrawer) startDrawing() {
	pd.mut.Lock()
	if pd.visitCounter == 0 {
		pd.pchan = make(chan approximator.DiscreteNormalPoint)
		pd.approx = pd.approxf.CreateDiscreteApproximator()
		go pd.zmapper.DrawChannel(pd.pchan)
	}
	pd.visitCounter++
	pd.mut.Unlock()
}

func (pd *ParallelWithoutLightsDrawer) stopDrawing() {
	pd.mut.Lock()
	pd.visitCounter--
	if pd.visitCounter == 0 {
		close(pd.pchan)
		pd.approx = nil
	}
	pd.mut.Unlock()
}

func (pd *ParallelWithoutLightsDrawer) SetZmapper(zm mapper.NormalZmapper) {
	pd.zmapper = zm
}

func (pd *ParallelWithoutLightsDrawer) VisitCamera(cam *object.Camera) {
	pd.cam = cam
}

func (pd *ParallelWithoutLightsDrawer) VisitPolygonObject(po *object.PolygonObject) {
	pd.startDrawing()
	for _, v := range po.GetPolygons() {
		pd.approx.ApproximatePolygon(v, pd.pchan)
	}
	pd.stopDrawing()
}

func (pd *ParallelWithoutLightsDrawer) VisitObjectPool(light *object.ObjectPool) {
	pd.startDrawing()
	wg := &sync.WaitGroup{}
	for _, obj := range light.GetObjects() {
		wg.Add(1)
		go func() {
			obj.Accept(pd)
			wg.Done()
		}()
	}
	wg.Wait()
	pd.stopDrawing()
}

func (pd *ParallelWithoutLightsDrawer) GetImage() image.Image {
	return pd.zmapper
}

func (pd *ParallelWithoutLightsDrawer) GetZmapper() mapper.NormalZmapper {
	return pd.zmapper
}

func (pd *ParallelWithoutLightsDrawer) ResetImage() {
	pd.zmapper.Reset()
}

func (pd *ParallelWithoutLightsDrawer) GetWidth() int {
	return pd.zmapper.Bounds().Dx()
}

func (pd *ParallelWithoutLightsDrawer) GetHeight() int {
	return pd.zmapper.Bounds().Dy()
}
