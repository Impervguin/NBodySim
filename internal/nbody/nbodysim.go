package nbody

type NBodySim struct {
	nbody  *NBody
	solver NBodySolver
	engine NBodyEngine
	dt     float64
}

func NewNBodySim(nbody *NBody, solver NBodySolver, engine NBodyEngine, dt float64) *NBodySim {
	return &NBodySim{nbody: nbody, solver: solver, engine: engine, dt: dt}
}

func (sim *NBodySim) SolveStep() error {
	uNBody, err := sim.engine.Calculate(sim.nbody, sim.solver.CalculateBody, sim.dt)
	if err != nil {
		return err
	}
	sim.solver.UpdateSelf(uNBody)
	sim.nbody.UpdateSelf(uNBody)
	return nil
}

func (sim *NBodySim) GetNBody() *NBody { return sim.nbody }

func (sim *NBodySim) SetSolver(solver NBodySolver) {
	sim.solver = solver
}
