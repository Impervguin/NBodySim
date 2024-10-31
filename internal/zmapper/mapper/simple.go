package mapper

import (
	"NBodySim/internal/mathutils"
	"NBodySim/internal/object"
	"NBodySim/internal/zmapper/buffers"
	"image"
	"image/color"
)

type SimpleZmapper struct {
	dbuf          buffers.DepthBuffer
	sbuf          buffers.ScreenBuffer
	width, height int
}

type SimpleZmapperFabric struct{}

func NewSimpleZmapperFabric() *SimpleZmapperFabric {
	return &SimpleZmapperFabric{}
}

func (f *SimpleZmapperFabric) CreateZmapper(width, height int, background color.Color) Zmapper {
	return newSimpleZmapper(width, height, background, &buffers.DepthBufferInfFabric{})
}

func newSimpleZmapper(width, height int, background color.Color, df buffers.DepthBufferFabric) *SimpleZmapper {
	return &SimpleZmapper{
		width:  width,
		height: height,
		sbuf:   *buffers.NewScreenBuffer(width, height, background),
		dbuf:   df.CreateDepthBuffer(width, height),
	}
}

type zmapPoint struct {
	X, Y  int
	Z     float64
	Color color.Color
}

func (zm *SimpleZmapper) polygonGenerator(p *object.Polygon) <-chan zmapPoint {
	ch := make(chan zmapPoint)
	go func() {

		p1, p2, p3 := p.GetVertices()

		color := p.GetColor()
		if p1.Y < p2.Y {
			p1, p2 = p2, p1
		}
		if p1.Y < p3.Y {
			p1, p3 = p3, p1
		}
		if p2.Y < p3.Y {
			p2, p3 = p3, p2
		}
		// fmt.Println(p1.Y, p2.Y, p3.Y)
		p12 := mathutils.LinearYIntInterpolation(mathutils.ToInt(p1.X), mathutils.ToInt(p1.Y), mathutils.ToInt(p2.X), mathutils.ToInt(p2.Y))
		p13 := mathutils.LinearYIntInterpolation(mathutils.ToInt(p1.X), mathutils.ToInt(p1.Y), mathutils.ToInt(p3.X), mathutils.ToInt(p3.Y))
		p23 := mathutils.LinearYIntInterpolation(mathutils.ToInt(p2.X), mathutils.ToInt(p2.Y), mathutils.ToInt(p3.X), mathutils.ToInt(p3.Y))

		p23 = p23[:len(p23)-1]
		p123 := append(p23, p12...)

		_, z12 := mathutils.LinearXInterpolation(mathutils.ToInt(p1.Y), p1.Z, mathutils.ToInt(p2.Y), p2.Z)
		_, z13 := mathutils.LinearXInterpolation(mathutils.ToInt(p1.Y), p1.Z, mathutils.ToInt(p3.Y), p3.Z)
		_, z23 := mathutils.LinearXInterpolation(mathutils.ToInt(p2.Y), p2.Z, mathutils.ToInt(p3.Y), p3.Z)

		z23 = z23[:len(z23)-1]
		z123 := append(z23, z12...)

		med := len(p123) / 2
		if p123[med][0] < p13[med][0] {
			for i := range p123 {
				z := z123[i]
				dz := (z13[i] - z123[i]) / float64(p13[i][0]-p123[i][0])
				for x := p123[i][0]; x <= p13[i][0]; x++ {
					ch <- zmapPoint{X: x, Y: p123[i][1], Z: z, Color: color}
					z += dz
				}
			}
		} else {
			for i := range p123 {
				z := z13[i]
				dz := (z123[i] - z13[i]) / float64(p123[i][0]-p13[i][0])
				for x := p13[i][0]; x <= p123[i][0]; x++ {
					ch <- zmapPoint{X: x, Y: p13[i][1], Z: z, Color: color}
					z += dz
				}
			}
		}
		close(ch)
	}()
	return ch
}

func (zm *SimpleZmapper) processPolygon(p *object.Polygon) {
	ch := zm.polygonGenerator(p)
	for zmp := range ch {
		zm.setPoint(zmp.X, zmp.Y, zmp.Z, zmp.Color)
	}
}

func (zm *SimpleZmapper) setPoint(x, y int, z float64, color color.Color) {
	ok, _ := zm.dbuf.PutPoint(x, y, z)
	if ok {
		zm.sbuf.PutPoint(x, y, color)
	}
}

func (zm *SimpleZmapper) VisitPolygonObject(po *object.PolygonObject) {
	for _, p := range po.GetPolygons() {
		zm.processPolygon(p)
	}
}

func (zm *SimpleZmapper) VisitCamera(cam *object.Camera) {
	// Nothing to do here
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
