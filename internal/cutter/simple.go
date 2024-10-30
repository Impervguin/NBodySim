package cutter

import (
	"NBodySim/internal/object"
	"NBodySim/internal/vectormath"
)

type SimpleCamCutter struct {
	camera                   *object.Camera
	left, right, bottom, top vectormath.Vector3d
	distance                 float64
}

func NewSimpleCamCutter(camera *object.Camera) *SimpleCamCutter {
	// calculating camera view normals
	d := camera.GetDistance()
	px, py := camera.GetWidth(), camera.GetHeight()
	right := vectormath.Vector3d{-2 * d / px, 0, 1}
	left := vectormath.Vector3d{2 * d / px, 0, 1}
	bottom := vectormath.Vector3d{0, 2 * d / py, 1}
	top := vectormath.Vector3d{0, -2 * d / py, 1}
	right.Normalize()
	top.Normalize()
	left.Normalize()
	bottom.Normalize()
	return &SimpleCamCutter{
		camera:   camera,
		left:     left,
		right:    right,
		bottom:   bottom,
		top:      top,
		distance: d,
	}
}

func (c *SimpleCamCutter) SeePoint(point *vectormath.Vector3d) bool {
	if point.Z < c.distance {
		return false
	}
	if point.Dot(&c.left) < 0 {
		// fmt.Println(point.Dot(&c.left))
		// fmt.Println(point)
		return false
	}
	if point.Dot(&c.right) < 0 {
		return false
	}
	if point.Dot(&c.bottom) < 0 {
		return false
	}
	if point.Dot(&c.top) < 0 {
		return false
	}
	return true
}

func (c *SimpleCamCutter) SeePolygonObject(po *object.PolygonObject) bool {
	for _, v := range po.GetVertices() {
		if c.SeePoint(v) {
			return true
		}
	}
	return false
}

func (c *SimpleCamCutter) SeePolygon(p *object.Polygon) bool {
	v1, v2, v3 := p.GetVertices()
	if c.SeePoint(v1) {
		return true
	}
	if c.SeePoint(v2) {
		return true
	}
	if c.SeePoint(v3) {
		return true
	}
	return false
}

func (c *SimpleCamCutter) CutPolygon(p *object.Polygon) []object.Polygon {
	res := make([]object.Polygon, 0, 2)
	v1, v2, v3 := p.GetVertices()
	if c.SeePoint(v1) && c.SeePoint(v2) && c.SeePoint(v3) {
		res = append(res, *object.NewPolygon(v1.Copy(), v2.Copy(), v3.Copy(), p.GetColor()))
		return res
	}
	// res = append(res, *p)
	return res
}

func (c *SimpleCamCutter) VisitPolygonObject(po *object.PolygonObject) {
	see := c.SeePolygonObject(po)
	polygons := po.GetPolygons()
	po.ResetVertices()
	po.ResetPolygons()
	if !see {
		return
	}
	for _, polygon := range polygons {
		if c.SeePolygon(polygon) {
			cutPolygon := c.CutPolygon(polygon)
			for _, cut := range cutPolygon {
				po.AddPolygon(&cut)
				v1, v2, v3 := cut.GetVertices()
				po.AddVertex(v1)
				po.AddVertex(v2)
				po.AddVertex(v3)
			}
		}
	}
}

func (c *SimpleCamCutter) VisitCamera(cam *object.Camera) {
	// Nothing to do here, just a placeholder
}
