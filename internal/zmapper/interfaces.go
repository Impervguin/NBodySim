package zmapper

import (
	"NBodySim/internal/object"
	"image"
	"image/color"
)

type Zmapper interface {
	image.Image
	VisitPolygonObject(po *object.PolygonObject)
	VisitCamera(cam *object.Camera)
	GetScreenFunction() ScreenFunction
	Reset()
}

type ZmapperFabric interface {
	CreateZmapper(width, height int, background color.Color) Zmapper
}
