package objectdrawer

import (
	"NBodySim/internal/object"
	"NBodySim/internal/zmapper/approximator"
	"NBodySim/internal/zmapper/approximator/colorist"
	"NBodySim/internal/zmapper/mapper"
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

type ObjectDrawerWithoutLights interface {
	VisitPolygonObject(po *object.PolygonObject)
	VisitCamera(cam *object.Camera)
	VisitObjectPool(pool *object.ObjectPool)
	GetZmapper() mapper.NormalZmapper
	SetZmapper(zm mapper.NormalZmapper)
	GetImage() image.Image
	GetWidth() int
	GetHeight() int
	ResetImage()
}

type ObjectDrawerWithoutLightsFabric interface {
	CreateObjectDrawerWithoutLights() ObjectDrawerWithoutLights
}
