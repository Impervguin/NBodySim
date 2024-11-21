package conveyer

import (
	"NBodySim/internal/cutter"
	"NBodySim/internal/mathutils/vector"
	"NBodySim/internal/object"
	"NBodySim/internal/simulation"
	"NBodySim/internal/transform"
	"NBodySim/internal/zmapper/objectdrawer"
	"image"
)

type SimplestSimulationConveyer struct {
	drawer objectdrawer.ObjectDrawerWithoutLights
	sim    *simulation.Simulation
}

func NewSimplestSimulationConveyer(fabric objectdrawer.ObjectDrawerWithoutLightsFabric, sim *simulation.Simulation) *SimplestSimulationConveyer {
	return &SimplestSimulationConveyer{
		drawer: fabric.CreateObjectDrawerWithoutLights(),
		sim:    sim,
	}
}

func (sc *SimplestSimulationConveyer) GetImage() image.Image {
	return sc.drawer.GetImage()
}

func (sc *SimplestSimulationConveyer) Convey() error {
	sc.drawer.ResetImage()

	objs := sc.sim.GetObjectsClone()
	camo := sc.sim.GetCamera().Clone()
	cam, _ := camo.(*object.Camera)

	view := object.NewCameraViewAction(cam)

	objs.Transform(view)
	cam.Transform(view)

	persp := object.NewPerspectiveTransform(cam)
	canvas := transform.NewViewportToCanvas(float64(sc.drawer.GetWidth()), float64(sc.drawer.GetHeight()))

	// With knowing, that screen coordinate system starts at (0, 0) and ends at (width, height),
	move := transform.NewMoveAction(vector.NewVector3d(float64(sc.drawer.GetWidth())/2, float64(sc.drawer.GetHeight())/2, 0))

	cut := cutter.NewSimpleCamCutter(cam)
	objs.Accept(cut)

	objs.Transform(persp)
	objs.Transform(canvas)
	objs.Transform(move)

	objs.Accept(sc.drawer)

	return nil
}
