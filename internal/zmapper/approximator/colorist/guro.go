package colorist

import (
	"NBodySim/internal/mathutils"
	"NBodySim/internal/mathutils/vector"
	"NBodySim/internal/object"
	"image/color"
)

type GuroColorModel struct {
	C1 color.RGBA
	C2 color.RGBA
	C3 color.RGBA
}

type GuroColorist struct {
	view   vector.Vector3d
	lights []object.Light
}

func NewGuroColorist() *GuroColorist {
	return &GuroColorist{view: *vector.NewVector3d(0, 0, 0), lights: make([]object.Light, 0)}
}

func (c *GuroColorist) VisitPolygonObject(po *object.PolygonObject) {
	for _, p := range po.GetPolygons() {
		c.processPolygon(p)
	}
}

func (c *GuroColorist) processPolygon(p *object.Polygon) {
	normal := p.GetNormal()
	pcolor := p.GetColor()
	color := GuroColorModel{
		C1: color.RGBA{0, 0, 0, 255},
		C2: color.RGBA{0, 0, 0, 255},
		C3: color.RGBA{0, 0, 0, 255},
	}
	for _, l := range c.lights {
		v1, v2, v3 := p.GetVertices()
		l1color := l.CalculateLightContribution(*v1, c.view, *normal, pcolor)
		color.C1 = mathutils.AddRGBA(color.C1, l1color)
		l2color := l.CalculateLightContribution(*v2, c.view, *normal, pcolor)
		color.C2 = mathutils.AddRGBA(color.C2, l2color)
		l3color := l.CalculateLightContribution(*v3, c.view, *normal, pcolor)
		color.C3 = mathutils.AddRGBA(color.C3, l3color)
	}
	p.SetColorModel(&color)
}

func (c *GuroColorist) VisitPointLight(light *object.PointLight) {
	c.lights = append(c.lights, light)
}

func (c *GuroColorist) VisitCamera(cam *object.Camera) {
	c.view = cam.GetCenter()
}
