package cutter

import (
	"NBodySim/internal/object"
)

// Cuts unseen parts of objects after transforming to camera view
type Cutter interface {
	VisitPolygonObject(po *object.PolygonObject)
	VisitCamera(cam *object.Camera)
}

type CutterFabric interface {
	CreateCutter() Cutter
}