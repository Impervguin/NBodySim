package shadowdrawer

import (
	"NBodySim/internal/object"
	"NBodySim/internal/zmapper/mapper"
	"NBodySim/internal/zmapper/shadowapproximator"
	"sync"
)

type ParallelShadowPerObjectDrawer struct {
	SimpleShadowObjectDrawer
}

type ParallelShadowPerObjectDrawerFabric struct {
	SimpleShadowObjectDrawerFabric
}

func NewParallelShadowPerObjectDrawerFabric(zf mapper.ZmapperFabric, af shadowapproximator.ShadowDiscreteApproximatorFabric) *ParallelShadowPerObjectDrawerFabric {
	return &ParallelShadowPerObjectDrawerFabric{
		SimpleShadowObjectDrawerFabric: *NewSimpleShadowObjectDrawerFabric(zf, af),
	}
}

func (pd *ParallelShadowPerObjectDrawerFabric) CreatePerObjectDrawer() ShadowObjectDrawer {
	return newParallelShadowPerObjectDrawer(pd.zf, pd.af)
}

func newParallelShadowPerObjectDrawer(zf mapper.ZmapperFabric, af shadowapproximator.ShadowDiscreteApproximatorFabric) *ParallelShadowPerObjectDrawer {
	return &ParallelShadowPerObjectDrawer{
		SimpleShadowObjectDrawer: *newSimpleShadowObjectDrawer(zf, af),
	}
}

func (pd *ParallelShadowPerObjectDrawer) VisitObjectPool(op *object.ObjectPool) {
	pd.startDrawing()
	wait := sync.WaitGroup{}
	for _, obj := range op.GetObjects() {
		wait.Add(1)
		go func() {
			obj.Accept(pd)
			wait.Done()
		}()
	}
	wait.Wait()
	pd.stopDrawing()
}
