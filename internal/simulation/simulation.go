package simulation

import (
	"NBodySim/internal/nbody"
	"NBodySim/internal/object"
	"NBodySim/internal/vectormath"
	"fmt"
)

type Simulation struct {
	objects object.ObjectPool
	nbody   nbody.NBody
	sim     nbody.NBodySim
	bodyObj map[int64]int64
	camera  *object.Camera
	lights  []*object.Object
}

func NewSimulation() *Simulation {
	sim := &Simulation{}
	sim.init()
	return sim
}

func (s *Simulation) init() {
	s.nbody = *nbody.NewNBody(make([]*nbody.Body, 0))
	s.sim = *nbody.NewNBodySim(&s.nbody, nbody.NewEulerSolver(&s.nbody), nbody.NewIterativeNbodyEngine(), 0.01)
	s.objects = *object.NewObjectPool()
	s.bodyObj = make(map[int64]int64)
	s.camera = object.NewCamera(*vectormath.NewVector3d(0, 0, 0), *vectormath.NewVector3d(0, 0, 1), *vectormath.NewVector3d(0, 1, 0), 1, 1, 1)
}

func (s *Simulation) SetCamera(cam *object.Camera) error {
	if cam == nil {
		return fmt.Errorf("camera cannot be nil")
	}
	s.camera = cam
	return nil
}

func (s *Simulation) GetCamera() *object.Camera {
	return s.camera
}

func (s *Simulation) SetSolver(solver nbody.NBodySolver) {
	s.sim.SetSolver(solver)
}

func (s *Simulation) SetEngine(engine nbody.NBodyEngine) {
	s.sim.SetEngine(engine)
}

func (s *Simulation) SetDt(dt float64) {
	s.sim.SetDt(dt)
}

func (s *Simulation) Accept(visitor SimulationVisitor) {
	visitor.VisitSimulation(s)
}
func (s *Simulation) GetObjectsClone() *object.ObjectPool {
	return s.objects.Clone()
}

func (s *Simulation) AddObject(obj *SimulationObject) error {
	if obj.body == nil {
		return fmt.Errorf("body not set for object")
	}
	if _, ok := s.bodyObj[obj.body.Id]; ok {
		return fmt.Errorf("body with id %d already exists", obj.body.Id)
	}
	if _, ok := s.objects.GetObject(obj.obj.GetId()); ok {
		return fmt.Errorf("object with id %d already exists", obj.obj.GetId())
	}
	_, err := s.nbody.AddBody(obj.body)
	if err != nil {
		return err
	}

	s.objects.PutObject(obj.obj)
	s.bodyObj[obj.body.Id] = obj.obj.GetId()
	return nil
}
