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

type RefactoredSimulationConveyer struct {
	drawer objectdrawer.ObjectDrawerWithoutLights
	sim    *simulation.Simulation
}

func NewRefactoredSimulationConveyer(fabric objectdrawer.ObjectDrawerWithoutLightsFabric, sim *simulation.Simulation) *RefactoredSimulationConveyer {
	return &RefactoredSimulationConveyer{
		drawer: fabric.CreateObjectDrawerWithoutLights(),
		sim:    sim,
	}
}

func (sc *RefactoredSimulationConveyer) GetImage() image.Image {
	return sc.drawer.GetImage()
}

func (sc *RefactoredSimulationConveyer) Convey() error {
	sc.drawer.ResetImage()

	imobjs := sc.sim.GetImaginaryObjectsClone()
	objs := sc.sim.GetObjectsClone()
	lights := sc.sim.GetLightsClone()
	camo := sc.sim.GetCamera().Clone()
	cam, _ := camo.(*object.Camera)

	view := object.NewCameraViewAction(cam)

	imobjs.Transform(view)
	objs.Transform(view)
	cam.Transform(view)
	lights.Transform(view)

	persp := object.NewPerspectiveTransform(cam)
	canvas := transform.NewViewportToCanvas(float64(sc.drawer.GetWidth()), float64(sc.drawer.GetHeight()))
	// With knowing, that screen coordinate system starts at (0, 0) and ends at (width, height),
	move := transform.NewMoveAction(vector.NewVector3d(float64(sc.drawer.GetWidth())/2, float64(sc.drawer.GetHeight())/2, 0))

	bcut := cutter.NewBackwardsCutter(cam)
	objs.Accept(bcut)
	imobjs.Accept(bcut)

	cut := cutter.NewSimpleCamCutter(cam)
	objs.Accept(cut)
	imobjs.Accept(cut)

	imobjs.Transform(persp)
	imobjs.Transform(canvas)
	imobjs.Transform(move)

	objs.Transform(persp)
	objs.Transform(canvas)
	objs.Transform(move)

	objs.Accept(sc.drawer)

	mapper := sc.drawer.GetZmapper()

	moveBack := transform.NewMoveAction(vector.NewVector3d(float64(-sc.drawer.GetWidth())/2, float64(-sc.drawer.GetHeight())/2, 0))
	canvasBack := transform.NewViewportToCanvas(1/float64(sc.drawer.GetWidth()), 1/float64(sc.drawer.GetHeight()))
	revpersp := object.NewReversePerspectiveTransform(cam)
	moveBack.ApplyAfter(canvasBack)
	canvasBack.ApplyAfter(revpersp)

	mapper.ApplyLight(lights, moveBack)
	imobjs.Accept(sc.drawer)

	return nil
}
