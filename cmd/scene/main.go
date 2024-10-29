package main

import (
	"NBodySim/internal/builder"
	"NBodySim/internal/conveyer"
	"NBodySim/internal/nbody"
	"NBodySim/internal/object"
	"NBodySim/internal/reader"
	"NBodySim/internal/simulation"
	"NBodySim/internal/transform"
	"NBodySim/internal/vectormath"
	"NBodySim/internal/zmapper"
	"image/color"
	"math"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
)

func main() {

	read, _ := reader.NewObjReader("/home/impervguin/Projects/NBodySim/models/6_hexahedron.obj")
	dir := builder.NewPolygonObjectDirector(&builder.ClassicPolygonFactory{}, read)
	cube, err := dir.Construct()
	if err != nil {
		panic(err)
	}

	read, _ = reader.NewObjReader("/home/impervguin/Projects/NBodySim/models/8_octahedron.obj")
	dir = builder.NewPolygonObjectDirector(&builder.ClassicPolygonFactory{}, read)
	cube2, err := dir.Construct()
	if err != nil {
		panic(err)
	}
	cube2.Transform(transform.NewMoveAction(vectormath.NewVector3d(-5, 0, 0)))

	cam := object.NewCamera(
		*vectormath.NewVector3d(0, 0, -10),
		*vectormath.NewVector3d(0, 0, 1),
		*vectormath.NewVector3d(0, 1, 0),
		1, 1, 1,
	)

	sim := simulation.NewSimulation()
	sim.SetCamera(cam)
	body1 := nbody.NewBody(vectormath.NewVector3d(0, 0, 0), vectormath.NewVector3d(0, 0, 0), 1)
	body2 := nbody.NewBody(vectormath.NewVector3d(-5, 0, 0), vectormath.NewVector3d(0, 0, 0), 1)
	sim.AddObject(simulation.NewSimulationObject(cube, body1))
	sim.AddObject(simulation.NewSimulationObject(cube2, body2))

	myApp := app.New()
	myWindow := myApp.NewWindow("3dSim")
	width, height := 800., 800.
	myWindow.Resize(fyne.NewSize(float32(width), float32(height)))

	myWindow.SetFixedSize(true)

	go func() {
		time.Sleep(time.Second)
		width, height = float64(width)*float64(myWindow.Canvas().Scale()), float64(height)*float64(myWindow.Canvas().Scale())

		conv := conveyer.NewSimulationConveyer(
			zmapper.NewSimpleZmapperFabric(),
			int(width),
			int(height),
			color.Black,
			sim,
		)
		var nraster *canvas.Raster = canvas.NewRasterFromImage(conv.GetImage())

		myWindow.SetContent(nraster)
		for {
			time.Sleep(time.Millisecond * 34)
			cam.Transform(transform.NewRotateAction(vectormath.NewVector3d(math.Pi/30, math.Pi/30, 0)))
			conv.Convey()
			myWindow.Content().Refresh()
		}
	}()

	myWindow.ShowAndRun()
}
