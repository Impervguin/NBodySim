package shadow

import (
	"NBodySim/internal/mathutils/vector"
	"NBodySim/internal/object"
	"NBodySim/internal/transform"
	"math"
	"sync"
)

type PointShadowMap struct {
	resolution int
	ligth      object.PointLight

	forward ShadowMap
	back    ShadowMap
	left    ShadowMap
	right   ShadowMap
	top     ShadowMap
	bottom  ShadowMap
}

func NewPointShadowMap(resolution int, light *object.PointLight) *PointShadowMap {
	pos := light.GetCenter()
	pshm := &PointShadowMap{
		resolution: resolution,
		ligth:      *light,
		forward:    *NewShadowMap(resolution, *object.NewCamera(pos, *vector.NewVector3d(0, 0, 1), *vector.NewVector3d(0, 1, 0), 1, 1, 1)),
	}
	backcam := pshm.forward.cam.Clone().(*object.Camera)
	backcam.Transform(transform.NewRotateActionCenter(&pos, vector.NewVector3d(0, math.Pi, 0)))
	pshm.back = *NewShadowMap(resolution, *backcam)

	leftcam := pshm.forward.cam.Clone().(*object.Camera)
	leftcam.Transform(transform.NewRotateActionCenter(&pos, vector.NewVector3d(0, math.Pi/2, 0)))
	pshm.left = *NewShadowMap(resolution, *leftcam)

	rightcam := pshm.forward.cam.Clone().(*object.Camera)
	rightcam.Transform(transform.NewRotateActionCenter(&pos, vector.NewVector3d(0, -math.Pi/2, 0)))
	pshm.right = *NewShadowMap(resolution, *rightcam)

	topcam := pshm.forward.cam.Clone().(*object.Camera)
	topcam.Transform(transform.NewRotateActionCenter(&pos, vector.NewVector3d(math.Pi/2, 0, 0)))
	pshm.top = *NewShadowMap(resolution, *topcam)

	bottomcam := pshm.forward.cam.Clone().(*object.Camera)
	bottomcam.Transform(transform.NewRotateActionCenter(&pos, vector.NewVector3d(-math.Pi/2, 0, 0)))
	pshm.bottom = *NewShadowMap(resolution, *bottomcam)

	return pshm
}

func (p *PointShadowMap) VisitObjectPool(pool *object.ObjectPool) {
	wg := &sync.WaitGroup{}
	maps := []*ShadowMap{&p.forward, &p.back, &p.left, &p.right, &p.top, &p.bottom}
	for _, m := range maps {
		wg.Add(1)
		go func() {
			pool.Accept(m)
			wg.Done()
		}()
	}
	wg.Wait()
}

func (p *PointShadowMap) VisitPolygonObject(polygon *object.PolygonObject) {
	p.forward.VisitPolygonObject(polygon)
	p.back.VisitPolygonObject(polygon)
	p.left.VisitPolygonObject(polygon)
	p.right.VisitPolygonObject(polygon)
	p.top.VisitPolygonObject(polygon)
	p.bottom.VisitPolygonObject(polygon)
}

func (p *PointShadowMap) VisitCamera(cam *object.Camera) {
	p.forward.VisitCamera(cam)
	p.back.VisitCamera(cam)
	p.left.VisitCamera(cam)
	p.right.VisitCamera(cam)
	p.top.VisitCamera(cam)
	p.bottom.VisitCamera(cam)
}

func (p *PointShadowMap) PointInShadow(v vector.Vector3d) bool {
	if p.forward.PointInShadow(v) ||
		p.back.PointInShadow(v) ||
		p.left.PointInShadow(v) ||
		p.right.PointInShadow(v) ||
		p.top.PointInShadow(v) ||
		p.bottom.PointInShadow(v) {
		return true
	}
	return false
}
