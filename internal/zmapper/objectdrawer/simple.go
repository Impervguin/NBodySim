package objectdrawer

import (
	"NBodySim/internal/mathutils/vector"
	"NBodySim/internal/object"
	"NBodySim/internal/zmapper/approximator"
	"NBodySim/internal/zmapper/approximator/colorist"
	"NBodySim/internal/zmapper/mapper"
	"image"
	"sync"
)

type SimpleObjectDrawer struct {
	lights       []object.Light
	zmapper      mapper.Zmapper
	approxf      approximator.DiscreteApproximatorFabric
	approx       approximator.DiscreteApproximator
	view         vector.Vector3d
	cam          *object.Camera
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
		approxf:      af,
		approx:       nil,
		view:         vector.Vector3d{0, 0, 0},
		mut:          sync.Mutex{},
		visitCounter: 0,
		pchan:        nil,
	}
}

func (sd *SimpleObjectDrawer) startDrawing() {
	sd.mut.Lock()
	if sd.visitCounter == 0 {
		sd.pchan = make(chan approximator.DiscreteFlatPoint)
		sd.approx = sd.approxf.CreateDiscreteApproximator(sd.view)
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
		sd.approx = nil
	}
	sd.mut.Unlock()
}

func (sd *SimpleObjectDrawer) SetView(view vector.Vector3d) {
	sd.view = view
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

func (sd *SimpleObjectDrawer) GetColorist(view vector.Vector3d) colorist.Colorist {
	return sd.approxf.GetColorist(view)
}
