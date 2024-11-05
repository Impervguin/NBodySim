package objectdrawer

import (
	"NBodySim/internal/mathutils/vector"
	"NBodySim/internal/object"
	"NBodySim/internal/zmapper/approximator/colorist"
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
	GetColorist(view vector.Vector3d) colorist.Colorist
}

type ObjectDrawerFabric interface {
	CreateObjectDrawer() ObjectDrawer
}
