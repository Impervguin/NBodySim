package mathutils

import (
	"math"
)

func IAbs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}

func ToInt(f float64) int {
	return int(math.Round(f))
}

func LinearXIntInterpolation(x1, y1, x2, y2 int) [][]int {
	if x1 == x2 {
		return [][]int{{x1, y1}}
	}
	dx := x2 - x1
	dy := y2 - y1
	res := make([][]int, IAbs(dx)+1)
	stepy := dy / dx
	for x := x1; x <= x2; x++ {
		y := int(math.Round(float64(y1) + float64(stepy)*(float64(x)-float64(x1))))
		res[x-x1] = []int{x, y}
	}
	return res
}

func LinearYIntInterpolation(x1, y1, x2, y2 int) [][]int {
	if y1 == y2 {
		return [][]int{{x1, y1}}
	}
	if y1 > y2 {
		x1, x2 = x2, x1
		y1, y2 = y2, y1
	}
	dx := x2 - x1
	dy := y2 - y1
	res := make([][]int, IAbs(dy)+1)
	stepx := float64(dx) / float64(dy)
	x := float64(x1)
	for y := y1; y <= y2; y++ {
		res[y-y1] = []int{ToInt(x), y}
		x += stepx
	}

	return res
}

func LinearXInterpolation(x1 int, y1 float64, x2 int, y2 float64) ([]int, []float64) {
	if x1 == x2 {
		return []int{x1}, []float64{y1}
	}
	if x1 > x2 {
		y1, y2 = y2, y1
		x1, x2 = x2, x1
	}
	dx := x2 - x1
	dy := y2 - y1
	resx := make([]int, IAbs(dx)+1)
	resy := make([]float64, IAbs(dx)+1)
	stepy := dy / float64(dx)
	y := y1
	for x := x1; x <= x2; x++ {
		resx[x-x1] = x
		resy[x-x1] = y
		y += stepy
	}
	return resx, resy
}

func LinearIntInterpolation(x1, y1, x2, y2 int) [][]int {
	if x1 == x2 && y1 == y2 {
		return [][]int{{x1, y1}}
	}
	dx := x2 - x1
	dy := y2 - y1
	if IAbs(dx) > IAbs(dy) {
		return LinearXIntInterpolation(x1, y1, x2, y2)
	} else {
		return LinearYIntInterpolation(x1, y1, x2, y2)
	}
}

type DepthPoint struct {
	X, Y int
	Z    float64
}
