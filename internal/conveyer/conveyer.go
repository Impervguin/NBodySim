package conveyer

import (
	"NBodySim/internal/cutter"
	"NBodySim/internal/mathutils/vector"
	"NBodySim/internal/object"
	"NBodySim/internal/simulation"
	"NBodySim/internal/transform"
	"NBodySim/internal/zmapper/mapper"
	"image"
	"image/color"
)

type SimulationConveyer struct {
	zmapper mapper.Zmapper
	swidth  int
	sheight int
	sim     *simulation.Simulation
}

func NewSimulationConveyer(fabric mapper.ZmapperFabric, swidth, sheight int, background color.Color, sim *simulation.Simulation) *SimulationConveyer {
	return &SimulationConveyer{
		zmapper: fabric.CreateZmapper(swidth, sheight, background),
		swidth:  swidth,
		sheight: sheight,
		sim:     sim,
	}
}

func (sc *SimulationConveyer) GetImage() image.Image {
	return sc.zmapper
}

func (sc *SimulationConveyer) Convey() error {
	sc.zmapper.Reset()
	objs := sc.sim.GetObjectsClone()
	cam := sc.sim.GetCamera()

	view := object.NewCameraViewAction(cam)
	canvas := transform.NewViewportToCanvas(float64(sc.zmapper.Bounds().Dx()), float64(sc.zmapper.Bounds().Dy()))
	persp := object.NewPerspectiveTransform(cam)

	objs.Transform(view)

	cut := cutter.NewSimpleCamCutter(cam)
	objs.Accept(cut)

	objs.Transform(canvas)
	objs.Transform(persp)

	// With knowing, that screen coordinate system starts at (0, 0) and ends at (width, height),
	move := transform.NewMoveAction(vector.NewVector3d(float64(sc.swidth)/2, float64(sc.sheight)/2, 0))
	objs.Transform(move)

	objs.Accept(sc.zmapper)

	return nil
}
