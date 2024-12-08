package main

import (
	"NBodySim/internal/builder"
	"NBodySim/internal/conveyer"
	"NBodySim/internal/mathutils/vector"
	"NBodySim/internal/object"
	"NBodySim/internal/reader"
	"NBodySim/internal/simulation"
	"NBodySim/internal/transform"
	"NBodySim/internal/zmapper/approximator"
	"NBodySim/internal/zmapper/buffers"
	"NBodySim/internal/zmapper/mapper"
	"NBodySim/internal/zmapper/objectdrawer"
	"fmt"
	"image/color"
	"math"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
)

const FPS = 20

func main() {

	read, _ := reader.NewObjReader("/home/impervguin/Projects/NBodySim/models/6_hexahedron.obj")
	dir := builder.NewPolygonObjectDirector(&builder.ClassicPolygonFactory{}, read)
	cube, err := dir.Construct()
	if err != nil {
		panic(err)
	}
	cube.Transform(transform.NewMoveAction(vector.NewVector3d(0, 0, 30)))

	read, _ = reader.NewObjReader("/home/impervguin/Projects/NBodySim/models/6_hexahedron.obj")
	dir = builder.NewPolygonObjectDirector(&builder.ClassicPolygonFactory{}, read)
	cube2, err := dir.Construct()
	if err != nil {
		panic(err)
	}
	cube2.Transform(transform.NewMoveAction(vector.NewVector3d(30, 0, 0)))

	read, _ = reader.NewObjReader("/home/impervguin/Projects/NBodySim/models/6_hexahedron.obj")
	dir = builder.NewPolygonObjectDirector(&builder.ClassicPolygonFactory{}, read)
	cube3, err := dir.Construct()
	if err != nil {
		panic(err)
	}
	cube3.Transform(transform.NewMoveAction(vector.NewVector3d(0, 0, -30)))
	// fmt.Println(cube3.GetCenter())

	read, _ = reader.NewObjReader("/home/impervguin/Projects/NBodySim/models/6_hexahedron.obj")
	dir = builder.NewPolygonObjectDirector(&builder.ClassicPolygonFactory{}, read)
	cube4, err := dir.Construct()
	if err != nil {
		panic(err)
	}
	cube4.Transform(transform.NewMoveAction(vector.NewVector3d(-30, 0, 0)))

	cam := object.NewCamera(
		*vector.NewVector3d(0, 0, -120),
		*vector.NewVector3d(0, 0, 1),
		*vector.NewVector3d(0, 1, 0),
		1, 1, 1,
	)

	light1 := object.NewPointLightShadow(color.RGBA{255, 255, 255, 255}, *vector.NewVector3d(0, 10, 0))
	light2 := object.NewPointLightShadow(color.RGBA{255, 255, 255, 255}, *vector.NewVector3d(5, 0, 5))
	// light3 := object.NewPointLight(color.RGBA{255, 255, 255, 255}, *vector.NewVector3d(0, 0, -15))
	// light4 := object.NewPointLight(color.RGBA{255, 255, 255, 255}, *vector.NewVector3d(0, 10, -5))

	sim := simulation.NewSimulation()
	sim.SetCamera(cam)
	sim.AddObject(cube, *vector.NewVector3d(0.2, 0, 0), 100000000000)
	sim.AddObject(cube2, *vector.NewVector3d(0, 0, -.2), 100000000000)
	sim.AddObject(cube3, *vector.NewVector3d(-.2, 0, 0), 100000000000)
	sim.AddObject(cube4, *vector.NewVector3d(0, 0, .2), 100000000000)
	sim.AddLight(light1)
	sim.AddLight(light2)
	// sim.AddLight(light3)
	// sim.AddLight(light4)
	sim.SetDt(0.00001)

	myApp := app.New()
	myWindow := myApp.NewWindow("3dSim")
	width, height := 800., 800.
	myWindow.Resize(fyne.NewSize(float32(width), float32(height)))

	myWindow.SetFixedSize(true)

	go func() {
		time.Sleep(time.Second)
		width, height = float64(width)*float64(myWindow.Canvas().Scale()), float64(height)*float64(myWindow.Canvas().Scale())
		cam.Transform(transform.NewRotateAction(vector.NewVector3d(-math.Pi/4, 0, 0)))
		drawerfac := objectdrawer.NewParallelWithoutLightsDrawerFabric(mapper.NewParallelZmapperWithNormalsFabric(int(width), int(height), color.Black, &buffers.DepthBufferNullFabric{}), approximator.NewFlatNormalApproximatorFabric())
		conv := conveyer.NewRefactoredSimulationConveyer(
			drawerfac,
			sim,
		)
		var nraster *canvas.Raster = canvas.NewRasterFromImage(conv.GetImage())

		myWindow.SetContent(nraster)
		FPStime := time.Second / FPS
		timeToSleep := FPStime
		for {
			time.Sleep(timeToSleep)
			startTime := time.Now()
			sim.UpdateFor(1)
			// cam.Transform(transform.NewRotateAction(vector.NewVector3d(0, math.Pi/60, 0)))
			conv.Convey()
			nraster.Refresh()
			endTime := time.Now()
			doneIn := endTime.Sub(startTime)
			fmt.Println(doneIn)
			timeToSleep = FPStime
		}
	}()

	myWindow.ShowAndRun()
}
