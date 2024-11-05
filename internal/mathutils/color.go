package mathutils

import (
	"image/color"
)

func ToRGBA(c color.Color) color.RGBA {
	r, g, b, a := c.RGBA()
	return color.RGBA{R: uint8(r), G: uint8(g), B: uint8(b), A: uint8(a)}
}

func AddRGBA(c1, c2 color.Color) color.RGBA {
	r1, g1, b1, a1 := c1.RGBA()
	r2, g2, b2, a2 := c2.RGBA()
	return color.RGBA{
		R: min(uint8((r1+r2)>>8), 255),
		G: min(uint8((g1+g2)>>8), 255),
		B: min(uint8((b1+b2)>>8), 255),
		A: min(uint8((a1+a2)>>8), 255),
	}
}
