package cutter

import (
	"NBodySim/internal/mathutils/plane"
	"NBodySim/internal/mathutils/vector"
	"NBodySim/internal/object"
)

type SimpleCamCutter struct {
	camera                         *object.Camera
	left, right, bottom, top, back *plane.Plane
	planes                         []*plane.Plane
	distance                       float64
}

func NewSimpleCamCutter(camera *object.Camera) *SimpleCamCutter {
	// calculating camera view normals
	d := camera.GetDistance()
	px, py := camera.GetWidth(), camera.GetHeight()
	right := plane.NewPlane(*vector.NewVector3d(-2*d/px, 0, 1), 0)
	left := plane.NewPlane(*vector.NewVector3d(2*d/px, 0, 1), 0)
	bottom := plane.NewPlane(*vector.NewVector3d(0, 2*d/py, 1), 0)
	top := plane.NewPlane(*vector.NewVector3d(0, -2*d/py, 1), 0)
	back := plane.NewPlane(*vector.NewVector3d(0, 0, 1), d)
	return &SimpleCamCutter{
		camera:   camera,
		left:     left,
		right:    right,
		bottom:   bottom,
		top:      top,
		back:     back,
		planes:   []*plane.Plane{left, right, bottom, top, back},
		distance: d,
	}
}

func (c *SimpleCamCutter) SeePoint(point *vector.Vector3d) bool {
	for _, plane := range c.planes {
		if plane.IsPointBehind(point) {
			return false
		}
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

func (c *SimpleCamCutter) CutPolygon(p *object.Polygon) []*object.Polygon {
	v1, v2, v3 := p.GetVertices()
	res := append(make([]*object.Polygon, 0, 4), object.NewPolygon(v1, v2, v3, p.GetColor()))
	if c.SeePoint(v1) && c.SeePoint(v2) && c.SeePoint(v3) {
		return res
	}

	for _, plane := range c.planes {
		res = c.CutPolygonsAtPlane(res, plane)
	}
	return res
}

func (c *SimpleCamCutter) CutPolygonsAtPlane(p []*object.Polygon, plane *plane.Plane) []*object.Polygon {
	res := make([]*object.Polygon, 0, len(p)*2)
	for _, poly := range p {
		res = append(res, c.CutPolygonAtPlane(poly, plane)...)
	}
	return res
}

func (c *SimpleCamCutter) CutPolygonAtPlane(p *object.Polygon, plane *plane.Plane) []*object.Polygon {
	seen := make([]*vector.Vector3d, 0, 3)
	unseen := make([]*vector.Vector3d, 0, 3)
	v1, v2, v3 := p.GetVertices()
	if plane.IsPointInFront(v1) {
		seen = append(seen, v1)
	} else {
		unseen = append(unseen, v1)
	}
	if plane.IsPointInFront(v2) {
		seen = append(seen, v2)
	} else {
		unseen = append(unseen, v2)
	}
	if plane.IsPointInFront(v3) {
		seen = append(seen, v3)
	} else {
		unseen = append(unseen, v3)
	}

	if len(seen) == 3 {
		return []*object.Polygon{p}
	}
	if len(seen) == 1 {
		vseen := seen[0]
		v1, v2 := unseen[0], unseen[1]
		intersec1, _ := plane.Intersection(vseen, v1)
		intersec2, _ := plane.Intersection(vseen, v2)
		return []*object.Polygon{object.NewPolygon(vseen, intersec1, intersec2, p.GetColor())}
	}
	if len(seen) == 2 {
		vunseen := unseen[0]
		v1, v2 := seen[0], seen[1]
		intersec1, _ := plane.Intersection(vunseen, v1)
		intersec2, _ := plane.Intersection(vunseen, v2)
		return []*object.Polygon{
			object.NewPolygon(v1, v2, intersec1, p.GetColor()),
			object.NewPolygon(v2, intersec1, intersec2, p.GetColor()),
		}
	}
	return make([]*object.Polygon, 0)
}

func (c *SimpleCamCutter) VisitPolygonObject(po *object.PolygonObject) {
	see := c.SeePolygonObject(po)
	polygons := po.GetPolygons()
	po.ResetVertices()
	po.ResetPolygons()
	if !see {
		return
	}
	vertMap := make(map[*vector.Vector3d]struct{}, 2*len(polygons))
	for _, polygon := range polygons {
		if c.SeePolygon(polygon) {
			cutPolygon := c.CutPolygon(polygon)
			for _, cut := range cutPolygon {
				po.AddPolygon(cut)
				v1, v2, v3 := cut.GetVertices()
				vertMap[v1] = struct{}{}
				vertMap[v2] = struct{}{}
				vertMap[v3] = struct{}{}
			}
		}
	}
	for v := range vertMap {
		po.AddVertex(v)
	}
}

func (c *SimpleCamCutter) VisitCamera(cam *object.Camera) {
	// Nothing to do here, just a placeholder
}

func (c *SimpleCamCutter) VisitPointLight(light *object.PointLight) {
	// Nothing to do here, just a placeholder
}
