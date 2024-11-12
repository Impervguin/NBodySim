package shadowdrawer

import (
	"NBodySim/internal/mathutils/vector"
	"NBodySim/internal/object"
	"NBodySim/internal/transform"
	"NBodySim/internal/zmapper/approximator"
	"NBodySim/internal/zmapper/approximator/colorist"
	"NBodySim/internal/zmapper/mapper"
	"NBodySim/internal/zmapper/shadowapproximator"
	"NBodySim/internal/zmapper/shadowmapper"
	"image"
	"sync"
)

type SimpleShadowObjectDrawer struct {
	lights       []object.Light
	shadows      *shadowmapper.ShadowMapper
	zmapper      mapper.Zmapper
	approxf      shadowapproximator.ShadowDiscreteApproximatorFabric
	approx       shadowapproximator.ShadowDiscreteApproximator
	view         vector.Vector3d
	cam          *object.Camera
	mut          sync.Mutex
	visitCounter int
	pchan        chan approximator.DiscreteFlatPoint
}

type SimpleShadowObjectDrawerFabric struct {
	zf mapper.ZmapperFabric
	af shadowapproximator.ShadowDiscreteApproximatorFabric
}

func NewSimpleShadowObjectDrawerFabric(zf mapper.ZmapperFabric, af shadowapproximator.ShadowDiscreteApproximatorFabric) *SimpleShadowObjectDrawerFabric {
	return &SimpleShadowObjectDrawerFabric{
		zf: zf,
		af: af,
	}
}

func (f *SimpleShadowObjectDrawerFabric) CreateShadowObjectDrawer() ShadowObjectDrawer {
	return newSimpleShadowObjectDrawer(f.zf, f.af)
}

func newSimpleShadowObjectDrawer(zf mapper.ZmapperFabric, af shadowapproximator.ShadowDiscreteApproximatorFabric) *SimpleShadowObjectDrawer {
	return &SimpleShadowObjectDrawer{
		lights:       []object.Light{},
		shadows:      nil,
		zmapper:      zf.CreateZmapper(),
		approxf:      af,
		approx:       nil,
		mut:          sync.Mutex{},
		visitCounter: 0,
		pchan:        nil,
	}
}

func (sd *SimpleShadowObjectDrawer) startDrawing() {
	sd.mut.Lock()
	if sd.visitCounter == 0 {
		sd.pchan = make(chan approximator.DiscreteFlatPoint)
		sd.approx = sd.approxf.CreateShadowDiscreteApproximator()
		sd.approx.VisitShadowMapper(sd.shadows)

		move := transform.NewMoveAction(vector.NewVector3d(float64(-sd.GetWidth())/2, float64(-sd.GetHeight())/2, 0))
		canvas := transform.NewViewportToCanvas(1/float64(sd.GetWidth()), 1/float64(sd.GetHeight()))
		revpersp := object.NewReversePerspectiveTransform(sd.cam)
		move.ApplyAfter(canvas)
		canvas.ApplyAfter(revpersp)

		sd.approx.ToShadowTransform(move)

		go sd.zmapper.DrawChannel(sd.pchan)
	}
	sd.visitCounter++
	sd.mut.Unlock()
}

func (sd *SimpleShadowObjectDrawer) stopDrawing() {
	sd.mut.Lock()
	sd.visitCounter--
	if sd.visitCounter == 0 {
		close(sd.pchan)
		sd.approx = nil
	}
	sd.mut.Unlock()
}

func (sd *SimpleShadowObjectDrawer) SetView(view vector.Vector3d) {
	sd.view = view
}

func (sd *SimpleShadowObjectDrawer) VisitObjectPool(op *object.ObjectPool) {
	sd.startDrawing()
	for _, obj := range op.GetObjects() {
		obj.Accept(sd)
	}
	sd.stopDrawing()
}

func (sd *SimpleShadowObjectDrawer) VisitPolygonObject(po *object.PolygonObject) {
	sd.startDrawing()
	for _, v := range po.GetPolygons() {
		sd.approx.ApproximatePolygon(v, sd.pchan)
	}
	sd.stopDrawing()
}

func (sd *SimpleShadowObjectDrawer) VisitShadowMapper(mapper *shadowmapper.ShadowMapper) {
	sd.shadows = mapper
}

func (sd *SimpleShadowObjectDrawer) VisitCamera(cam *object.Camera) {
	sd.cam = cam
}

func (sd *SimpleShadowObjectDrawer) VisitPointLight(light *object.PointLight) {
	sd.lights = append(sd.lights, light)
}

func (sd *SimpleShadowObjectDrawer) GetImage() image.Image {
	return sd.zmapper
}

func (sd *SimpleShadowObjectDrawer) ResetImage() {
	sd.zmapper.Reset()
	sd.lights = make([]object.Light, 0)
}

func (sd *SimpleShadowObjectDrawer) GetWidth() int {
	return sd.zmapper.Bounds().Dx()
}

func (sd *SimpleShadowObjectDrawer) GetHeight() int {
	return sd.zmapper.Bounds().Dy()
}

func (sd *SimpleShadowObjectDrawer) GetColorist() colorist.Colorist {
	return sd.approxf.GetColorist()
}

func (sd *SimpleShadowObjectDrawer) SetPointDepth(p *approximator.DiscreteFlatPoint) {
	sd.zmapper.SetPointDepth(p)
}
