package simulation2d

import (
	"NBodySim/internal/nbody"
	"image/color"
)

type Body2d struct {
	body   *nbody.Body
	radius float64
	color  color.Color
}

func NewBody2d(body *nbody.Body, radius float64, color color.Color) *Body2d {
	return &Body2d{body: body, radius: radius, color: color}
}
