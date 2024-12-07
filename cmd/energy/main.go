package main

import (
	"NBodySim/internal/mathutils/vector"
	"NBodySim/internal/nbody"
	"fmt"
	"os"
)

const dt = 0.01
const maxTime = 1000.
const outFile = "energy.log"

func main() {
	nb := nbody.NewNBody(nbody.NewEulerSolver(), nbody.NewIterativeNbodyEngine())

	body1 := nbody.NewOnlyBody(*vector.NewVector3d(0, 0, 40), *vector.NewVector3d(0.2, 0, 0), 100000000000)
	body2 := nbody.NewOnlyBody(*vector.NewVector3d(40, 0, 0), *vector.NewVector3d(0, 0, -0.2), 100000000000)
	body3 := nbody.NewOnlyBody(*vector.NewVector3d(0, 0, -40), *vector.NewVector3d(-0.2, 0, 0), 100000000000)
	body4 := nbody.NewOnlyBody(*vector.NewVector3d(-40, 0, 0), *vector.NewVector3d(0, 0, 0.2), 100000000000)
	nb.AddBody(body1)
	nb.AddBody(body2)
	nb.AddBody(body3)
	nb.AddBody(body4)

	envis := nbody.NewEnergyVisitor()
	envis.VisitNBody(nb)
	initEnergy := envis.GetTotalEnergy()
	f, err := os.Create(outFile)
	if err != nil {
		panic(err)
	}
	// fmt.Printf("Initial total energy: %.2f\n", initEnergy)
	for t := dt; t < maxTime; t += dt {
		nb.SolveStep(dt)
		envis = nbody.NewEnergyVisitor()
		envis.VisitNBody(nb)
		energy := envis.GetTotalEnergy()
		diff := energy - initEnergy
		diffPercent := diff / initEnergy * 100
		f.WriteString(fmt.Sprintf("%.2f %.2f %.2f %.2f\n", t, initEnergy, energy, diffPercent))
	}
	f.Close()
}
