package shadowmapper

import (
	"NBodySim/internal/mathutils/vector"
	"NBodySim/internal/object"
	"NBodySim/internal/zmapper/shadow"
	"sync"
)

type Shadow interface {
	PointInShadow(p vector.Vector3d) bool
	SurfacePointInShadow(p vector.Vector3d, normal vector.Vector3d) bool
	object.ObjectVisitor
}

// Creates a new Shadow
type ShadowMapper struct {
	resolution int
	objs       *object.ObjectPool
}

func NewShadowMapper(resolution int) *ShadowMapper {
	return &ShadowMapper{resolution: resolution, objs: object.NewObjectPool()}
}

func (sm *ShadowMapper) VisitCamera(camera *object.Camera) {

}

func (sm *ShadowMapper) VisitPolygonObject(po *object.PolygonObject) {
	sm.objs.PutObject(po.Clone())
}

func (sm *ShadowMapper) VisitObjectPool(pool *object.ObjectPool) {
	wg := sync.WaitGroup{}
	for _, obj := range pool.GetObjects() {
		wg.Add(1)
		go func(obj object.Object) {
			defer wg.Done()
			obj.Accept(sm)
		}(obj)
	}
	wg.Wait()
}

func (sm *ShadowMapper) VisitPointLight(p *object.PointLight) {

}

func (sm *ShadowMapper) VisitPointLightShadow(p *object.PointLightShadow) {
	shad := shadow.NewPointShadowMap(sm.resolution, p)
	sm.objs.Accept(shad)
	p.SetShadowModel(shad)
}
