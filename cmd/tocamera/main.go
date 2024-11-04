package main

import (
	"NBodySim/internal/builder"
	"NBodySim/internal/conveyer"
	"NBodySim/internal/mathutils/vector"
	"NBodySim/internal/object"
	"NBodySim/internal/reader"
	"NBodySim/internal/simulation"
	"NBodySim/internal/transform"
	"NBodySim/internal/zmapper/mapper"
	"image/color"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
)

func main() {

	read, _ := reader.NewObjReader("/home/impervguin/Projects/NBodySim/models/20_icosahedron.obj")
	dir := builder.NewPolygonObjectDirector(&builder.ClassicPolygonFactory{}, read)
	cube, err := dir.Construct()
	if err != nil {
		panic(err)
	}

	cam := object.NewCamera(
		*vector.NewVector3d(0, 0, -100),
		*vector.NewVector3d(0, 0, 1),
		*vector.NewVector3d(0, 1, 0),
		1, 1, 1,
	)

	sim := simulation.NewSimulation()
	sim.SetCamera(cam)
	sim.AddObject(cube, *vector.NewVector3d(0, 0, 0), 10000000000)

	sim.SetDt(0.01)

	myApp := app.New()
	myWindow := myApp.NewWindow("3dSim")
	width, height := 800., 800.
	myWindow.Resize(fyne.NewSize(float32(width), float32(height)))

	myWindow.SetFixedSize(true)

	go func() {
		time.Sleep(time.Second)
		width, height = float64(width)*float64(myWindow.Canvas().Scale()), float64(height)*float64(myWindow.Canvas().Scale())
		// cam.Transform(transform.NewRotateAction(vector.NewVector3d(math.Pi/4, -math.Pi/4, 0)))
		cube.Transform(transform.NewMoveAction(vector.NewVector3d(0, 0, 1000)))
		conv := conveyer.NewSimulationConveyer(
			// mapper.NewSimpleZmapperFabric(),
			mapper.NewParallelPerPolygonZmapperFabric(),
			int(width),
			int(height),
			color.White,
			sim,
		)
		var nraster *canvas.Raster = canvas.NewRasterFromImage(conv.GetImage())

		myWindow.SetContent(nraster)
		for {
			time.Sleep(time.Millisecond * 34)
			sim.UpdateFor(1)
			// cam.Transform(transform.NewRotateAction(vector.NewVector3d(0, math.Pi/60, 0)))
			cube.Transform(transform.NewMoveAction(vector.NewVector3d(0, 0, -10)))
			conv.Convey()
			myWindow.Content().Refresh()
		}
	}()

	myWindow.ShowAndRun()
}
