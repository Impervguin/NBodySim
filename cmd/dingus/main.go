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
	"image/png"
	"math"
	"os"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
)

const FPS = 20

func DegroidCsvOutput(filename string, drawer objectdrawer.ObjectDrawerWithoutLights) error {
	im := drawer.GetImage()
	w, h := im.Bounds().Size().X, im.Bounds().Size().Y
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			col := im.At(j, i)
			r, _, _, _ := col.RGBA()
			rf := float64(r) / float64(65535)
			fmt.Fprintf(f, "%.4f", rf)
			if j < w-1 {
				fmt.Fprint(f, ",")
			}
		}
		fmt.Fprintln(f)
	}
	return nil
}

func SavePng(filename string, drawer objectdrawer.ObjectDrawerWithoutLights) error {
	im := drawer.GetImage()
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	err = png.Encode(f, im)
	if err != nil {
		return err
	}
	return nil

}

func main() {
	read, _ := reader.NewObjReader("/home/impervguin/Projects/NBodySim/models/dinguser.obj")
	dir := builder.NewPolygonObjectDirector(&builder.ClassicPolygonFactory{}, read)
	dingus, err := dir.Construct()
	if err != nil {
		panic(err)
	}
	dingus.Transform(transform.NewScaleAction(vector.NewVector3d(1./250, 1./250, 1./250)))
	dingus.Transform(transform.NewMoveAction(vector.NewVector3d(0, -3, 0)))
	dobj := dingus.(*object.PolygonObject)
	dobj.SetColor(color.White)

	cam := object.NewCamera(
		*vector.NewVector3d(0, 0, -20),
		*vector.NewVector3d(0, 0, 1),
		*vector.NewVector3d(0, -1, 0),
		1, 1, 1,
	)

	// fmt.Println(dingus.GetCenter())

	light1 := object.NewPointLightShadow(color.RGBA{255, 255, 255, 255}, *vector.NewVector3d(0, 10, 0))
	light2 := object.NewPointLightShadow(color.RGBA{255, 255, 255, 255}, *vector.NewVector3d(5, 10, 5))

	sim := simulation.NewSimulation()
	sim.SetCamera(cam)
	sim.AddObject(dingus, *vector.NewVector3d(0, 0, 0), 1)
	sim.AddLight(light1)
	sim.AddLight(light2)

	myApp := app.New()
	myWindow := myApp.NewWindow("3dSim")
	width, height := 1440., 900.
	myWindow.Resize(fyne.NewSize(float32(width), float32(height)))

	myWindow.SetFixedSize(true)

	go func() {
		time.Sleep(time.Second)
		width, height = float64(width)*float64(myWindow.Canvas().Scale()), float64(height)*float64(myWindow.Canvas().Scale())
		fmt.Println(width, height)
		cam.Transform(transform.NewRotateAction(vector.NewVector3d(math.Pi/6, 0, 0)))
		cam.Transform(transform.NewRotateAction(vector.NewVector3d(0, math.Pi, 0)))
		drawerfac := objectdrawer.NewParallelWithoutLightsDrawerFabric(mapper.NewParallelZmapperWithNormalsFabric(int(width), int(height), color.RGBA{137, 142, 140, 255}, &buffers.DepthBufferNullFabric{}), approximator.NewFlatNormalApproximatorFabric())
		conv := conveyer.NewRefactoredShadowSimulationConveyer(
			drawerfac,
			sim,
		)
		var nraster *canvas.Raster = canvas.NewRasterFromImage(conv.GetImage())

		myWindow.SetContent(nraster)
		// FPStime := time.Second / FPS
		// timeToSleep := FPStime
		startTime := time.Now()
		conv.Convey()
		nraster.Refresh()
		endTime := time.Now()
		doneIn := endTime.Sub(startTime)
		fmt.Println(doneIn)
		// DegroidCsvOutput("output.csv", conv.GetDrawer())
		// fmt.Println("outed")
		err = SavePng("dingus.png", conv.GetDrawer())
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("png")
	}()

	myWindow.ShowAndRun()
}
