package objectdrawer

import (
	"NBodySim/internal/object"
	"NBodySim/internal/zmapper/approximator"
	"NBodySim/internal/zmapper/approximator/colorist"
	"image"
)

type ObjectDrawer interface {
	VisitPolygonObject(po *object.PolygonObject)
	VisitCamera(cam *object.Camera)
	VisitPointLight(light *object.PointLight)
	VisitObjectPool(pool *object.ObjectPool)
	SetPointDepth(p *approximator.DiscreteFlatPoint)
	GetImage() image.Image
	GetWidth() int
	GetHeight() int
	ResetImage()
	GetColorist() colorist.Colorist
}

type ObjectDrawerFabric interface {
	CreateObjectDrawer() ObjectDrawer
}
