package simulation

type SimulationVisitor interface {
	VisitSimulation(s *Simulation)
}
