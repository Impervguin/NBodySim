package simulation

import (
	"NBodySim/internal/mathutils/vector"
	"NBodySim/internal/nbody"
	"NBodySim/internal/object"
	"fmt"
)

type Simulation struct {
	objects          object.ObjectPool
	imaginaryObjects object.ObjectPool
	nbody            nbody.NBody
	camera           *object.Camera
	lights           object.LightPool
	dt               float64
	timeMoment       float64
}

func NewSimulation() *Simulation {
	sim := &Simulation{}
	sim.init()
	return sim
}

func (s *Simulation) init() {
	s.nbody = *nbody.NewNBody(nbody.NewEulerSolver(), nbody.NewIterativeNbodyEngine())
	s.dt = 0.01
	s.timeMoment = 0
	s.objects = *object.NewObjectPool()
	s.imaginaryObjects = *object.NewObjectPool()
	s.lights = *object.NewLightPool()
	s.camera = object.NewCamera(*vector.NewVector3d(0, 0, 0), *vector.NewVector3d(0, 0, 1), *vector.NewVector3d(0, 1, 0), 1, 1, 1)
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
	s.nbody.SetSolver(solver)
}

func (s *Simulation) SetEngine(engine nbody.NBodyEngine) {
	s.nbody.SetEngine(engine)
}

func (s *Simulation) SetDt(dt float64) {
	s.dt = dt
}

func (s *Simulation) Update() {
	s.timeMoment += s.dt
	s.nbody.SolveStep(s.dt)
	body, _ := s.nbody.GetBody(2)
	fmt.Println("Id:", 2, body.GetPosition())
}

func (s *Simulation) UpdateFor(t float64) {
	times := int(t / s.dt)
	s.timeMoment += float64(times) * s.dt
	s.nbody.SolveSteps(times, s.dt)
}

func (s *Simulation) Accept(visitor SimulationVisitor) {
	visitor.VisitSimulation(s)
}
func (s *Simulation) GetObjectsClone() *object.ObjectPool {
	return s.objects.Clone()
}

func (s *Simulation) AddObject(obj object.Object, velocity vector.Vector3d, mass float64) error {
	phys := NewPhysicalBody(obj, velocity, mass)
	if _, ok := s.objects.GetObject(obj.GetId()); ok {
		return fmt.Errorf("object with id %d already exists", obj.GetId())
	}
	if _, ok := s.nbody.GetBody(obj.GetId()); ok {
		return fmt.Errorf("body with id %d already exists", obj.GetId())
	}
	err := s.nbody.AddBody(phys)
	if err != nil {
		return nil
	}
	s.objects.PutObject(obj)

	return nil
}

func (s *Simulation) AddLight(light object.Light) error {
	if _, ok := s.lights.GetLight(light.GetId()); ok {
		fmt.Println(light.GetId())
		return fmt.Errorf("light with id %d already exists", light.GetId())
	}
	s.lights.PutLight(light)
	return nil
}

func (s *Simulation) GetLightsClone() *object.LightPool {
	return s.lights.Clone()
}

func (s *Simulation) RemoveObject(id int64) error {
	_, ok := s.objects.GetObject(id)
	if !ok {
		return fmt.Errorf("object with id %d does not exist", id)
	}
	err := s.nbody.RemoveBody(id)
	if err != nil {
		return err
	}
	s.objects.RemoveObject(id)
	return nil
}

func (s *Simulation) AddImaginaryObject(obj object.Object) error {
	if _, ok := s.imaginaryObjects.GetObject(obj.GetId()); ok {
		return fmt.Errorf("object with id %d already exists", obj.GetId())
	}
	s.imaginaryObjects.PutObject(obj)
	return nil
}

func (s *Simulation) GetImaginaryObjectsClone() *object.ObjectPool {
	return s.imaginaryObjects.Clone()
}

func (s *Simulation) RemoveImaginaryObject(objId int64) error {
	_, ok := s.imaginaryObjects.GetObject(objId)
	if !ok {
		return fmt.Errorf("object with id %d does not exist", objId)
	}
	s.imaginaryObjects.RemoveObject(objId)
	return nil
}
