package colorist

import "NBodySim/internal/object"

type Colorist interface {
	VisitPolygonObject(po *object.PolygonObject)
	VisitCamera(cam *object.Camera)
	VisitPointLight(light *object.PointLight)
	VisitPointLightShadow(light *object.PointLightShadow)
	VisitObjectPool(pool *object.ObjectPool)
}
