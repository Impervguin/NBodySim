package mapper

import (
	"NBodySim/internal/object"
	"NBodySim/internal/zmapper/buffers"
	"image"
	"image/color"
)

type Zmapper interface {
	image.Image
	VisitPolygonObject(po *object.PolygonObject)
	VisitCamera(cam *object.Camera)
	GetScreenFunction() buffers.ScreenFunction
	Reset()
}

type ZmapperFabric interface {
	CreateZmapper(width, height int, background color.Color) Zmapper
}
