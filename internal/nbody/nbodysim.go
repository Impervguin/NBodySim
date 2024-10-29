package nbody

import "fmt"

type NBody struct {
	nbody  map[int64]Body
	solver NBodySolver
	engine NBodyEngine
}

func NewNBody(solver NBodySolver, engine NBodyEngine) *NBody {
	return &NBody{nbody: make(map[int64]Body), solver: solver, engine: engine}
}

func (sim *NBody) ResetSolver() {
	sim.solver.Reset()
	bodies := make([]PhysBody, 0, len(sim.nbody))
	for _, body := range sim.nbody {
		bodies = append(bodies, *FromBody(body))
	}
	sim.solver.UpdateSelf(bodies)
}

func (sim *NBody) AddBody(body Body) error {
	if _, ok := sim.nbody[body.GetId()]; ok {
		return fmt.Errorf("body with id %d already exists", body.GetId())
	}
	sim.nbody[body.GetId()] = body
	sim.ResetSolver()
	return nil
}

func (sim *NBody) RemoveBody(id int64) error {
	if _, ok := sim.nbody[id]; !ok {
		return fmt.Errorf("body with id %d does not exist", id)
	}
	delete(sim.nbody, id)
	sim.ResetSolver()
	return nil
}

func (sim *NBody) UpdateBody(body Body) error {
	if _, ok := sim.nbody[body.GetId()]; !ok {
		return fmt.Errorf("body with id %d does not exist", body.GetId())
	}
	sim.nbody[body.GetId()] = body
	return nil
}

func (sim *NBody) GetBody(id int64) (Body, bool) {
	b, ok := sim.nbody[id]
	return b, ok
}

func (sim *NBody) Clone() *NBody {
	newNBody := make(map[int64]Body, len(sim.nbody))
	for id, body := range sim.nbody {
		newNBody[id] = body.Clone()
	}
	return &NBody{nbody: newNBody, solver: sim.solver, engine: sim.engine}
}

func (sim *NBody) SolveStep(dt float64) error {
	if sim.solver == nil || sim.engine == nil {
		return fmt.Errorf("solver and engine must be set")
	}
	bodies := make([]PhysBody, 0, len(sim.nbody))
	for _, body := range sim.nbody {
		bodies = append(bodies, *FromBody(body))
	}
	err := sim.engine.Calculate(bodies, sim.solver, dt)
	if err != nil {
		return err
	}
	sim.solver.UpdateSelf(bodies)
	for _, phbody := range bodies {
		sim.nbody[phbody.Id].SetVelocity(phbody.Velocity)
		sim.nbody[phbody.Id].SetPosition(phbody.Position)
	}
	return nil
}

func (sim *NBody) SolveSteps(times int, dt float64) error {
	if sim.solver == nil || sim.engine == nil {
		return fmt.Errorf("solver and engine must be set")
	}
	bodies := make([]PhysBody, 0, len(sim.nbody))
	for _, body := range sim.nbody {
		bodies = append(bodies, *FromBody(body))
	}
	for i := 0; i < times; i++ {
		err := sim.engine.Calculate(bodies, sim.solver, dt)
		if err != nil {
			return err
		}
		sim.solver.UpdateSelf(bodies)
	}
	for _, phbody := range bodies {
		sim.nbody[phbody.Id].SetVelocity(phbody.Velocity)
		sim.nbody[phbody.Id].SetPosition(phbody.Position)
	}
	return nil
}

func (sim *NBody) SetSolver(solver NBodySolver) {
	sim.solver = solver
	sim.ResetSolver()
}

func (sim *NBody) SetEngine(engine NBodyEngine) {
	sim.engine = engine
}
