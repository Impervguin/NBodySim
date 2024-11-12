package mapper

import (
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
