package cutter

import (
	"NBodySim/internal/mathutils/vector"
	"NBodySim/internal/object"
)

// Cuts unseen parts of objects after transforming to camera view
type Cutter interface {
	VisitPolygonObject(po *object.PolygonObject)
	VisitCamera(cam *object.Camera)
	VisitObjectPool(pool *object.ObjectPool)
	SeePoint(point *vector.Vector3d) bool
}

type CutterFabric interface {
	CreateCutter() Cutter
}
