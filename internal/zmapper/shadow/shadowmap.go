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
	cutter     *cutter.SimpleCamCutter
	bcutter    *cutter.BackwardsCutter
	drawer     objectdrawer.ObjectDrawer
	toMap      transform.TransformAction
}

const SimpleBias = 0.0001
const MinBias = 0.0001
const MaxBias = 0.001

func NewShadowMap(resolution int, cam object.Camera) *ShadowMap {
	m := &ShadowMap{resolution: resolution,
		cam:     cam,
		cutter:  cutter.NewSimpleCamCutter(&cam),
		bcutter: cutter.NewBackwardsCutter(&cam),
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
	// cop.Accept(m.bcutter)
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

	ap := approximator.DiscreteFlatPoint{X: mathutils.ToInt(p.X), Y: mathutils.ToInt(p.Y), Z: 0, Color: color.Black}
	m.drawer.SetPointDepth(&ap)

	if (ap.Z) == math.Inf(1) {
		return false
	}
	return ap.Z > p.Z+SimpleBias
}

func (m *ShadowMap) SurfacePointInShadow(p vector.Vector3d, normal vector.Vector3d) bool {
	if !m.PointOnMap(p) {
		return false
	}
	lightPos := m.cam.GetCenter()
	lightDir := vector.SubtractVectors(&p, &lightPos)
	lightDir.Normalize()
	biasCoeff := math.Abs(lightDir.Dot(&normal))
	m.cam.GetViewAction().ApplyToVector(&p)
	m.cam.GetPerspectiveTransform().ApplyToVector(&p)
	m.toMap.ApplyToVector(&p)

	ap := approximator.DiscreteFlatPoint{X: mathutils.ToInt(p.X), Y: mathutils.ToInt(p.Y), Z: 0, Color: color.Black}
	m.drawer.SetPointDepth(&ap)

	if (ap.Z) == math.Inf(1) {
		return false
	}

	bias := math.Max(MaxBias*(1-biasCoeff), MinBias)
	return ap.Z > p.Z+bias
}
