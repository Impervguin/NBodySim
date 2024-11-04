package object

type ObjectVisitor interface {
	VisitPolygonObject(po *PolygonObject)
	VisitCamera(cam *Camera)
}

type LightVisitor interface {
	VisitPointLight(light *PointLight)
}
