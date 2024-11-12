package conveyer

import (
	"NBodySim/internal/cutter"
	"NBodySim/internal/mathutils/vector"
	"NBodySim/internal/object"
	"NBodySim/internal/simulation"
	"NBodySim/internal/transform"
	"NBodySim/internal/zmapper/objectdrawer"
	"NBodySim/internal/zmapper/shadowmapper"
	"image"
)

type SimulationConveyer struct {
	drawer objectdrawer.ObjectDrawer
	sim    *simulation.Simulation
}

func NewSimulationConveyer(fabric objectdrawer.ObjectDrawerFabric, sim *simulation.Simulation) *SimulationConveyer {
	return &SimulationConveyer{
		drawer: fabric.CreateObjectDrawer(),
		sim:    sim,
	}
}

func (sc *SimulationConveyer) GetImage() image.Image {
	return sc.drawer.GetImage()
}

func (sc *SimulationConveyer) Convey() error {
	sc.drawer.ResetImage()

	objs := sc.sim.GetObjectsClone()
	shadows := shadowmapper.NewShadowMapper(512)
	lights := sc.sim.GetLightsClone()
	camo := sc.sim.GetCamera().Clone()
	cam, _ := camo.(*object.Camera)

	view := object.NewCameraViewAction(cam)
	canvas := transform.NewViewportToCanvas(float64(sc.drawer.GetWidth()), float64(sc.drawer.GetHeight()))
	persp := object.NewPerspectiveTransform(cam)
	// With knowing, that screen coordinate system starts at (0, 0) and ends at (width, height),
	move := transform.NewMoveAction(vector.NewVector3d(float64(sc.drawer.GetWidth())/2, float64(sc.drawer.GetHeight())/2, 0))

	objs.Transform(view)
	cam.Transform(view)
	lights.Transform(view)

	lights.Accept(shadows)
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

	objs.Accept(sc.drawer)

	return nil
}
