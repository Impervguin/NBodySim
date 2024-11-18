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

	objs := sc.sim.GetObjectsClone()
	lights := sc.sim.GetLightsClone()
	camo := sc.sim.GetCamera().Clone()
	cam, _ := camo.(*object.Camera)

	view := object.NewCameraViewAction(cam)

	objs.Transform(view)
	cam.Transform(view)
	lights.Transform(view)
	persp := object.NewPerspectiveTransform(cam)
	canvas := transform.NewViewportToCanvas(float64(sc.drawer.GetWidth()), float64(sc.drawer.GetHeight()))
	// With knowing, that screen coordinate system starts at (0, 0) and ends at (width, height),
	move := transform.NewMoveAction(vector.NewVector3d(float64(sc.drawer.GetWidth())/2, float64(sc.drawer.GetHeight())/2, 0))

	cut := cutter.NewSimpleCamCutter(cam)
	objs.Accept(cut)

	// tobj, _ := objs.GetObject(1)
	// pobj, _ := tobj.(*object.PolygonObject)

	// norm := pobj.GetPolygons()[0].Clone()

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

	// fmt.Println(norm.GetNormal())
	// norm.TransformNormal(persp)
	// norm.TransformNormal(canvas)
	// norm.TransformNormal(move)
	// norm.TransformNormal(moveBack)
	// fmt.Println(norm.GetNormal())
	// os.Exit(1)

	// p := vector.NewVector3d(0, 0, 5)
	// view.ApplyToVector(p)
	// fmt.Println(p)
	// persp.ApplyToVector(p)
	// canvas.ApplyToVector(p)
	// move.ApplyToVector(p)
	// moveBack.ApplyToVector(p)
	// // fmt.Println(p)
	// var l object.Light
	// var b bool
	// for id := 0; true; id++ {
	// 	l, b = lights.GetLight(int64(id))
	// 	if b {
	// 		break
	// 	}
	// }

	// lc := l.GetCenter()
	// fmt.Println(vector.SubtractVectors(&lc, p))

	return nil
}
