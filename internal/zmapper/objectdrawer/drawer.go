package objectdrawer

import (
	"NBodySim/internal/object"
	"image"
)

type ObjectDrawer interface {
	VisitPolygonObject(po *object.PolygonObject)
	VisitCamera(cam *object.Camera)
	VisitPointLight(light *object.PointLight)
	VisitObjectPool(pool *object.ObjectPool)
	GetImage() image.Image
	GetWidth() int
	GetHeight() int
	ResetImage()
}

type ObjectDrawerFabric interface {
	CreateObjectDrawer() ObjectDrawer
}
