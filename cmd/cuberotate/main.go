package main

import (
	"NBodySim/internal/builder"
	"NBodySim/internal/object"
	"NBodySim/internal/reader"
	"NBodySim/internal/transform"
	"NBodySim/internal/vectormath"
	"NBodySim/internal/zmapper"
	"fmt"
	"image/color"
	"math"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
)

func main() {

	read, _ := reader.NewObjReader("/home/impervguin/Projects/NBodySim/models/banana.obj")
	dir := builder.NewPolygonObjectDirector(&builder.ClassicPolygonFactory{}, read)
	cube, err := dir.Construct()
	if err != nil {
		panic(err)
	}

	cam := object.NewCamera(
		*vectormath.NewVector3d(0, 0, -400),
		*vectormath.NewVector3d(0, 0, 1),
		*vectormath.NewVector3d(0, 1, 0),
		1, 1, 1,
	)
	myApp := app.New()
	myWindow := myApp.NewWindow("3dSim")
	myWindow.Resize(fyne.NewSize(1000, 1000))

	myWindow.SetFixedSize(true)

	go func() {
		time.Sleep(time.Second)
		width, height := float64(1000)*float64(myWindow.Canvas().Scale()), float64(1000)*float64(myWindow.Canvas().Scale())
		fmt.Println(width, height)
		var nview *object.CameraViewAction
		var cccube *object.PolygonObject
		var nzmap *zmapper.Zmapper = zmapper.NewZmapper(int(width), int(height), color.RGBA{R: 255, B: 255, G: 255, A: 255}, &zmapper.DepthBufferInfFabric{})
		var nraster *canvas.Raster = canvas.NewRasterWithPixels(nzmap.GetScreenFunction())
		myWindow.SetContent(nraster)
		for {
			time.Sleep(time.Millisecond * 50)
			cam.Transform(transform.NewRotateAction(vectormath.NewVector3d(-math.Pi/12, -math.Pi/12, 0)))
			nview = object.NewCameraViewAction(cam)
			cccube, _ = (cube.Clone()).(*object.PolygonObject)
			cccube.Transform(nview)
			cccube.Transform(transform.NewViewportToCanvas(float64(width), float64(height)))
			cccube.Transform(object.NewPerspectiveTransform(cam))

			// cccube.Transform(transform.NewScaleAction(vectormath.NewVector3d(float64(myWindow.Canvas().Scale()), float64(myWindow.Canvas().Scale()), 1)))
			cccube.Transform(transform.NewMoveAction(vectormath.NewVector3d(float64(width)/2, float64(height)/2, 0)))

			nzmap.Reset()
			cccube.Accept(nzmap)

			myWindow.Content().Refresh()
		}
	}()

	myWindow.ShowAndRun()
}
