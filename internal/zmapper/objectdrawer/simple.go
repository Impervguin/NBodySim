package objectdrawer

import (
	"NBodySim/internal/object"
	"NBodySim/internal/zmapper/approximator"
	"NBodySim/internal/zmapper/mapper"
	"image"
	"sync"
)

type SimpleObjectDrawer struct {
	lights       []object.Light
	zmapper      mapper.Zmapper
	approx       approximator.DiscreteApproximator
	mut          sync.Mutex
	visitCounter int
	pchan        chan approximator.DiscreteFlatPoint
}

type SimpleObjectDrawerFabric struct {
	zf mapper.ZmapperFabric
	af approximator.DiscreteApproximatorFabric
}

func NewSimpleObjectDrawerFabric(zf mapper.ZmapperFabric, af approximator.DiscreteApproximatorFabric) *SimpleObjectDrawerFabric {
	return &SimpleObjectDrawerFabric{
		zf: zf,
		af: af,
	}
}

func (f *SimpleObjectDrawerFabric) CreateObjectDrawer() ObjectDrawer {
	return newSimpleObjectDrawer(f.zf, f.af)
}

func newSimpleObjectDrawer(zf mapper.ZmapperFabric, af approximator.DiscreteApproximatorFabric) *SimpleObjectDrawer {
	return &SimpleObjectDrawer{
		lights:       []object.Light{},
		zmapper:      zf.CreateZmapper(),
		approx:       af.CreateDiscreteApproximator(),
		mut:          sync.Mutex{},
		visitCounter: 0,
		pchan:        nil,
	}
}

func (sd *SimpleObjectDrawer) startDrawing() {
	sd.mut.Lock()
	if sd.visitCounter == 0 {
		sd.pchan = make(chan approximator.DiscreteFlatPoint)
		go sd.zmapper.DrawChannel(sd.pchan)
	}
	sd.visitCounter++
	sd.mut.Unlock()
}

func (sd *SimpleObjectDrawer) stopDrawing() {
	sd.mut.Lock()
	sd.visitCounter--
	if sd.visitCounter == 0 {
		close(sd.pchan)
	}
	sd.mut.Unlock()
}

func (sd *SimpleObjectDrawer) VisitObjectPool(op *object.ObjectPool) {
	sd.startDrawing()
	for _, obj := range op.GetObjects() {
		obj.Accept(sd)
	}
	sd.stopDrawing()
}

func (sd *SimpleObjectDrawer) VisitPolygonObject(po *object.PolygonObject) {
	sd.startDrawing()
	for _, v := range po.GetPolygons() {
		sd.approx.ApproximatePolygon(v, sd.pchan)
	}
	sd.stopDrawing()
}

func (sd *SimpleObjectDrawer) VisitCamera(cam *object.Camera) {
	// Nothing to do here
}

func (sd *SimpleObjectDrawer) VisitPointLight(light *object.PointLight) {
	sd.lights = append(sd.lights, light)
}

func (sd *SimpleObjectDrawer) GetImage() image.Image {
	return sd.zmapper
}

func (sd *SimpleObjectDrawer) ResetImage() {
	sd.zmapper.Reset()
	sd.lights = make([]object.Light, 0)
}

func (sd *SimpleObjectDrawer) GetWidth() int {
	return sd.zmapper.Bounds().Dx()
}

func (sd *SimpleObjectDrawer) GetHeight() int {
	return sd.zmapper.Bounds().Dy()
}
