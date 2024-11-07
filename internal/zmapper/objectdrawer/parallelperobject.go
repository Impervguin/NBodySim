package objectdrawer

import (
	"NBodySim/internal/object"
	"NBodySim/internal/zmapper/approximator"
	"NBodySim/internal/zmapper/mapper"
	"sync"
)

type ParallelPerObjectDrawer struct {
	SimpleObjectDrawer
}

type ParallelPerObjectDrawerFabric struct {
	SimpleObjectDrawerFabric
}

func NewParallelPerObjectDrawerFabric(zf mapper.ZmapperFabric, af approximator.DiscreteApproximatorFabric) *ParallelPerObjectDrawerFabric {
	return &ParallelPerObjectDrawerFabric{
		SimpleObjectDrawerFabric: *NewSimpleObjectDrawerFabric(zf, af),
	}
}

func (pd *ParallelPerObjectDrawerFabric) CreatePerObjectDrawer() ObjectDrawer {
	return newParallelPerObjectDrawer(pd.zf, pd.af)
}

func newParallelPerObjectDrawer(zf mapper.ZmapperFabric, af approximator.DiscreteApproximatorFabric) *ParallelPerObjectDrawer {
	return &ParallelPerObjectDrawer{
		SimpleObjectDrawer: *newSimpleObjectDrawer(zf, af),
	}
}

func (pd *ParallelPerObjectDrawer) VisitObjectPool(op *object.ObjectPool) {
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
