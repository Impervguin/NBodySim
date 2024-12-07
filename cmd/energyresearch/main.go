package main

import (
	"NBodySim/internal/mathutils/vector"
	"NBodySim/internal/nbody"
	"fmt"
	"math"
)

type ResearchUnit struct {
	TimeMax    float64
	Dt         float64
	MinPercent float64
	MaxPercent float64
	SumPercent float64
	Count      int
	Nb         *nbody.NBody
}

func NewResearchUnit(tm, dt float64, nb *nbody.NBody) *ResearchUnit {
	return &ResearchUnit{
		TimeMax:    tm,
		Dt:         dt,
		MinPercent: math.Inf(1),
		MaxPercent: math.Inf(-1),
		SumPercent: 0.,
		Count:      0,
		Nb:         nb,
	}
}

func (ru *ResearchUnit) DoResearch() {
	envis := nbody.NewEnergyVisitor()
	ru.Nb.Accept(envis)
	initEnergy := envis.GetTotalEnergy()
	for t := ru.Dt; t < ru.TimeMax; t += ru.Dt {
		ru.Nb.SolveStep(ru.Dt)
		envis = nbody.NewEnergyVisitor()
		envis.VisitNBody(ru.Nb)
		energy := envis.GetTotalEnergy()
		diff := energy - initEnergy
		diffPercent := math.Abs(diff / initEnergy * 100)
		ru.SumPercent += diffPercent
		ru.Count++
		ru.MinPercent = math.Min(ru.MinPercent, diffPercent)
		ru.MaxPercent = math.Max(ru.MaxPercent, diffPercent)
	}
}

func (ru *ResearchUnit) LatexTableString() string {
	return fmt.Sprintf("%.0f & %.4g & %.4g & %.4g & %.4g \\\\ \\hline", ru.TimeMax, ru.Dt, ru.MinPercent, ru.MaxPercent, ru.SumPercent/float64(ru.Count))
}

func Research1() *ResearchUnit {
	nb := nbody.NewNBody(nbody.NewEulerSolver(), nbody.NewIterativeNbodyEngine())

	body1 := nbody.NewOnlyBody(*vector.NewVector3d(0, 0, 50), *vector.NewVector3d(0.2, 0, 0), 100000000000)
	body2 := nbody.NewOnlyBody(*vector.NewVector3d(50, 0, 0), *vector.NewVector3d(0, 0, -0.2), 100000000000)
	body3 := nbody.NewOnlyBody(*vector.NewVector3d(0, 0, -50), *vector.NewVector3d(-0.2, 0, 0), 100000000000)
	body4 := nbody.NewOnlyBody(*vector.NewVector3d(-50, 0, 0), *vector.NewVector3d(0, 0, 0.2), 100000000000)
	nb.AddBody(body1)
	nb.AddBody(body2)
	nb.AddBody(body3)
	nb.AddBody(body4)
	return NewResearchUnit(90, 0.01, nb)
}

func Research2() *ResearchUnit {
	nb := nbody.NewNBody(nbody.NewEulerSolver(), nbody.NewIterativeNbodyEngine())

	body1 := nbody.NewOnlyBody(*vector.NewVector3d(0, 0, 50), *vector.NewVector3d(0.2, 0, 0), 100000000000)
	body2 := nbody.NewOnlyBody(*vector.NewVector3d(50, 0, 0), *vector.NewVector3d(0, 0, -0.2), 100000000000)
	body3 := nbody.NewOnlyBody(*vector.NewVector3d(0, 0, -50), *vector.NewVector3d(-0.2, 0, 0), 100000000000)
	body4 := nbody.NewOnlyBody(*vector.NewVector3d(-50, 0, 0), *vector.NewVector3d(0, 0, 0.2), 100000000000)
	nb.AddBody(body1)
	nb.AddBody(body2)
	nb.AddBody(body3)
	nb.AddBody(body4)
	return NewResearchUnit(90, 0.001, nb)
}

func Research3() *ResearchUnit {
	nb := nbody.NewNBody(nbody.NewEulerSolver(), nbody.NewIterativeNbodyEngine())

	body1 := nbody.NewOnlyBody(*vector.NewVector3d(0, 0, 50), *vector.NewVector3d(0.2, 0, 0), 100000000000)
	body2 := nbody.NewOnlyBody(*vector.NewVector3d(50, 0, 0), *vector.NewVector3d(0, 0, -0.2), 100000000000)
	body3 := nbody.NewOnlyBody(*vector.NewVector3d(0, 0, -50), *vector.NewVector3d(-0.2, 0, 0), 100000000000)
	body4 := nbody.NewOnlyBody(*vector.NewVector3d(-50, 0, 0), *vector.NewVector3d(0, 0, 0.2), 100000000000)
	nb.AddBody(body1)
	nb.AddBody(body2)
	nb.AddBody(body3)
	nb.AddBody(body4)
	return NewResearchUnit(90, 0.0001, nb)
}

func main() {
	rus := []*ResearchUnit{
		Research1(),
		Research2(),
		Research3(),
	}
	for _, ru := range rus {
		ru.DoResearch()
		fmt.Printf("%s\n", ru.LatexTableString())
	}
}
