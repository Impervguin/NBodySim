package object

type ObjectVisitor interface {
	VisitPolygonObject(po *PolygonObject)
	VisitCamera(cam *Camera)
	VisitObjectPool(pool *ObjectPool)
}

type LightVisitor interface {
	VisitPointLight(light *PointLight)
}


