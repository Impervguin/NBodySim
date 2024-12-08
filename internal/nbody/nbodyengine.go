package nbody

type NBodyEngine interface {
	Calculate(bodies []PhysBody, solver NBodySolver, dt float64) error
}

type IterativeNbodyEngine struct{}

func NewIterativeNbodyEngine() *IterativeNbodyEngine {
	return &IterativeNbodyEngine{}
}

func (e *IterativeNbodyEngine) Calculate(bodies []PhysBody, solver NBodySolver, dt float64) error {
	for i := range bodies {
		_, err := solver.CalculateBody(&bodies[i], dt)
		if err != nil {
			return err
		}
	}
	return nil
}
