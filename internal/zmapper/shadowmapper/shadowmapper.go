package shadowmapper

import (
	"NBodySim/internal/mathutils/vector"
	"NBodySim/internal/object"
	"NBodySim/internal/zmapper/shadow"
	"sync"
)

type Shadow interface {
	PointInShadow(p vector.Vector3d) bool

	object.ObjectVisitor
}

type ShadowMapper struct {
	resolution int
	cam        *object.Camera
	lights     []object.Light
	shadows    map[int64]Shadow
}

func NewShadowMapper(resolution int) *ShadowMapper {
	return &ShadowMapper{resolution: resolution, lights: make([]object.Light, 0), shadows: make(map[int64]Shadow)}
}

func (sm *ShadowMapper) AddLight(light object.Light) {
	light.Accept(sm)
}

func (sm *ShadowMapper) VisitPointLight(l *object.PointLight) {
	sm.lights = append(sm.lights, l)
	sm.shadows[l.GetId()] = shadow.NewPointShadowMap(sm.resolution, l)
}

func (sm *ShadowMapper) VisitObjectPool(l *object.ObjectPool) {
	wg := sync.WaitGroup{}
	for _, shadow := range sm.shadows {
		wg.Add(1)
		go func() {
			l.Accept(shadow)
			wg.Done()
		}()
	}
	wg.Wait()
}

func (sm *ShadowMapper) VisitPolygonObject(po *object.PolygonObject) {
	for _, shadow := range sm.shadows {
		po.Accept(shadow)
	}
}

func (sm *ShadowMapper) VisitCamera(l *object.Camera) {
	for _, shadow := range sm.shadows {
		l.Accept(shadow)
	}
}

func (sm *ShadowMapper) InShadowBy(id int64, p vector.Vector3d) bool {
	if shadow, ok := sm.shadows[id]; ok {
		return shadow.PointInShadow(p)
	}
	return false
}
