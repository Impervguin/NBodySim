package mapper

import (
	"NBodySim/internal/zmapper/approximator"
	"NBodySim/internal/zmapper/buffers"
	"image"
	"image/color"
)

type SimpleZmapper struct {
	dbuf          buffers.DepthBuffer
	sbuf          buffers.ScreenBuffer
	width, height int
}

type SimpleZmapperFabric struct {
	width, height int
	background    color.Color
}

func NewSimpleZmapperFabric(width, height int, background color.Color) *SimpleZmapperFabric {
	return &SimpleZmapperFabric{
		width:      width,
		height:     height,
		background: background,
	}
}

func (f *SimpleZmapperFabric) CreateZmapper() Zmapper {
	return newSimpleZmapper(f.width, f.height, f.background, &buffers.DepthBufferInfFabric{})
}

func newSimpleZmapper(width, height int, background color.Color, df buffers.DepthBufferFabric) *SimpleZmapper {
	return &SimpleZmapper{
		width:  width,
		height: height,
		sbuf:   *buffers.NewScreenBuffer(width, height, background),
		dbuf:   df.CreateDepthBuffer(width, height),
	}
}

func (zm *SimpleZmapper) setPoint(x, y int, z float64, color color.Color) {
	ok, _ := zm.dbuf.PutPoint(x, y, z)
	if ok {
		zm.sbuf.PutPoint(x, y, color)
	}
}

func (zm *SimpleZmapper) DrawChannel(ch <-chan approximator.DiscreteFlatPoint) {
	for dp := range ch {
		zm.setPoint(dp.X, dp.Y, dp.Z, dp.Color)
	}
}

func (zm *SimpleZmapper) GetScreenFunction() buffers.ScreenFunction {
	return func(x, y, w, h int) color.Color {
		return zm.sbuf.GetPoint(x, y)
	}
}

func (zm *SimpleZmapper) Reset() {
	zm.sbuf.Reset()
	zm.dbuf.Reset()
}

func (zm *SimpleZmapper) ColorModel() color.Model {
	return color.RGBAModel
}

func (zm *SimpleZmapper) Bounds() image.Rectangle {
	return image.Rect(0, 0, zm.width, zm.height)
}

func (zm *SimpleZmapper) At(x, y int) color.Color {
	return zm.sbuf.GetPoint(x, y)
}

func (zm *SimpleZmapper) SetPointDepth(p *approximator.DiscreteFlatPoint) {
	p.Z, _ = zm.dbuf.GetDepth(p.X, p.Y)
}
