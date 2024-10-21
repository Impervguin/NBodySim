package main

import (
	"NBodySim/internal/transform"
	"NBodySim/internal/vectormath"
	"fmt"
	"math"
)

func main() {
	p1 := vectormath.NewVector3d(1, 0.5, 0.3)
	action := transform.NewRotateAction(vectormath.NewVector3d(math.Pi/4, 0, 0))
	// action := transform.NewScaleActionCenter(vectormath.NewVector3d(1, 1, 0), vectormath.NewVector3d(5, 2, 3))
	action.ApplyToVector(p1)
	fmt.Println(p1.X, p1.Y, p1.Z)
}
