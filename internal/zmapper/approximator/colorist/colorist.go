package colorist

import "NBodySim/internal/object"

type Colorist interface {
	VisitPolygonObject(po *object.PolygonObject)
	VisitCamera(cam *object.Camera)
	VisitPointLight(light *object.PointLight)
	VisitObjectPool(pool *object.ObjectPool)
}
