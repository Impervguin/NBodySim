package mapper

import (
	"NBodySim/internal/object"
	"NBodySim/internal/transform"
	"NBodySim/internal/zmapper/approximator"
	"NBodySim/internal/zmapper/buffers"
	"image"
)

type Zmapper interface {
	image.Image
	DrawChannel(ch <-chan approximator.DiscreteFlatPoint)
	SetPointDepth(p *approximator.DiscreteFlatPoint)
	GetScreenFunction() buffers.ScreenFunction
	Reset()
}

type ZmapperFabric interface {
	CreateZmapper() Zmapper
}

type NormalZmapper interface {
	image.Image
	DrawChannel(ch <-chan approximator.DiscreteNormalPoint)
	GetPoint(x, y int) *approximator.DiscreteNormalPoint
	GetScreenFunction() buffers.ScreenFunction
	ApplyLight(lp *object.LightPool, tolights transform.TransformAction)
	Reset()
}

type NormalZmapperFabric interface {
	CreateNormalZmapper() NormalZmapper
}
