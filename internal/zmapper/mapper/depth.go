package mapper

import (
	"NBodySim/internal/zmapper/approximator"
	"NBodySim/internal/zmapper/buffers"
	"image"
	"image/color"
)

type DepthZmapper struct {
	dbuf          buffers.DepthBuffer
	background    color.Color
	width, height int
}

type DepthZmapperFabric struct {
	width, height int
	background    color.Color
}

func NewDepthZmapperFabric(width, height int, background color.Color) *DepthZmapperFabric {
	return &DepthZmapperFabric{
		width:      width,
		height:     height,
		background: background,
	}
}

func (f *DepthZmapperFabric) CreateZmapper() Zmapper {
	return newDepthZmapper(f.width, f.height, f.background, &buffers.DepthBufferInfFabric{})
}

func (f *DepthZmapper) SetPointDepth(p *approximator.DiscreteFlatPoint) {
	p.Z, _ = f.dbuf.GetDepth(p.X, p.Y)
}

func newDepthZmapper(width, height int, background color.Color, df buffers.DepthBufferFabric) *DepthZmapper {
	return &DepthZmapper{
		width:      width,
		height:     height,
		background: background,
		dbuf:       df.CreateDepthBuffer(width, height),
	}
}

func (zm *DepthZmapper) setPoint(x, y int, z float64, _ color.Color) {
	zm.dbuf.PutPoint(x, y, z)
}

func (zm *DepthZmapper) DrawChannel(ch <-chan approximator.DiscreteFlatPoint) {
	for dp := range ch {
		zm.setPoint(dp.X, dp.Y, dp.Z, dp.Color)
	}
}

func (zm *DepthZmapper) GetScreenFunction() buffers.ScreenFunction {
	return func(x, y, w, h int) color.Color {
		return zm.background
	}
}

func (zm *DepthZmapper) Reset() {
	zm.dbuf.Reset()
}

func (zm *DepthZmapper) ColorModel() color.Model {
	return color.RGBAModel
}

func (zm *DepthZmapper) Bounds() image.Rectangle {
	return image.Rect(0, 0, zm.width, zm.height)
}

func (zm *DepthZmapper) At(x, y int) color.Color {
	return zm.background
}
