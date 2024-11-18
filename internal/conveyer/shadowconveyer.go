package conveyer

import (
	"NBodySim/internal/cutter"
	"NBodySim/internal/mathutils/vector"
	"NBodySim/internal/object"
	"NBodySim/internal/simulation"
	"NBodySim/internal/transform"
	"NBodySim/internal/zmapper/shadowdrawer"
	"NBodySim/internal/zmapper/shadowmapper"
	"image"
)

type ShadowSimulationConveyer struct {
	shadows shadowmapper.ShadowMapper
	drawer  shadowdrawer.ShadowObjectDrawer
	sim     *simulation.Simulation
}

func NewShadowSimulationConveyer(fabric shadowdrawer.ShadowObjectDrawerFabric, sim *simulation.Simulation) *ShadowSimulationConveyer {
	return &ShadowSimulationConveyer{
		drawer: fabric.CreateShadowObjectDrawer(),
		sim:    sim,
	}
}

func (sc *ShadowSimulationConveyer) GetImage() image.Image {
	return sc.drawer.GetImage()
}

func (sc *ShadowSimulationConveyer) Convey() error {
	sc.drawer.ResetImage()

	objs := sc.sim.GetObjectsClone()
	shadows := shadowmapper.NewShadowMapper(256)
	lights := sc.sim.GetLightsClone()
	camo := sc.sim.GetCamera().Clone()
	cam, _ := camo.(*object.Camera)

	view := object.NewCameraViewAction(cam)
	canvas := transform.NewViewportToCanvas(float64(sc.drawer.GetWidth()), float64(sc.drawer.GetHeight()))
	persp := object.NewPerspectiveTransform(cam)
	// With knowing, that screen coordinate system starts at (0, 0) and ends at (width, height),
	move := transform.NewMoveAction(vector.NewVector3d(float64(sc.drawer.GetWidth())/2, float64(sc.drawer.GetHeight())/2, 0))

	lights.Accept(shadows)
	cam.Accept(shadows)

	objs.Transform(view)
	cam.Transform(view)
	lights.Transform(view)

	objs.Accept(shadows)

	cut := cutter.NewSimpleCamCutter(cam)
	objs.Accept(cut)

	colorist := sc.drawer.GetColorist()
	cam.Accept(colorist)
	lights.Accept(colorist)
	objs.Accept(colorist)

	objs.Transform(persp)
	objs.Transform(canvas)
	objs.Transform(move)

	sc.drawer.VisitShadowMapper(shadows)
	cam.Accept(sc.drawer)
	objs.Accept(sc.drawer)

	return nil
}
