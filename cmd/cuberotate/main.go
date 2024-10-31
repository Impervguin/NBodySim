package main

import (
	"NBodySim/internal/builder"
	"NBodySim/internal/mathutils/vector"
	"NBodySim/internal/object"
	"NBodySim/internal/reader"
	"NBodySim/internal/transform"
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
		*vector.NewVector3d(0, 0, -400),
		*vector.NewVector3d(0, 0, 1),
		*vector.NewVector3d(0, 1, 0),
		1, 1, 1,
	)
	myApp := app.New()
	myWindow := myApp.NewWindow("3dSim")
	width, height := 800., 800.
	myWindow.Resize(fyne.NewSize(float32(width), float32(height)))

	myWindow.SetFixedSize(true)

	go func() {
		time.Sleep(time.Second)
		width, height = float64(width)*float64(myWindow.Canvas().Scale()), float64(height)*float64(myWindow.Canvas().Scale())
		fmt.Println(width, height)
		var nview *object.CameraViewAction
		var cccube *object.PolygonObject
		var nzmap *zmapper.Zmapper = zmapper.NewZmapper(int(width), int(height), color.RGBA{R: 255, B: 255, G: 255, A: 255}, &zmapper.DepthBufferInfFabric{})
		var nraster *canvas.Raster = canvas.NewRasterWithPixels(nzmap.GetScreenFunction())
		myWindow.SetContent(nraster)
		for {
			time.Sleep(time.Millisecond * 50)
			cam.Transform(transform.NewRotateAction(vector.NewVector3d(-math.Pi/12, -math.Pi/12, 0)))
			nview = object.NewCameraViewAction(cam)
			cccube, _ = (cube.Clone()).(*object.PolygonObject)
			cccube.Transform(nview)
			cccube.Transform(transform.NewViewportToCanvas(float64(width), float64(height)))
			cccube.Transform(object.NewPerspectiveTransform(cam))

			// cccube.Transform(transform.NewScaleAction(vector.NewVector3d(float64(myWindow.Canvas().Scale()), float64(myWindow.Canvas().Scale()), 1)))
			cccube.Transform(transform.NewMoveAction(vector.NewVector3d(float64(width)/2, float64(height)/2, 0)))

			nzmap.Reset()
			cccube.Accept(nzmap)

			myWindow.Content().Refresh()
		}
	}()

	myWindow.ShowAndRun()
}
