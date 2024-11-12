package mathutils

import (
	"image/color"
)

func MultRGBA64(c color.Color, coeff float64) color.RGBA64 {
	r, g, b, _ := c.RGBA()
	return color.RGBA64{
		R: uint16(min(float64(r)*coeff, 65535)),
		G: uint16(min(float64(g)*coeff, 65535)),
		B: uint16(min(float64(b)*coeff, 65535)),
	}
}

func ToRGBA(c color.Color) color.RGBA {
	r, g, b, a := c.RGBA()
	return color.RGBA{R: uint8(r >> 8), G: uint8(g >> 8), B: uint8(b >> 8), A: uint8(a >> 8)}
}

func AddRGBA(c1, c2 color.Color) color.RGBA {
	r1, g1, b1, a1 := c1.RGBA()
	r2, g2, b2, a2 := c2.RGBA()
	// fmt.Println(r1, r2)
	return color.RGBA{
		R: uint8(min((r1+r2)>>8, 255)),
		G: uint8(min((g1+g2)>>8, 255)),
		B: uint8(min((b1+b2)>>8, 255)),
		A: uint8(min((a1+a2)>>8, 255)),
	}
}

func AddRGBA64(c1, c2 color.Color) color.RGBA64 {
	r1, g1, b1, a1 := c1.RGBA()
	r2, g2, b2, a2 := c2.RGBA()
	return color.RGBA64{
		R: uint16(min(r1+r2, 65535)),
		G: uint16(min(g1+g2, 65535)),
		B: uint16(min(b1+b2, 65535)),
		A: uint16(min(a1+a2, 65535)),
	}
}

func LinearColorInterpolation(x1, x2 int, c1, c2 color.Color) []color.RGBA64 {
	r1, g1, b1, a1 := c1.RGBA()
	if x1 == x2 {
		return []color.RGBA64{
			{
				R: uint16(r1),
				G: uint16(g1),
				B: uint16(b1),
				A: uint16(a1),
			},
		}
	}
	if x1 > x2 {
		x1, x2 = x2, x1
		c1, c2 = c2, c1
	}
	dx := x2 - x1

	r1, g1, b1, a1 = c1.RGBA()
	r2, g2, b2, a2 := c2.RGBA()
	dr, dg, db, da := (int16(r2)-int16(r1))/int16(dx), (int16(g2)-int16(g1))/int16(dx), (int16(b2)-int16(b1))/int16(dx), (int16(a2)-int16(a1))/int16(dx)
	c := color.RGBA64{R: uint16(r1), G: uint16(g1), B: uint16(b1), A: uint16(a1)}
	res := make([]color.RGBA64, 0, dx+1)
	for x := x1; x <= x2; x++ {
		res = append(res, c)
		c.R = uint16(int16(c.R) + dr)
		c.G = uint16(int16(c.G) + dg)
		c.B = uint16(int16(c.B) + db)
		c.A = uint16(int16(c.A) + da)
	}
	return res
}
