package shadow

import (
	"NBodySim/internal/cutter"
	"NBodySim/internal/mathutils"
	"NBodySim/internal/mathutils/vector"
	"NBodySim/internal/object"
	"NBodySim/internal/transform"
	"NBodySim/internal/zmapper/approximator"
	"NBodySim/internal/zmapper/mapper"
	"NBodySim/internal/zmapper/objectdrawer"
	"image/color"
	"math"
)

type ShadowMap struct {
	resolution int
	cam        object.Camera
	cutter     cutter.Cutter
	drawer     objectdrawer.ObjectDrawer
	bias       float64
	toMap      transform.TransformAction
}

const BiasCoeff = 0.005

func NewShadowMap(resolution int, cam object.Camera) *ShadowMap {
	m := &ShadowMap{resolution: resolution,
		cam:    cam,
		bias:   BiasCoeff,
		cutter: cutter.NewSimpleCamCutter(&cam),
		drawer: objectdrawer.NewParallelPerObjectDrawerFabric(
			mapper.NewParallelDepthZmapperFabric(resolution, resolution, color.Black),
			approximator.NewDepthApproximatorFabric(),
		).CreateObjectDrawer(),
	}
	toMap := transform.NewViewportToCanvas(float64(resolution), float64(resolution))
	toMap.ApplyAfter(transform.NewMoveAction(vector.NewVector3d(float64(resolution)/2, float64(resolution)/2, 0)))
	m.toMap = toMap
	return m
}

func (m *ShadowMap) VisitObjectPool(pool *object.ObjectPool) {
	cop := pool.Clone()
	cop.Transform(m.cam.GetViewAction())
	cop.Accept(m.cutter)
	cop.Transform(m.cam.GetPerspectiveTransform())
	cop.Transform(m.toMap)
	cop.Accept(m.drawer)
}

func (m *ShadowMap) VisitPolygonObject(object *object.PolygonObject) {
	cop := object.Clone()
	cop.Transform(m.cam.GetViewAction())
	cop.Accept(m.cutter)
	cop.Transform(m.cam.GetPerspectiveTransform())
	cop.Transform(m.toMap)
	cop.Accept(m.drawer)
}

func (m *ShadowMap) VisitCamera(cam *object.Camera) {
	m.cam.Transform(cam.GetViewAction())
}

func (m *ShadowMap) PointOnMap(p vector.Vector3d) bool {
	m.cam.GetViewAction().ApplyToVector(&p)
	return m.cutter.SeePoint(&p)
}

func (m *ShadowMap) PointInShadow(p vector.Vector3d) bool {
	if !m.PointOnMap(p) {
		return false
	}
	m.cam.GetViewAction().ApplyToVector(&p)
	m.cam.GetPerspectiveTransform().ApplyToVector(&p)
	m.toMap.ApplyToVector(&p)
	// fmt.Println(p)

	ap := approximator.DiscreteFlatPoint{X: mathutils.ToInt(p.X), Y: mathutils.ToInt(p.Y), Z: 0, Color: color.Black}
	m.drawer.SetPointDepth(&ap)
	// fmt.Println(p, ap)

	if (ap.Z) == math.Inf(1) {
		return false
	}
	return ap.Z > p.Z+m.bias
}
