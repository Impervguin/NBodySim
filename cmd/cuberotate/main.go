package main

import (
	"NBodySim/internal/object"
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

var vertices []vectormath.Vector3d = []vectormath.Vector3d{
	*vectormath.NewVector3d(1, 1, 1),
	*vectormath.NewVector3d(-1, 1, 1),
	*vectormath.NewVector3d(-1, -1, 1),
	*vectormath.NewVector3d(1, -1, 1),
	*vectormath.NewVector3d(1, 1, -1),
	*vectormath.NewVector3d(-1, 1, -1),
	*vectormath.NewVector3d(-1, -1, -1),
	*vectormath.NewVector3d(1, -1, -1),
}
var polygons []object.Polygon = []object.Polygon{
	*object.NewPolygon(&vertices[0], &vertices[1], &vertices[2], color.RGBA{255, 0, 0, 255}),
	*object.NewPolygon(&vertices[0], &vertices[2], &vertices[3], color.RGBA{255, 0, 0, 255}),
	*object.NewPolygon(&vertices[4], &vertices[0], &vertices[3], color.RGBA{0, 255, 0, 255}),
	*object.NewPolygon(&vertices[4], &vertices[3], &vertices[7], color.RGBA{0, 255, 0, 255}),
	*object.NewPolygon(&vertices[5], &vertices[4], &vertices[7], color.RGBA{0, 0, 255, 255}),
	*object.NewPolygon(&vertices[5], &vertices[7], &vertices[6], color.RGBA{0, 0, 255, 255}),
	*object.NewPolygon(&vertices[1], &vertices[5], &vertices[6], color.RGBA{255, 255, 0, 255}),
	*object.NewPolygon(&vertices[1], &vertices[6], &vertices[2], color.RGBA{255, 255, 0, 255}),
	*object.NewPolygon(&vertices[4], &vertices[5], &vertices[1], color.RGBA{255, 0, 255, 255}),
	*object.NewPolygon(&vertices[4], &vertices[1], &vertices[0], color.RGBA{255, 0, 255, 255}),
	*object.NewPolygon(&vertices[2], &vertices[6], &vertices[7], color.RGBA{0, 255, 255, 255}),
	*object.NewPolygon(&vertices[2], &vertices[7], &vertices[3], color.RGBA{0, 255, 255, 255}),
}

var cube object.PolygonObject = *object.NewPolygonObject(vertices, polygons, *vectormath.NewVector3d(0, 0, 0))

func main() {
	cam := object.NewCamera(
		*vectormath.NewVector3d(0, 0, -5),
		*vectormath.NewVector3d(0, 0, 1),
		*vectormath.NewVector3d(0, 1, 0),
		1, 1, 1,
	)
	myApp := app.New()
	myWindow := myApp.NewWindow("3dSim")
	myWindow.Resize(fyne.NewSize(1000, 1000))
	myWindow.SetFixedSize(true)

	go func() {
		var nview *object.CameraViewAction
		var cccube *object.PolygonObject
		var nzmap *zmapper.Zmapper = zmapper.NewZmapper(1000, 1000, color.RGBA{R: 255, B: 255, G: 255, A: 255}, &zmapper.DepthBufferInfFabric{})
		var nraster *canvas.Raster = canvas.NewRasterWithPixels(nzmap.GetScreenFunction())
		myWindow.SetContent(nraster)
		for {
			time.Sleep(time.Millisecond * 50)
			cam.Transform(transform.NewRotateAction(vectormath.NewVector3d(-math.Pi/12, -math.Pi/12, 0)))
			nview = object.NewCameraViewAction(cam)
			cccube, _ = (cube.Clone()).(*object.PolygonObject)
			cccube.Transform(nview)
			cccube.Transform(transform.NewViewportToCanvas(1000, 1000))
			cccube.Transform(object.NewPerspectiveTransform(cam))
			cccube.Transform(transform.NewMoveAction(vectormath.NewVector3d(500, 500, 0)))

			nzmap.Reset()
			cccube.Accept(nzmap)

			myWindow.Content().Refresh()
		}
	}()

	myWindow.ShowAndRun()
}
