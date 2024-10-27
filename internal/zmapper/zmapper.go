package zmapper

import (
	"NBodySim/internal/mathutils"
	"NBodySim/internal/object"
	"fmt"
	"image/color"
)

type ScreenFunction func(x, y, w, h int) color.Color

type Zmapper struct {
	dbuf          DepthBuffer
	sbuf          ScreenBuffer
	width, height int
}

func NewZmapper(width, height int, background color.Color, df DepthBufferFabric) *Zmapper {
	return &Zmapper{
		width:  width,
		height: height,
		sbuf:   *NewScreenBuffer(width, height, background),
		dbuf:   df.CreateDepthBuffer(width, height),
	}
}

type zmapPoint struct {
	X, Y  int
	Z     float64
	Color color.Color
}

func (zm *Zmapper) polygonGenerator(p *object.Polygon) <-chan zmapPoint {
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

func (zm *Zmapper) processPolygon(p *object.Polygon) {
	ch := zm.polygonGenerator(p)
	for zmp := range ch {
		ok, err := zm.dbuf.PutPoint(zmp.X, zmp.Y, zmp.Z)
		if err != nil {
			fmt.Println("leee, error writing to depth buffer:", err)
		}
		if ok {
			zm.sbuf.PutPoint(zmp.X, zmp.Y, zmp.Color)
		}
	}
}

func (zm *Zmapper) VisitPolygonObject(po *object.PolygonObject) {
	for _, p := range po.GetPolygons() {
		zm.processPolygon(p)
	}
}

func (zm *Zmapper) VisitCamera(cam *object.Camera) {
	// Nothing to do here
}

func (zm *Zmapper) GetScreenFunction() ScreenFunction {
	return func(x, y, w, h int) color.Color {
		return zm.sbuf.GetPoint(x, y)
	}
}

func (zm *Zmapper) Reset() {
	zm.sbuf.Reset()
	zm.dbuf.Reset()
}
